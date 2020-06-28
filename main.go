package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// region no heartbeat client
//func ClientHandleError(err error, when string) {
//	if err != nil {
//		fmt.Println(err, when)
//		os.Exit(1)
//	}
//
//}
//
//func main() {
//
//	//拨号远程地址，简历tcp连接
//	conn, err := net.Dial("tcp", "127.0.0.1:7373")
//	ClientHandleError(err, "client conn error")
//
//	//预先准备消息缓冲区
//	buffer := make([]byte, 1024)
//
//	//准备命令行标准输入
//	reader := bufio.NewReader(os.Stdin)
//
//	for {
//		lineBytes, _, _ := reader.ReadLine()
//		conn.Write(lineBytes)
//		n, err := conn.Read(buffer)
//		ClientHandleError(err, "client read error")
//
//		serverMsg := string(buffer[0:n])
//		fmt.Printf("服务端msg", serverMsg)
//		if serverMsg == "bye" {
//			break
//		}
//
//	}
//
//}
// endregion no heartbeat client

// region with heartbeat client
func main() {
	server := "127.0.0.1:7373"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		Log(os.Stderr, "Fatal error:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		Log("Fatal error:", err.Error())
		os.Exit(1)
	}
	Log(conn.RemoteAddr().String(), "connection succcess!")

	sender(conn)
	Log("send over")
}
func sender(conn *net.TCPConn) {
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " Hello I'm MyHeartbeat Client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			os.Exit(1)
		}
		Log("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}
	for i := 0; i < 2; i++ {
		time.Sleep(12 * time.Second)
	}
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + " Hi I'm MyHeartbeat Client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), "Fatal error: ", err)
			os.Exit(1)
		}
		Log("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}

}
func Log(v ...interface{}) {
	fmt.Println(v...)
	return
}

// endregion with heartbeat client
