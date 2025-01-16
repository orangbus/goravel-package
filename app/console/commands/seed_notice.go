package commands

import (
	"fmt"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/orangbus/rabbitmq/facades"
	"log"
	"math/rand"
	"time"
)

type SeedNotice struct {
}

// go run . artisan mq:notice
func (receiver *SeedNotice) Signature() string {
	return "mq:notice"
}

// Description The console command description.
func (receiver *SeedNotice) Description() string {
	return "生成测试的mq消息"
}

// Extend The console command extend.
func (receiver *SeedNotice) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *SeedNotice) Handle(ctx console.Context) error {
	type_msg := false
	type_publish := false
	type_routing := false
	type_top := true
	// 发送mq的普通消息
	if type_msg {
		go func() {
			var total int64
			for {
				total++
				err := facades.Rabbitmq().Msg(fmt.Sprintf("%d 测试mq消息", total))
				if err != nil {
					log.Printf("rabbitmq 普通消息发送失败：%s", err.Error())
				}
				log.Printf("rabbitmq 普通消息发送成功：%d", total)
				time.Sleep(time.Second)
			}
		}()
	}

	// 订阅模式：这里面的消息可以有多个消费者消费
	if type_publish {
		go func() {
			var total2 int64
			for {
				total2++
				err := facades.Rabbitmq().Publish("logs", fmt.Sprintf("%d 订阅消息", total2))
				if err != nil {
					log.Printf("rabbitmq 订阅消息发送失败：%s", err.Error())
				}
				log.Printf("rabbitmq 订阅消息发送成功：%d", total2)
				if total2%2 == 0 {
					time.Sleep(time.Second)
				}
			}
		}()
	}

	// 路由模式：
	if type_routing {
		go func() {
			key := "info"
			var total3 int64
			for {
				total3++
				if total3%2 == 0 {
					key = "info"
				} else {
					key = "error"
				}
				log.Printf("total3:%d -> %d -> %s", total3, total3%2, key)
				err := facades.Rabbitmq().Routing("logs_msg", key, fmt.Sprintf("%d 【key:%s】路由消息", total3, key))
				if err != nil {
					log.Printf("rabbitmq 路由消息发送失败：%s", err.Error())
				}
				log.Printf("[key:%s]rabbitmq 路由消息发送成功：%d", key, total3)
				time.Sleep(time.Second)
			}
		}()
	}

	// 主题模式
	if type_top {
		go func() {
			var total4 int64
			country := []string{"country1", "country2"}
			province := []string{"province1", "province2"}
			city := []string{"city1", "city2"}

			key := fmt.Sprintf("%s.%s.%s", randomVal(country), randomVal(province), randomVal(city))
			for {
				total4++
				err := facades.Rabbitmq().Topic("city", key, fmt.Sprintf("%d 【key:%s】主题消息", total4, key))
				if err != nil {
					log.Printf("rabbitmq 主题消息发送失败：%s", err.Error())
				}
				log.Printf("[key:%s]rabbitmq 主题消息发送成功：%d", key, total4)
				time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	select {}
	return nil
}

func randomVal(list []string) string {
	randomIndex := rand.Intn(len(list))
	return list[randomIndex]
}
