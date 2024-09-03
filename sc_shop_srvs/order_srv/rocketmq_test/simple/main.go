package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	//创建生产者
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.239.128:9876"}))
	if err != nil {
		panic("生成producer失败")
	}
	//启动
	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}
	//发送消息
	res, err := p.SendSync(context.Background(), primitive.NewMessage("imooc1", []byte("this is imooc1")))
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
	} else {
		fmt.Printf("发送成功: %s\n", res.String())
	}
	//关闭
	if err = p.Shutdown(); err != nil {
		panic("关闭producer失败")
	}

}
