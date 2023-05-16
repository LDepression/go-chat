/**
 * @Author: lenovo
 * @Description:
 * @File:  producer
 * @Version: 1.0.0
 * @Date: 2023/05/16 19:00
 */

package producer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go-chat/internal/global"
)

func SendMsgToMQ(mID uint) {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{fmt.Sprintf("%s:%d", global.Settings.Rocketmq.Addr, global.Settings.Rocketmq.Port)}))
	if err != nil {
		panic("生成producer失败")
	}
	if err = p.Start(); err != nil {
		panic(err)
	}
	uID := fmt.Sprintf("accountID:%d", mID)
	SystemMsg := "welcome to chatroom!"
	res, err := p.SendSync(context.Background(), primitive.NewMessage(uID, []byte(SystemMsg)))
	if err != nil {
		fmt.Println("发送失败", err)
	} else {
		fmt.Println("发送成功,res:", res.String())
	}
	if err := p.Shutdown(); err != nil {
		panic(err)
	}
}
