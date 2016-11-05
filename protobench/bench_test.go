package protobench

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

// protoc --go_out=. msg.proto
// go test -test.bench=.*

func BenchmarkUnmarshalA(b *testing.B) {
	benchmark(b, genA())
}
func BenchmarkUnmarshalB(b *testing.B) {
	benchmark(b, genB())
}

func BenchmarkAStd(b *testing.B) {
	save := proto.UnmarshalKind
	proto.UnmarshalKind = proto.UnmarshalStd
	benchmark(b, genA())
	proto.UnmarshalKind = save
}

func BenchmarkATable(b *testing.B) {
	save := proto.UnmarshalKind
	proto.UnmarshalKind = proto.UnmarshalTable
	benchmark(b, genA())
	proto.UnmarshalKind = save
}

func BenchmarkAGen(b *testing.B) {
	save := proto.UnmarshalKind
	proto.UnmarshalKind = proto.UnmarshalGen
	benchmark(b, genA())
	proto.UnmarshalKind = save
}

func BenchmarkBStd(b *testing.B) {
	save := proto.UnmarshalKind
	proto.UnmarshalKind = proto.UnmarshalStd
	benchmark(b, genB())
	proto.UnmarshalKind = save
}

func BenchmarkBTable(b *testing.B) {
	save := proto.UnmarshalKind
	proto.UnmarshalKind = proto.UnmarshalTable
	benchmark(b, genB())
	proto.UnmarshalKind = save
}

func BenchmarkBGen(b *testing.B) {
	save := proto.UnmarshalKind
	proto.UnmarshalKind = proto.UnmarshalGen
	benchmark(b, genB())
	proto.UnmarshalKind = save
}

func genA() proto.Message {
	var a A
	a.F1 = 892
	a.F2 = 33
	a.F3 = 300
	a.F4 = true
	a.F5 = 37.3
	a.F6 = 99.9
	a.F7 = 22
	a.F8 = 23
	a.F9 = -32
	a.F10 = -111
	return &a
}

func genB() proto.Message {
	return &B{
		F1: []int64{7, 77, 777, 7777, 77777, 777777},
		F2: &B2{
			F1: []float32{3.875},
			F2: -99,
			F3: []*B3{
				&B3{
					F1: []string{"abc", "yyy"},
					F2: "abcdef",
				},
				&B3{
					F1: []string{"abc", "yyy", "zzz"},
					F2: "abcdefghi",
				},
			},
		},
		F3: []*B2{},
		F4: "foo",
	}
}

func benchmark(b *testing.B, m proto.Message) {
	b.ReportAllocs()

	// Encode the message.
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}

	// Decode, then encode again, just to make sure unmarshal works.
	m.Reset()
	err = proto.Unmarshal(data, m)
	if err != nil {
		panic(err)
	}
	data2, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, data2) {
		fmt.Printf("%v\n", data)
		fmt.Printf("%v\n", data2)
		panic("bytes not equal")
	}

	// Repeatedly unmarshal.
	var buf proto.Buffer
	for i := 0; i < b.N; i++ {
		m.Reset()
		buf.SetBuf(data)
		err = buf.Unmarshal(m)
		if err != nil {
			panic(err)
		}
	}
}
