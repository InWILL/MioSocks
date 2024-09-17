# MioSocks

![GitHub License](https://img.shields.io/github/license/InWILL/MioSocks)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/InWILL/MioSocks/total)

### English

Tcp2socks implementation by WinDivert.

The project provides a solution to convert tcp connections to socks5.

From 2023.11.29, MapleStory-Japan updated Nexon Game Security(NGS), it had been detecting [Netch](https://github.com/netchx/netch) Process Mode as suspicious activity. Tun Mode can use normally, but Tun mode means all connections have to be proxyed. So I base on WinDivert to design a tcp2socks software.

## Download

Download available at <https://github.com/InWILL/MioSocks/releases/latest>.

## Getting Started

### Prerequisites

* [Microsoft Visual C++ Redistributable](https://docs.microsoft.com/en-US/cpp/windows/latest-supported-vc-redist?view=msvc-170)

## Example

settings.json for MapleStory
```
{
	"socks5_server": "127.0.0.1",
	"socks5_port": 2801,
	"tcp_port": 2805,
	"WindivertOpen": "tcp and localAddr != :: and remoteAddr != :: and (tcp.DstPort == 8484 or tcp.DstPort == 8585 or tcp.DstPort == 8586 or tcp.DstPort == 8589 or tcp.DstPort == 8787 or tcp.DstPort == 8788 or tcp.DstPort == 443 or tcp.SrcPort == 2805)"
}
```

## Built With

* [WinDivert](https://github.com/basil00/WinDivert) - Used to handle tcp connection

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/InWILL/MioSocks/tags). 

## Contributors

See also the list of [contributors](https://github.com/InWILL/MioSocks/graphs/contributors) who participated in this project.

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.