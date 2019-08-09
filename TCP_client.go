package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

//수신까지 걸리는 프로세스타임
/*func process_run_time() int64 {
	elaspedtime := time.Since()
	fmt.Println(elaspedtime)
	return 0
}*/

//수신직후 현재시간 스탬프
func nanotime() int64 {
	now := time.Now()
	epoch := now.UnixNano()
	return epoch
}
func main() {

	argu := os.Args
	if len(argu) == 1 {
		fmt.Println("plz provide host:port")
		return
	}

	CONNECT := argu[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		now := time.Now() //현재 시간 타임스탬프
		fmt.Println("->: " + message)
		fmt.Printf("Time Now = %d", nanotime())
		elaspedtime := time.Since(now) //프로세스 처리시간 타임 스탬프
		fmt.Printf("Process Time : %s\n", elaspedtime)
	}
}
