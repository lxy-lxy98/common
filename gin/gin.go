package main

import (
	"fmt"
	"log"
	"strconv"

	"net/http"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
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

func main() {
	go func() { //新监听另一个端口作为pprof http
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	router := gin.Default()
	prefix := router.Group("api/v1")
	prefix.GET("/fb", createFBS)
	router.Run(":8101")
}

func createFBS(c *gin.Context) {
	log.Println(c.Request.URL)
	var fbs []int
	//c.GetQuery()
	v, ok := c.GetQuery("key")
	if !ok {
		c.JSON(404, gin.H{
			"rtn":     404,
			"message": "error",
		})
		return
	}
	n, _ := strconv.Atoi(v)
	for i := 0; i < 1000000; i++ {
		fbs = GetFibonacciSerie(n)
	}

	c.Writer.Write([]byte(fmt.Sprintf("%v", fbs)))
	if len(fbs) != 0 {
		c.JSON(0, gin.H{
			"rtn":     0,
			"message": "success",
		})
		return
	}
}

//url  c.Query("***")
//form c.PostForm("***")
