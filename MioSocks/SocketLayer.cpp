#include "SocketLayer.h"

#include <windows.h>
#include <psapi.h>
#include <shlwapi.h>
#include <stdio.h>
#include <stdlib.h>

#pragma comment(lib,"Shlwapi.lib")

#define INET6_ADDRSTRLEN    45

Connection conns[UINT16_MAX + 1];

void handle_syn(HANDLE handle, Connection *conn)
{
	SYN* syn = conn->syn;
	PWINDIVERT_IPHDR ip_header = NULL;
	PWINDIVERT_TCPHDR tcp_header = NULL;
	if (syn == NULL)
	{
		return;
	}
	WinDivertHelperParsePacket(
		syn->packet, 
		syn->packet_len, 
		&ip_header, 
		NULL, 
		NULL, 
		NULL, 
		NULL, 
		&tcp_header, 
		NULL, 
		NULL, 
		NULL, 
		NULL, 
		NULL
	);
	if (ip_header == NULL || tcp_header == NULL)
	{
		free(syn);
		return;
	}
	
	switch (conn->state)
	{
	case STATE_WHITELISTED:
	{
		syn->addr.Outbound = TRUE;
		WinDivertHelperCalcChecksums(syn->packet, syn->packet_len, &syn->addr, 0);
		if (!WinDivertSend(handle, syn->packet, syn->packet_len, NULL, &syn->addr))
		{
			// Handle send error
		}
		free(syn);
		break;
	}
	case STATE_SYN_WAIT:
	{
		conn->SrcAddr = ip_header->SrcAddr;
		conn->SrcPort = tcp_header->SrcPort;
		conn->DstAddr = ip_header->DstAddr;
		conn->DstPort = tcp_header->DstPort;

		UINT32 dst_addr = ip_header->DstAddr;
		tcp_header->DstPort = htons(2805);
		ip_header->DstAddr = ip_header->SrcAddr;
		ip_header->SrcAddr = dst_addr;

		printf("[SOCKET] [CONNECT] %u:%u %u:%u\n", conn->SrcAddr, conn->SrcPort, conn->DstAddr, conn->DstPort);

		syn->addr.Outbound = FALSE;
		WinDivertHelperCalcChecksums(syn->packet, syn->packet_len, &syn->addr, 0);
		if (!WinDivertSend(handle, syn->packet, syn->packet_len, NULL, &syn->addr))
		{
			// Handle send error
		}
		free(syn);
		break;
	}
	}
}

// Pend a CONNECT:
void pend_connect(HANDLE handle, UINT16 local_port, State state)
{
	Connection *conn = &conns[local_port];
	switch (conn->state)
	{
	case STATE_NOT_CONNECTED:
	case STATE_FIN_WAIT:
		conn->state = state;
		handle_syn(handle, conn);
		break;
	default:
		break;
	}
}

void Socket_Layer_Process()
{
	HANDLE handle;          // WinDivert socket handle
	HANDLE inject;			// WinDivert network handle
	HANDLE process;
	DWORD path_len;

	char path[MAX_PATH];
	char* filename{};

	// Open some filter
	handle = WinDivertOpen(
		"tcp and (event == CONNECT or event == CLOSE) and localAddr != :: and remoteAddr != ::", 
		WINDIVERT_LAYER_SOCKET, 
		11234, 
		WINDIVERT_FLAG_SNIFF | WINDIVERT_FLAG_RECV_ONLY
	);
	inject = WinDivertOpen(
		"false",
		WINDIVERT_LAYER_NETWORK,
		1751,
		WINDIVERT_FLAG_SEND_ONLY
	);
	if (handle == INVALID_HANDLE_VALUE || inject == INVALID_HANDLE_VALUE)
	{
		// Handle error
		exit(EXIT_FAILURE);
	}

	// Main capture-modify-inject loop:
	while (TRUE)
	{
		// Packet address
		WINDIVERT_ADDRESS addr;
		if (!WinDivertRecv(handle, NULL, 0, NULL, &addr))
		{
			// Handle recv error
			continue;
		}

		/*printf(" pid=");
		printf("%u", addr.Socket.ProcessId);

		printf(" program=");*/
		process = OpenProcess(PROCESS_QUERY_LIMITED_INFORMATION, FALSE,
			addr.Socket.ProcessId);
		path_len = 0;
		if (process != NULL)
		{
			path_len = GetProcessImageFileNameA(process, path, sizeof(path));
			CloseHandle(process);
		}
		if (path_len != 0)
		{
			filename = PathFindFileNameA(path);
			/*printf("%s", filename);*/
		}

		UINT16 local_port = htons(addr.Socket.LocalPort);
		switch (addr.Event)
		{
		case WINDIVERT_EVENT_SOCKET_CONNECT:
		{
			if (strcmp(filename, "Chrome.exe") == 0)
			{
				pend_connect(inject, local_port, STATE_SYN_WAIT);
			}
			else
			{
				pend_connect(inject, local_port, STATE_WHITELISTED);
			}
			break;
		}
		case WINDIVERT_EVENT_SOCKET_CLOSE:
			Connection* conn = &conns[local_port];
			SYN* syn = conn->syn;
			if (conn->state == STATE_NOT_CONNECTED ||
				conn->state == STATE_FIN_WAIT)
			{
				/* Connection Closed */
				continue;
			}

			printf("[SOCKET] [CLOSE] %u:%u %u:%u\n", conn->SrcAddr, conn->SrcPort, conn->DstAddr, conn->DstPort);

			free(syn);
			conn->state = STATE_FIN_WAIT;
			conn->syn = NULL;
			break;
		}
	}
}