package Reflect

import (
	resultSlicePageLimit "common/reflect"
	"fmt"
	"testing"
)

func Test_CommonResultSlicePageLimit(t *testing.T) {
	tests := []struct {
		offset, maxResult int
		inSlice           interface{}
		outSlice          interface{}
		want              int
	}{
		{0, 1, []int{0}, &[]int{}, -1},
	}
	for _, tt := range tests {
		got := resultSlicePageLimit.CommonResultSlicePageLimit(tt.offset, tt.maxResult, tt.inSlice, tt.outSlice)
		if got != tt.want {
			t.Errorf("CommonResultSlicePageLimit(%v): got %v, want %v", tt, got, tt.want)
		}
	}
	offset, maxResult := 0, 2
	inSlice := []int{0, 1}
	var outSlice []int
	offset = resultSlicePageLimit.CommonResultSlicePageLimit(offset, maxResult, inSlice, &outSlice)
	fmt.Println(offset, maxResult, inSlice, &outSlice)
}
