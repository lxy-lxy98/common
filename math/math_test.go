package math

import (
	"fmt"
	"testing"
)

func TestFloatDecimal(t *testing.T) {
	f := 501.118918915
	prec := 2
	//1.   10*10 = 100
	//2.   f*100 = 50111.****
	//3.   50111
	//4.   501.11
	fmt.Println(FloatDecimal(f, prec))
}

func Test_FloatDecimal(t *testing.T) {
	tests := []struct {
		value, want float64
		bit         int
	}{
		{0.130004, 0.13, 5},
	}
	for _, tt := range tests {
		got := FloatDecimal(tt.value, tt.bit)
		t.Logf("want %v, got %v", tt.want, got)
		if !FloatEqual(got, tt.want, 0.000001) {
			t.Fatalf("FloatDecimal err test example %v", tt)
		}
	}
}
