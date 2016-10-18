package prototest

import "testing"

import "github.com/golang/protobuf/proto"

func TestFullCustom(t *testing.T) {
	var x Small2

	x.A = proto.Int32(5)
	x.B = proto.Uint32(7)
	x.C = proto.String("abc")
	x.G = &Small2_G{X: proto.Float32(2.5)}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}
	var y Small2
	if y.MergeFullCustom(data) != nil {
		panic("bad merge")
	}
	if x.GetA() != y.GetA() {
		panic("bad A")
	}
	if x.GetB() != y.GetB() {
		panic("bad B")
	}
	if x.GetC() != y.GetC() {
		panic("bad C: " + y.GetC())
	}
	if x.GetG().GetX() != y.GetG().GetX() {
		panic("bad G.X")
	}
}

func TestTableDriven(t *testing.T) {
	var x Small2

	x.A = proto.Int32(5)
	x.B = proto.Uint32(7)
	x.C = proto.String("abc")
	x.G = &Small2_G{X: proto.Float32(2.5)}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}
	var y Small2
	if y.MergeTableDriven(data) != nil {
		panic("bad merge")
	}
	if x.GetA() != y.GetA() {
		panic("bad A")
	}
	if x.GetB() != y.GetB() {
		panic("bad B")
	}
	if x.GetC() != y.GetC() {
		panic("bad C: " + y.GetC())
	}
	if x.GetG().GetX() != y.GetG().GetX() {
		panic("bad G.X")
	}
}

func BenchmarkUnmarshal2(b *testing.B) {
	b.ReportAllocs()
	var x Small2

	x.A = proto.Int32(5)
	x.B = proto.Uint32(7)
	x.C = proto.String("abc")
	x.G = &Small2_G{X: proto.Float32(2.5)}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small2

	for i := 0; i < b.N; i++ {
		err = proto.Unmarshal(data, &y)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkUnmarshal2Buf(b *testing.B) {
	b.ReportAllocs()
	var x Small2

	x.A = proto.Int32(5)
	x.B = proto.Uint32(7)
	x.C = proto.String("abc")
	x.G = &Small2_G{X: proto.Float32(2.5)}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small2
	var p proto.Buffer

	for i := 0; i < b.N; i++ {
		y.Reset()
		p.SetBuf(data)
		err = p.Unmarshal(&y)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUnmarshal3(b *testing.B) {
	b.ReportAllocs()
	var x Small3

	x.A = 5
	x.B = 7
	x.C = "abc"
	x.G = &Group1{X: 2.5}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small3

	for i := 0; i < b.N; i++ {
		err = proto.Unmarshal(data, &y)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUnmarshal3Buf(b *testing.B) {
	b.ReportAllocs()
	var x Small3

	x.A = 5
	x.B = 7
	x.C = "abc"
	x.G = &Group1{X: 2.5}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small3
	var p proto.Buffer

	for i := 0; i < b.N; i++ {
		y.Reset()
		p.SetBuf(data)
		err = p.Unmarshal(&y)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMergeFullCustom2(b *testing.B) {
	b.ReportAllocs()
	var x Small2

	x.A = proto.Int32(5)
	x.B = proto.Uint32(7)
	x.C = proto.String("abc")
	x.G = &Small2_G{X: proto.Float32(2.5)}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small2

	for i := 0; i < b.N; i++ {
		y.Reset()
		y.MergeFullCustom(data)
	}
}

func BenchmarkMergeFullCustom3(b *testing.B) {
	b.ReportAllocs()
	var x Small3

	x.A = 5
	x.B = 7
	x.C = "abc"
	x.G = &Group1{X: 2.5}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small3

	for i := 0; i < b.N; i++ {
		y.Reset()
		y.MergeFullCustom(data)
	}
}

func BenchmarkMergeTableDriven2(b *testing.B) {
	b.ReportAllocs()
	var x Small2

	x.A = proto.Int32(5)
	x.B = proto.Uint32(7)
	x.C = proto.String("abc")
	x.G = &Small2_G{X: proto.Float32(2.5)}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small2

	for i := 0; i < b.N; i++ {
		y.Reset()
		y.MergeTableDriven(data)
	}
}

func BenchmarkMergeTableDriven3(b *testing.B) {
	b.ReportAllocs()
	var x Small3

	x.A = 5
	x.B = 7
	x.C = "abc"
	x.G = &Group1{X: 2.5}

	data, err := proto.Marshal(&x)
	if err != nil {
		panic(err)
	}

	var y Small3

	for i := 0; i < b.N; i++ {
		y.Reset()
		y.MergeTableDriven(data)
	}
}
