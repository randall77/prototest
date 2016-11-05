package protobench

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func BenchmarkSizeA(b *testing.B) {
	benchmarkSize(b, genA())
}
func BenchmarkSizeB(b *testing.B) {
	benchmarkSize(b, genB())
}

func benchmarkSize(b *testing.B, m proto.Message) {
	// Encode the message.
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}

	// Check the result.
	size := proto.SizeIt(m)
	if size != len(data) {
		panic("size wrong")
	}

	// Repeatedly compute size.
	var buf proto.Buffer
	s := 0
	for i := 0; i < b.N; i++ {
		s += buf.SizeIt(m)
	}
	sink = s
}

var sink int
