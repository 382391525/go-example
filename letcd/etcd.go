package letcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"sync"
	"time"
)

func EtcdGetPut() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer client.Close()
	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = client.Put(ctx, "e3w_test/fzf", "123456")
	cancel()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, "e3w_test/fzf")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func EtcdWatch() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer client.Close()
	rch := client.Watch(context.Background(), "e3w_test/fzf")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func EtcdSyncLock(id int, lockName string) {
	client := NewEtcdClient()
	defer client.Close()
	s, err := concurrency.NewSession(client)
	if err != nil {
		return
	}
	defer s.Close()
	// 创建一个etcd locker
	locker := concurrency.NewLocker(s, lockName)

	log.Printf("id：%v 尝试获取锁%v", id, lockName)
	locker.Lock()
	log.Printf("id:%v取得锁%v", id, lockName)

	// 模拟业务耗时
	time.Sleep(time.Millisecond * 300)

	locker.Unlock()
	log.Printf("id:%v释放锁%v", id, lockName)
}

func NewEtcdClient() *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("err", err)
		return nil
	}
	return client
}

func EtcdSyncLockTest() {
	lockName := "locker-test"
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			EtcdSyncLock(id, lockName)
		}(i)
	}
	wg.Wait()
	fmt.Println("etcd lock done")
}
