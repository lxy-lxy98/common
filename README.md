# common
golang服务，里面存在很多可以复用的多个模块，因此开发该项目统一维护，代码复用。

1. codes: 常用byte与float32相互转换，byte与float32数组互相转换，[]byte与string互相转换。
2.  math:Decimal保留浮点数后位数，Equal判断浮点数是否相同
3.  model:结果切片方法以及定时器
4.  reflect:利用反射的结构体数据万能填充方法
5. pprof:性能分析方法{
  pprof/file:使用方法
  1. go build prof.go
  2. ./prof
  3. go tool pprof prof cpu.prof

  pprof/http:使用方法
  1. go run 启动一个http服务后
  2. 重启一个终端
  3. go tool pprof http://localhost:8081/debug/pprof/profile?seconds=10   默认30秒，此命令监听端口10秒
  4. 浏览器调用
}

6. gin:开源的Gin框架