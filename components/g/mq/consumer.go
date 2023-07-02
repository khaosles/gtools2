package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/khaosles/gtools2/core/config"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: consumer.go
   @Author: khaosles
   @Time: 2023/7/2 09:33
   @Desc:
*/

func NewRocketmqConsumer(rocketmqCfg *config.Rocketmq, groupName string) rocketmq.PushConsumer {
	rlog.SetLogger(logger)
	pushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(rocketmqCfg.NameServer),
		consumer.WithGroupName(groupName),
		consumer.WithRetry(rocketmqCfg.Retry),
	)
	if err != nil {
		glog.Error("rocketmq consumer connection failed -> ", err.Error())
	}
	glog.Info("rocketmq consumer connection successful...")
	return pushConsumer
}
