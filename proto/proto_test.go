package proto

import (
	"code.google.com/p/goprotobuf/proto"
	// "fmt"
	"os"
	"testing"
)

var (
	msg        = &MyMessage{Count: proto.Int32(1234), Name: proto.String("protoName"), Quote: proto.String("protoQuote")}
	mardata, _ = proto.Marshal(msg)
	marTextStr = proto.MarshalTextString(msg)
)

func TestMarshal(t *testing.T) {
	msg_new := new(MyMessage)
	err := proto.Unmarshal([]byte(marTextStr), msg_new)
	checkerr(err)
}

func TestUnMarshalText(t *testing.T) {
	msg_new := new(MyMessage)
	err := proto.UnmarshalText(string(mardata), msg_new)
	checkerr(err)
}

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(msg)
		if checkerr(err) {
			break
		}
	}
}

func BenchmarkMarshalText(b *testing.B) {
	b.StopTimer()
	file, _ := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := proto.MarshalText(file, msg)
		if checkerr(err) {
			break
		}
	}
}

func BenchmarkUnMarshal(b *testing.B) {
	var msg_new MyMessage
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(mardata, &msg_new)
		if checkerr(err) {
			break
		}
	}
}

func BenchmarkUnMarshalText(b *testing.B) {
	var msg_new MyMessage
	for i := 0; i < b.N; i++ {
		err := proto.UnmarshalText(marTextStr, &msg_new)
		if checkerr(err) {
			break
		}
	}
}
