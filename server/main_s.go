package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/toukii/protobuf/Person"
	// "github.com/toukii/protobuf/Person2"
	"net"
)

const (
	ip   = ""
	port = 8080
)

func main() {
	Server()
}

func Server() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接...")
	_Server(listen)
}

func _Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 1024)
			for {
				// receive
				n, err := conn.Read(data)
				if checkerr(err) {
					break
				}
				fmt.Printf("\nread (%d) byte from %v :\n%v\n", n, conn.RemoteAddr(), data[:n])
				var msg Person.Person

				umerr := proto.Unmarshal(data[:n], &msg)
				if checkerr(umerr) {
					break
				}
				fmt.Printf("[Message]---->%v\n\n", msg.String())

				// msgstr := string(data[:n])
				// fmt.Printf("  msg to string:\n%s\n", msgstr)
				// umterr := proto.UnmarshalText(msgstr, &msg)
				// if checkerr(umterr) {
				// 	break
				// }
				// fmt.Printf("[MyMessage]:%#v\n%s\n", msg, msg.String())

				// // send

				// mterr := proto.MarshalText(conn, &msg)
				// if checkerr(mterr) {
				// 	break
				// }

				new_msg := Person.Person{Id: proto.Int32(111), Name: proto.String("Shiyongbin"), Email: proto.String("toukii@163.com")}
				data, err := proto.Marshal(&new_msg)
				if checkerr(err) {
					break
				}
				conn.Write(data)
			}

		}()
	}
}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
