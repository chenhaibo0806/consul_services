package main

import (
	"github.com/micro/go-micro/registry"//注意这些地址变了
	"github.com/micro/go-micro/web"//注意这些地址变了
	"github.com/micro/go-plugins/registry/consul"//注意这些地址变了
	"consul_services/userserver/routers"
)

var consulReg registry.Registry

func  init()  {
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	consulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
}

func main() {
	//初始化路由
	ginRouter := routers.InitRouters()

	//注册服务
	microService:= web.NewService(
		web.Name("userserver"),
		//web.RegisterTTL(time.Second*30),//设置注册服务的过期时间
		//web.RegisterInterval(time.Second*20),//设置间隔多久再次注册服务
		web.Address(":18001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)

	microService.Run()
}