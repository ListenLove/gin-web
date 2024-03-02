# Gin-Web
基于Gin框架的单体应用web项目模板
## 基本功能
1. 使用viper完成项目配置加载
2. 日志初始化
3. 初始化MySQL连接
4. 初始化Redis连接
5. 注册路由
6. 启动服务(优雅关机)

## 中间件
1. 请求头jwt token校验中间件
2. POST请求参数校验结果翻译中间件

## 项目结构
```
├── controller               // 请求处理控制器
├── dao                      // 数据库操作层
      ├── mysql              // mysql数据库操作
      ├── redis              // redis数据库操作
├── logger                   // 日志初始化
├── middleware               // 中间件
├── model                    // 数据模型
├── routes                   // 路由注册
├── service                  // 业务逻辑层
├── pkg                      // 工具包
├── settings                 // 配置文件初始化
├── config.yml               // 配置文件
├── docker-compose.yaml      // docker-compose配置文件
├── go.mod                   // go mod
├── main.go                  // 项目入口
├── LICENSE                  // LICENSE
├── README.md                // 项目说明
```