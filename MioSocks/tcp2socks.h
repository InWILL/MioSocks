#pragma once

#include "base.h"
#include <ws2tcpip.h>
#include "../socks5/socks5.h"

#pragma comment(lib,"socks5.lib")

int Tcp2Socks_Process();