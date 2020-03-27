package repository

import (
	"log"
	"sync"

	"github.com/coreos/etcd/clientv3"
)

//EtcdRepository etcd base repository
type EtcdRepository struct {
}

var (
	client         *clientv3.Client
	etcdClientOnce sync.Once
)

func (er *EtcdRepository) EtcdClient() *clientv3.Client {
	etcdClientOnce.Do(func() {
		//get config
		config := clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}}
		// 建立一个客户端
		if nClient, err := clientv3.New(config); err == nil {
			client = nClient
		} else {
			log.Fatalln(err)
		}
	})
	return client
}
