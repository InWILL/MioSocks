#include "NetworkLayer.h"

const int MAXBUF = 65536;

void pend_syn(HANDLE handle, UINT16 local_port, char* packet, UINT32 packet_len, WINDIVERT_ADDRESS* addr)
{
	SYN* syn = (SYN*)malloc(sizeof(SYN) + packet_len);
	if (syn == NULL)
	{
		exit(EXIT_FAILURE);
	}
	memcpy(&syn->addr, addr, sizeof(syn->addr));
	syn->packet_len = packet_len;
	memcpy(syn->packet, packet, packet_len);

	Connection* conn = &conns[local_port];
	SYN* old_syn = conn->syn;
	conn->syn = syn;
	handle_syn(handle, conn);
	free(old_syn);
}

int TCP_Proxy_Process()
{
	
	HANDLE handle;          // WinDivert handle

	// Open some filter
	handle = WinDivertOpen("tcp and localAddr != :: and remoteAddr != ::", WINDIVERT_LAYER_NETWORK, 0, 0);
	if (handle == INVALID_HANDLE_VALUE)
	{
		// Handle error
		exit(1);
	}

	// Main capture-modify-inject loop:
	while (TRUE)
	{
		static char packet[MAXBUF];    // Packet buffer
		static UINT packetLen;

		WINDIVERT_ADDRESS	addr; // Packet address
		PWINDIVERT_IPHDR	ip_header;
		PWINDIVERT_TCPHDR	tcp_header;
		if (!WinDivertRecv(handle, packet, sizeof(packet), &packetLen, &addr))
		{
			// Handle recv error
			continue;
		}

		// Modify packet.
		WinDivertHelperParsePacket(packet, packetLen, &ip_header, NULL, NULL,
			NULL, NULL, &tcp_header, NULL, NULL, NULL, NULL, NULL);

		if (ip_header == NULL || tcp_header == NULL)
		{
			printf("failed to parse packet :%d \n", GetLastError());
			continue;
		}

		UINT32 srcaddr = ip_header->SrcAddr;
		UINT16 srcport = ntohs(tcp_header->SrcPort);
		UINT32 dstaddr = ip_header->DstAddr;
		UINT16 dstport = ntohs(tcp_header->DstPort);

		if (ip_header->Version == 4)
		{
			if (tcp_header->SrcPort == htons(2805))
			{
				printf("[<-]%u:%u %u:%u\n", srcaddr, srcport, dstaddr, dstport);

				UINT32 dst_addr = ip_header->DstAddr;
				tcp_header->SrcPort = htons(conns[dstport].DstPort);
				ip_header->DstAddr = ip_header->SrcAddr;
				ip_header->SrcAddr = dst_addr;
				addr.Outbound = FALSE;
			}
			else if (dstport == 8484 || 
				dstport == 8585 || 
				dstport == 8586 || 
				dstport == 8589 || 
				dstport == 8787 || 
				dstport == 8788 || 
				dstport == 443
				)
			{
				printf("[->]%u:%u %u:%u\n", srcaddr, srcport, dstaddr, dstport);
				conns[srcport].SrcAddr = srcaddr;
				conns[srcport].SrcPort = srcport;
				conns[srcport].DstAddr = dstaddr;
				conns[srcport].DstPort = dstport;
				
				UINT32 dst_addr = ip_header->DstAddr;
				tcp_header->DstPort = htons(2805);
				ip_header->DstAddr = ip_header->SrcAddr;
				ip_header->SrcAddr = dst_addr;
				addr.Outbound = FALSE;
			}
		}

		WinDivertHelperCalcChecksums(packet, packetLen, &addr, 0);
		if (!WinDivertSend(handle, packet, packetLen, NULL, &addr))
		{
			// Handle send error
			continue;
		}
	}
}