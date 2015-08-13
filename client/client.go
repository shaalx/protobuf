package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/shaalx/protobuf/Person"
	"net"
	"time"
)

const (
	addr = "127.0.0.1:8080"
)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	Client(conn)
}

func Client(conn net.Conn) {
	msg := Person.Person{Name: proto.String("shiyongbin"), Id: proto.Int32(111)}
	for {
		time.Sleep(1e9)

		// send
		data, err := proto.Marshal(&msg)
		if checkerr(err) {
			break
		}
		conn.Write(data)

		// // receive
		receiveMsg(conn)

		time.Sleep(4e9)
	}

}

func receiveMsg(conn net.Conn) {
	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if checkerr(err) {
		return
	}
	fmt.Printf("\nread (%d) byte from %v :\n%v\n", n, conn.RemoteAddr(), buf[:n])

	protobuf := proto.NewBuffer(buf[:n])
	var msg Person.Person
	pumerr := protobuf.Unmarshal(&msg)
	if checkerr(pumerr) {
		return
	}
	fmt.Printf("[Message]----->%v\n\n", msg.String())

	//
	// umerr := proto.Unmarshal(buf[:n], &msg)
	// if checkerr(umerr) {
	// 	return
	// }
	// fmt.Printf("[MyMessage] %v\n", msg)

	// readStr := string(buf[:n])
	// fmt.Printf("read string (%d):\n%s\n", n, readStr)
	// umterr := proto.UnmarshalText(readStr, &msg)
	// if checkerr(umterr) {
	// 	return
	// }
	// fmt.Printf("[MyMessage] %v\n", msg)
}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
