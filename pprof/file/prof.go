package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	r = 10000
	c = 10000
)

func set(m *[r][c]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[r][c]int) {
	for i := 0; i < r; i++ {
		tmp := 0
		for j := 0; j < c; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	f, _ := os.Create("cpu.prof")
	//cpu pprof
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	//简单代码运算
	a := [r][c]int{}
	set(&a)
	calculate(&a)

	//mem profile
	f1, _ := os.Create("mem.prof")
	runtime.GC() //内存回收
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	f1.Close()

	//goroutine prof
	f2, _ := os.Create("goroutine.prof")
	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write groutine profile: ")
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
}
