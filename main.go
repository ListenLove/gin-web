package main

import (
	"context"
	"fmt"
	"gin-web/dao/mysql"
	"gin-web/dao/redis"
	"gin-web/logger"
	"gin-web/routes"
	"gin-web/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 1. 配置初始化
	if err := settings.Init(); err != nil {
		fmt.Printf("配置初始化失败: %v\n", err)
		return
	}
	// 2. 日志初始化
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("日志初始化失败: %v\n", err)
		return
	}
	defer zap.L().Sync() // 延迟关闭日志
	zap.L().Debug("logger init success")
	// 3. MySQL数据库初始化
	if err := mysql.Init(); err != nil {
		fmt.Printf("MySQL数据库初始化失败: %v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// 4. Redis数据库初始化
	if err := redis.Init(); err != nil {
		fmt.Printf("Redis数据库初始化失败: %v\n", err)
		return
	}
	defer redis.Close() // 程序退出关闭数据库连接

	// 5. 注册路由
	r := routes.Init()
	// 6. 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("监听并启动Gin服务失败: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("正在关闭服务中...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("服务关闭失败: ", zap.Error(err))
	}

	zap.L().Info("服务退出")
}
