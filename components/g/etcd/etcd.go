package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/khaosles/gtools2/core/config"
	glog "github.com/khaosles/gtools2/core/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

/*
   @File: etcd.go
   @Author: khaosles
   @Time: 2023/7/1 23:35
   @Desc:
*/

var client *clientv3.Client

func InitEtcd(etcdCfg *config.Etcd) {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdCfg.Nodes,
		DialTimeout: time.Duration(etcdCfg.Timeout) * time.Second,
		Username:    etcdCfg.Username,
		Password:    etcdCfg.Password,
	})
	if err != nil {
		glog.Error("etcd connection failed -> ", err.Error())
	}
	glog.Info("etcd connection successful...")
}

func Register(addr, serverName string, ttl int64) error {
	em, err := endpoints.NewManager(client, serverName)
	if err != nil {
		return err
	}
	lease, _ := client.Grant(context.TODO(), ttl)
	err = em.AddEndpoint(
		context.TODO(),
		fmt.Sprintf("%s/%s", serverName, addr),
		endpoints.Endpoint{Addr: addr},
		clientv3.WithLease(lease.ID),
	)
	if err != nil {
		return err
	}
	alive, err := client.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			<-alive
			glog.Debug(addr + " keep alive")
		}
	}()
	return nil
}

func UnRegister(addr, serverName string) error {
	em, err := endpoints.NewManager(client, serverName)
	if err != nil {
		return err
	}
	err = em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serverName, addr))
	if err != nil {
		return err
	}
	return nil
}

func Put(key, value string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	_, err := client.Put(ctx, key, value, opts...)
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string, opts ...clientv3.OpOption) ([]*mvccpb.KeyValue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	resp, err := client.Get(ctx, key, opts...)
	cancel()
	if err != nil {
		return nil, err
	}
	return resp.Kvs, nil
}

func Del(key string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	_, err := client.Delete(ctx, key, opts...)
	cancel()
	if err != nil {
		return err
	}
	return nil
}
