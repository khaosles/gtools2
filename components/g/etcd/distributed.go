package etcd

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

/*
   @File: distributed.go
   @Author: khaosles
   @Time: 2023/8/6 11:28
   @Desc:
*/

// ExecOnceTask 基于分布式的唯一执行任务
func ExecOnceTask(key string, ttl int, task func()) error {
	ctx := context.Background()
	// 建立session
	session, err := concurrency.NewSession(
		Client,
		concurrency.WithTTL(ttl),
		concurrency.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	mutex := concurrency.NewMutex(session, key)
	// 尝试获取锁
	err = mutex.TryLock(ctx)
	if err != nil {
		return err
	}
	// 执行操作
	task()
	// 执行完成
	// 释放锁
	_ = mutex.Unlock(ctx)
	return nil
}

// ExecAllTask 基于分布式的全部执行任务
func ExecAllTask(key string, ttl int, task func()) error {
	ctx := context.Background()
	// 建立session
	session, err := concurrency.NewSession(
		Client,
		concurrency.WithTTL(ttl),
		concurrency.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	mutex := concurrency.NewMutex(session, key)
	// 尝试获取锁
	err = mutex.Lock(ctx)
	if err != nil {
		return err
	}
	// 执行操作
	task()
	// 执行完成
	// 释放锁
	_ = mutex.Unlock(ctx)
	return nil
}

// OnceTask1 基于分布式的唯一执行任务
func OnceTask1(key string, ttl int64, cb func()) error {
	// 创建一个Etcd租约，设置过期时间
	lease, err := Client.Grant(context.Background(), ttl) // 租约5秒过期
	if err != nil {
		return err
	}
	// 尝试获取锁
	lock, err := Client.Txn(context.Background()).
		If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, fmt.Sprintf("This is a distributed lock by %s.", key), clientv3.WithLease(lease.ID))).
		Commit()
	if err != nil || !lock.Succeeded {
		return fmt.Errorf("failed to acquire distributed lock: %s", key)
	}
	// 自动续约
	_, err = Client.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return err
	}
	// 执行操作
	cb()
	// 执行完成
	// 释放锁
	_ = Del(key)
	return nil
}
