# MioSocks

![GitHub License](https://img.shields.io/github/license/InWILL/MioSocks)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/InWILL/MioSocks/total)

### English | [简体中文](https://github.com/InWILL/MioSocks#english--%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)

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


### [English](https://github.com/InWILL/MioSocks#english--%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87-1) | 简体中文

WinDivert 实现的 Tcp2socks。

该项目提供了将 tcp 连接转换为 socks5 的解决方案。

从 2023.11.29 开始，日服冒险岛更新了Nexon Game Security(NGS)，一直将 [Netch](https://github.com/netchx/netch) 的进程模式检测为可疑活动。虽然Tun模式可以正常使用，但Tun模式意味着所有连接都必须代理。所以我基于WinDivert设计了一个tcp2socks软件。

## 下载

可在 <https://github.com/InWILL/MioSocks/releases/latest> 下载。

## 入门

### 安装要求

* [Microsoft Visual C++ Redistributable](https://docs.microsoft.com/en-US/cpp/windows/latest-supported-vc-redist?view=msvc-170)

## 示例

settings.json for MapleStory
```
{
	"socks5_server": "127.0.0.1",
	"socks5_port": 2801,
	"tcp_port": 2805,
	"WindivertOpen": "tcp and localAddr != :: and remoteAddr != :: and (tcp.DstPort == 8484 or tcp.DstPort == 8585 or tcp.DstPort == 8586 or tcp.DstPort == 8589 or tcp.DstPort == 8787 or tcp.DstPort == 8788 or tcp.DstPort == 443 or tcp.SrcPort == 2805)"
}
```

## 内置

* [WinDivert](https://github.com/basil00/WinDivert) - 用于处理 tcp 连接

## 版本控制

我们使用 [SemVer](http://semver.org/) 进行版本控制。有关可用版本，请参阅 [此存储库上的标签](https://github.com/InWILL/MioSocks/tags)。

## 贡献者

另请参阅参与此项目的 [贡献者](https://github.com/InWILL/MioSocks/graphs/contributors) 列表。

## 许可证

该项目在 GPL-3.0 许可下获得许可 - 请参阅 [LICENSE](LICENSE) 文件了解详细信息。