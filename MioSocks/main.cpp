#include "base.h"
#include "NetworkLayer.h"
#include "tcp2socks.h"
#include "SocketLayer.h"

#include <io.h>
#include <fstream>
#include <nlohmann/json.hpp>

int main()
{
	if (access("settings.json",0) != 0)
	{
		std::ofstream output("settings.json");
		output << R"(
{
	"socks5_server": "127.0.0.1",
	"socks5_port": 2801,
	"tcp_port": 2805,
	"WindivertOpen": "tcp and localAddr != :: and remoteAddr != ::"
}
				)";
		output.close();
	}
	
	std::ifstream input("settings.json");
	nlohmann::json settings = nlohmann::json::parse(input);
	input.close();
	
	std::string server = settings["socks5_server"];
	strcpy(socks5_server, server.c_str());
	socks5_port = settings["socks5_port"];
	tcp_port = settings["tcp_port"];
	std::string divertopen = settings["WindivertOpen"];
	strcpy(windivertopen, divertopen.c_str());

	//std::thread th1(Socket_Layer_Process);
	std::thread th2(Tcp2Socks_Process);
	//th1.detach();
	th2.detach();
	TCP_Proxy_Process();
	return 0;
}