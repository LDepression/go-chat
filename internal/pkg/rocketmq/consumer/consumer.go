/**
 * @Author: lenovo
 * @Description:
 * @File:  consumer
 * @Version: 1.0.0
 * @Date: 2023/05/16 18:59
 */

package consumer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go-chat/internal/global"
	"go-chat/internal/model/chat"
)

func StartConsumer(accountID int64) {
	uID := fmt.Sprintf("accountID:%d", accountID)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{fmt.Sprintf("%s:%d", global.Settings.Rocketmq.Addr, global.Settings.Rocketmq.Port)}),
		consumer.WithGroupName(uID),
	)
	if err != nil {
		fmt.Println("创建消费者失败:", err)
		return
	}

	if err := c.Subscribe(uID, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			global.ChatMap.Send(accountID, chat.ClientSendMsg, msgs[i])
			fmt.Printf("获取到值：%v\n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println("订阅消息失败:", err)
		return
	}

	if err := c.Start(); err != nil {
		fmt.Println("启动消费者失败:", err)
		return
	}

	defer func() {
		if err := c.Shutdown(); err != nil {
			fmt.Println("关闭消费者失败:", err)
		} else {
			fmt.Println("消费者已关闭.")
		}
	}()

	select {}
}
