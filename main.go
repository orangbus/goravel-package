package main

import (
	mq "goravel/packages/rabbitmq/facades"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/facades"
	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
		}
	}()

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}

		os.Exit(0)
	}()

	// 接口mq的普通消息
	go func() {
		mq.Rabbitmq().Consume()
	}()

	// 2个消息者消费一个交换机的消息
	go func() {
		if err := mq.Rabbitmq().ConsumePublish("logs"); err != nil {
			log.Printf("第一个订阅消费者订阅失败：%s", err.Error())
		}
	}()
	go func() {
		if err := mq.Rabbitmq().ConsumePublish("logs"); err != nil {
			log.Printf("第2个订阅消费者订阅失败：%s", err.Error())
		}
	}()

	// 路由消息
	go func() {
		if err := mq.Rabbitmq().ConsumeRouting("logs_msg", "info"); err != nil {
			log.Printf("第一个路由消费者订阅失败：%s", err.Error())
		}
	}()
	go func() {
		if err := mq.Rabbitmq().ConsumeRouting("logs_msg", "error"); err != nil {
			log.Printf("第2个路由消费者订阅失败：%s", err.Error())
		}
	}()

	// 主题
	go func() {
		// 接受 country1.所有消息.所有消息
		if err := mq.Rabbitmq().ConsumeTopic("city", "country1.*.*"); err != nil {
			log.Printf("第2个路由消费者订阅失败：%s", err.Error())
		}

		// 接受 所有国家.province1.city1
		if err := mq.Rabbitmq().ConsumeTopic("city", "*.province1.city1"); err != nil {
			log.Printf("第2个路由消费者订阅失败：%s", err.Error())
		}
	}()

	select {}
}
