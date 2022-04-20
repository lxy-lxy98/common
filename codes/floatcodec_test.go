package codes

import (
	"io/ioutil"
	"math"
	"os"
	"testing"
)

func Test_byteToFloat32Array(t *testing.T) {
	file, err := os.Open("1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("read err %v", err)
	}
	content = content[:2112]
	feature := ByteConvertFloat32Array(content, len(content)/4)
	for _, value := range feature {
		if math.IsNaN(float64(value)) {
			t.Fatal("invalid feature value nan float")
		}
	}
	t.Log(feature)
}

func Benchmark_float32ToByte(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		Float32ConvertByte(1)
	}
}

func BenchmarkParallel_float32ToByte(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Float32ConvertByte(1)
		}
	})
}
