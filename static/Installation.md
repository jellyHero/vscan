# 编译

```sh
sudo apt install -y libpcap-dev golang git
git clone https://github.com/jellyHero/vscan.git
cd vscan
go build
```

# 安装/运行

1.在运行vscan之前，你必须先安装libpcap库

```sh
sudo apt install -y libpcap-dev
```

2.前往
[https://github.com/jellyHero/vscan/releases/](https://github.com/jellyHero/vscan/releases/)
下载vscan最新版运行:

## 运行时动态库版本问题

如果你运行的时候出现了`libpcap.so.0.8: cannot open shared object file: No such file or directory`的错误

请先检查libpcap库是否已经正常安装。
```sh
ls -all /lib64/libpcap*
```
如果有安装其他版本的libpcap库，可建立一个软连接到/lib64/libpcap.so.0.8即可正常运行程序

```sh
ln -s /lib64/libpcap.so.1.9.1 /lib64/libpcap.so.0.8
```