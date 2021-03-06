//https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/

package main

import (
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("... Error: ", err)
		os.Exit(0)
	}
}

func main() {
	//Create socket
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:23000")
	CheckError(err)

	//Listen
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	defer ServerConn.Close()

	ServerBuf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(ServerBuf)
		fmt.Println("... Received ", string(ServerBuf[0:n]), "from ", addr)

		if err != nil {
			fmt.Println("... Error: ", err)
		}
	}
}
