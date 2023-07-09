package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/khaosles/gtools2/core/config"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: producer.go
   @Author: khaosles
   @Time: 2023/7/2 09:33
   @Desc:
*/

func NewRocketmqProducer(rocketmqCfg *config.Rocketmq) rocketmq.Producer {
	rlog.SetLogLevel("error")
	newProducer, err := rocketmq.NewProducer(
		producer.WithNameServer(rocketmqCfg.NameServer),
		producer.WithRetry(rocketmqCfg.Retry),
	)
	if err != nil {
		glog.Error("rocketmq producer connection failed -> ", err.Error())
	}
	glog.Info("rocketmq producer connection successful...")
	return newProducer
}
