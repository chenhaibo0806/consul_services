//注册中心 Consul Windows安装与使用：
//参考文档：http://www.liuhaihua.cn/archives/679033.html   https://juejin.cn/post/6894518814555832328
//Windows版本：https://www.consul.io/downloads 官网下载win64版本 解压到 D:\consul
//用超管身份在目录 D:\consul 运行cmd ； 执行如下：
//D:\consul> sc.exe create "Consul" binPath="D:\consul\consule.exe agent -dev"
//[SC] CreateService 成功
//然后启动： .\consul.exe agent -dev
//D:\consul> .\consul.exe agent -dev
//==> Starting Consul agent...
//Version: '1.10.0'
//检查Consul服务是否正常使用：在浏览器中输入http://127.0.0.1:8500/ui/dc1/services 查看各个服务的健康状态；
//http://localhost:18001/users/     返回： {"code":200,"msg":"success","data":null}
//http://localhost:18002/orders/    返回： {"code":200,"msg":"get order success","data":null}
//项目目录结构：consul_services项目 下的 orderserver 和 userserver目录
