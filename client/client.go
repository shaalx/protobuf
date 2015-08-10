package main

import (
	"code.google.com/p/goprotobuf/proto"
	// "code.google.com/p/goprotobuf/proto/testdata"
	"fmt"
	"net"
	// "os"
	// "bytes"
	// "io"
	// "io/ioutil"
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
	var123 := int32(123)
	name := "myname"
	quote := "myquote"
	for {
		// send
		time.Sleep(3e9)
		msg := MyMessage{Count: &var123, Name: &name, Quote: &quote}
		mterr := proto.MarshalText(conn, &msg)
		if checkerr(mterr) {
			break
		}

		// // receive
		// bsbuf := make([]byte, 128)
		// dst := bytes.NewBuffer(bsbuf)
		// io.Copy(dst, conn)
		// b, err := ioutil.ReadAll(dst)
		// if checkerr(err) {
		// 	break
		// }
		// fmt.Printf("received %d byte from server.", len(b))
		// var rec_msg MyMessage
		// umterr := proto.UnmarshalText(string(b), &rec_msg)
		// if checkerr(umterr) {
		// 	break
		// }
		// fmt.Printf("%#v the msg:%s\n", rec_msg, rec_msg.toString())

		receiveMsg(conn)

		// buf := make([]byte, 128)
		// c, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("读取服务器数据异常:", err.Error())
		// }
		// fmt.Println(string(buf[0:c]))
		// var rec_msg2 MyMessage
		// umterr2 := proto.UnmarshalText(string(buf[:c]), &rec_msg2)
		// if checkerr(umterr2) {
		// 	break
		// }
		// fmt.Printf("%#v \n %s\n", rec_msg2, rec_msg2.toString())
	}

}

func receiveMsg(conn net.Conn) {
	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if checkerr(err) {
		return
	}
	readStr := string(buf[:n])
	fmt.Printf("read string (%d):\n%s\n", n, readStr)

	var msg MyMessage

	// umterr := proto.UnmarshalText(readStr, &msg)
	// if checkerr(umterr) {
	// 	return
	// }
	// fmt.Printf("[MyMessage] %v\n", msg)
	// fmt.Printf("[MyMessage] %v\n", msg.toString())

	//
	// umerr := proto.Unmarshal(buf[:n], &msg)
	// if checkerr(umerr) {
	// 	return
	// }
	// fmt.Printf("[MyMessage] %v\n", msg)
	// fmt.Printf("[MyMessage] %v\n", msg.toString())

	protobuf := proto.NewBuffer(buf[:n])
	// protobuf.Reset()
	// protobuf.SetBuf(buf[:n])
	pumerr := protobuf.Unmarshal(&msg)

	if checkerr(pumerr) {
		return
	}
	fmt.Printf("[MyMessage] %v\n", msg.String())
	fmt.Printf("[MyMessage] %v\n", msg.toString())
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
