package proto

import (
	"code.google.com/p/goprotobuf/proto"
	"code.google.com/p/goprotobuf/proto/testdata"
	"fmt"
	"os"
)

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}

func Marshal() {
	var Test testdata.GoEnum
	Test.Foo = new(testdata.FOO)
	fmt.Println(Test)

	data, err := proto.Marshal(&Test)
	if checkerr(err) {
		return
	}
	fmt.Println(data)

	var goe testdata.GoEnum
	err = proto.Unmarshal(data, &goe)
	if checkerr(err) {
		return
	}

	fmt.Println(goe)

	var i proto.Message
	err = proto.Unmarshal(data, i)
	if checkerr(err) {
		return
	}

	fmt.Println(i)
}

func marchalText() {
	m := &MyMessage{Count: proto.Int32(1234), Name: proto.String("protoName"), Quote: proto.String("protoQuote")}

	// MarshalTextString
	fmt.Println("// MarshalTextString...")
	mstr := proto.MarshalTextString(m)
	fmt.Println(mstr)
	var target MyMessage
	proto.UnmarshalText(mstr, &target)
	fmt.Printf("%#v\n", target)
	fmt.Printf("%s\n", target.toString())
	fmt.Println("// MarshalTextString...\n")

	// MarshalText
	fmt.Println("// MarshalText...")
	mterr := proto.MarshalText(os.Stdout, m)
	checkerr(mterr)
	var target1 MyMessage
	umterr := proto.UnmarshalText(mstr, &target1)
	checkerr(umterr)
	fmt.Printf("%#v\n", target1)
	fmt.Println(*(target1.Count))
	fmt.Println("// MarshalText...\n")

	// Marshal
	fmt.Println("// Marshal...")
	data, _ := proto.Marshal(m)
	fmt.Println(data)
	var target2 MyMessage
	uerr := proto.Unmarshal(data, &target2)
	checkerr(uerr)
	fmt.Printf("%#v\n", target2)
	fmt.Println("// Marshal...\n")

	// // MarshalSet
	// fmt.Println("// MarshalSet...")
	// var mset proto.MessageSet
	// merr := mset.Marshal(m)
	// checkerr(merr)
	// fmt.Println(mset.Item)
	// fmt.Println(mset.XXX_unrecognized)
	// var target3 MyMessage
	// mset.Unmarshal(&target3)
	// fmt.Printf("%#v\n", target3)
}

type MyMessage struct {
	Count *int32  `protobuf:"varint,1,req,name=count" json:"count,omitempty"`
	Name  *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Quote *string `protobuf:"bytes,3,opt,name=quote" json:"quote,omitempty"`

	// Count *int32  `protobuf:"varint,1,req,name=count" json:"count,omitempty"`
	// Name  *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Quote *string `protobuf:"bytes,3,opt,name=quote" json:"quote,omitempty"`

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
	return fmt.Sprintf("count:%d, name:%s, quote:%s", m.Count, m.Name, m.Quote)
}

func main1() {
	marchalText()
}
