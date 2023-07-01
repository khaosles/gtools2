package mq

import (
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/khaosles/gtools2/core/config"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: mq.go
   @Author: khaosles
   @Time: 2023/7/1 22:29
   @Desc:
*/

var logger = new(mqLogger)

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

func NewRocketmqProducer(rocketmqCfg *config.Rocketmq) rocketmq.Producer {
	rlog.SetLogger(logger)
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

type mqLogger struct {
	rlog.Logger
}

func (l mqLogger) Debug(msg string, fields map[string]interface{}) {
	glog.Debug(fmt.Sprintf("%s -> %+v", msg, fields))
}

func (l mqLogger) Info(msg string, fields map[string]interface{}) {
	glog.Info(fmt.Sprintf("%s -> %+v", msg, fields))
}

func (l mqLogger) Warning(msg string, fields map[string]interface{}) {
	glog.Warn(fmt.Sprintf("%s -> %+v", msg, fields))
}

func (l mqLogger) Error(msg string, fields map[string]interface{}) {
	glog.Error(fmt.Sprintf("%s -> %+v", msg, fields))
}

func (l mqLogger) Fatal(msg string, fields map[string]interface{}) {
	glog.Panic(fmt.Sprintf("%s -> %+v", msg, fields))
}

func (l mqLogger) Level(msglevel string) {
}

func (l mqLogger) OutputPath(path string) (err error) {
	return nil
}
