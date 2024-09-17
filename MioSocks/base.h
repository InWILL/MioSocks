#pragma once

#include <iostream>
#include <thread>
#include <winsock2.h>
#include "../WinDivert/windivert.h"

#pragma comment(lib,"../WinDivert/WinDivert.lib")
#pragma comment(lib,"Ws2_32.lib")

//struct NetTuple
//{
//	UINT32 SrcAddr;
//	UINT16 SrcPort;
//	UINT32 DstAddr;
//	UINT16 DstPort;
//	char ProcessName[MAX_PATH];
//};

enum State
{
    STATE_NOT_CONNECTED     = 0,
    STATE_SYN_WAIT          = 1,
    STATE_SYN_SEEN          = 2,
    STATE_SYNACK_SEEN       = 3,
    STATE_ESTABLISHED       = 4,
    STATE_FIN_WAIT          = 5,
    STATE_WHITELISTED       = 0xFF,
};

// Pended SYN.
struct SYN
{
    WINDIVERT_ADDRESS addr;
    UINT32 packet_len;
    UINT8 packet[];
};

struct Connection
{
    State state;
    UINT32 SrcAddr;
    UINT16 SrcPort;
    UINT32 DstAddr;
    UINT16 DstPort;
    SYN* syn;
};

extern Connection conns[UINT16_MAX + 1];