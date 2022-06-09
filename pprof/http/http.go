package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSerie(n int) []int {
	ret := make([]int, 2, n)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func index(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r)
	fmt.Println(string(body))
	resXml := `<?xml version="1.0" encoding="utf-8"?>
	<DATAS> 
	<DESCRIPTION>登录成功</DESCRIPTION>
	</DATAS>`
	w.Write([]byte(resXml))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = GetFibonacciSerie(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
