package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/*func protocolselector() {
}*/

func main() {
	argu := os.Args
	if len(argu) == 1 {
		fmt.Println("Plz provide port num")
		return
		//fmt.Println("Waiting to Packet")
	}
	PORT := ":" + argu[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		netdata, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netdata)) == "STOP" {
			fmt.Println("Exiting TCP Server")
			return
		}
		fmt.Print("->", string(netdata))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
