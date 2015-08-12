package main

import (
	// "log"
	// // 辅助库
	"code.google.com/p/goprotobuf/proto"
	"github.com/everfore/protobuf/Person"
	// "code.google.com/p/goprotobuf/proto/testdata"
	// // test.pb.go 的路径
	// "example"

	"fmt"
	// "io/ioutil"
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
				fmt.Printf("read %d byte.\n", n)
				msgstr := string(data[:n])
				fmt.Printf("  msg to string:\n%s\n", msgstr)
				var msg Person.Person

				umerr := proto.Unmarshal(data[:n], &msg)
				if checkerr(umerr) {
					break
				}
				fmt.Println(msg.String())
				// umterr := proto.UnmarshalText(msgstr, &msg)
				// if checkerr(umterr) {
				// 	break
				// }
				// fmt.Printf("[MyMessage]:%#v\n%s\n", msg, msg.toString())

				// // send
				// mterr := proto.MarshalText(conn, &msg)
				// if checkerr(mterr) {
				// 	break
				// }

				data, err := proto.Marshal(&msg)
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

type MyMessage struct {
	Count *int32  `protobuf:"varint,1,req,name=count" json:"count,omitempty"`
	Name  *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Quote *string `protobuf:"bytes,3,opt,name=quote" json:"quote,omitempty"`
	// Pet       []string             `protobuf:"bytes,4,rep,name=pet" json:"pet,omitempty"`
	// Inner     *InnerMessage        `protobuf:"bytes,5,opt,name=inner" json:"inner,omitempty"`
	// Others    []*OtherMessage      `protobuf:"bytes,6,rep,name=others" json:"others,omitempty"`
	// RepInner  []*InnerMessage      `protobuf:"bytes,12,rep,name=rep_inner" json:"rep_inner,omitempty"`
	// Bikeshed  *MyMessage_Color     `protobuf:"varint,7,opt,name=bikeshed,enum=testdata.MyMessage_Color" json:"bikeshed,omitempty"`
	// Somegroup *MyMessage_SomeGroup `protobuf:"group,8,opt,name=SomeGroup" json:"somegroup,omitempty"`
	// // This field becomes [][]byte in the generated code.
	// RepBytes         [][]byte                  `protobuf:"bytes,10,rep,name=rep_bytes" json:"rep_bytes,omitempty"`
	// Bigfloat         *float64                  `protobuf:"fixed64,11,opt,name=bigfloat" json:"bigfloat,omitempty"`
	// XXX_extensions   map[int32]proto.Extension `json:"-"`
	// XXX_unrecognized []byte                    `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (m *MyMessage) toString() string {
	return fmt.Sprintf("count:%d, name:%s, quote:%s", *(m.Count), *(m.Name), *(m.Quote))
}
