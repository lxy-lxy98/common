# common
golang服务，里面存在很多可以复用的多个模块，因此开发该项目统一维护，代码复用。

* codes: 常用byte与float32相互转换，byte与float32数组互相转换，[]byte与string互相转换。

* math:Decimal保留浮点数后位数，Equal判断浮点数是否相同

* model:通用的模型模块
* reflect:利用反射的结构体数据万能填充方法
* pprof:性能分析方法{
  pprof/file:使用方法
  ###  go build prof.go
  ### ./prof
  ### go tool pprof prof cpu.prof

  pprof/http:使用方法
  ### go run 启动一个http服务后
  ### 重启一个终端
  ### go tool pprof http://localhost:8081/debug/pprof/profile?seconds=10   默认30秒，此命令监听端口10秒
  ### 浏览器调用
}

* gin:开源的Gin框架