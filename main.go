package main

import (
	"TcpClient/msg"
	"bufio"
	"fmt"
	"net"
)

func main() {
	abc := msg.Req{Msg: "duwgadugwadui"}
	ss := make([]byte, 0)
	bb, _ := abc.XXX_Marshal(ss, true)
	fmt.Println(bb)
	ddd := &msg.Req{}
	err := ddd.XXX_Unmarshal(bb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ddd)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9930")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	fmt.Println("连接成功")
	defer conn.Close()
	if _, err = conn.Write(bb); err != nil {
		fmt.Printf("write failed , err : %v\n", err)
		return
		//break
	}
	fmt.Println(bb)
	reader := bufio.NewReader(conn)
	for {
		var buf [100]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到Client发来的数据：", recvStr)
		recvStr = "server" + recvStr
		conn.Write([]byte(recvStr))
	}
}

//}
