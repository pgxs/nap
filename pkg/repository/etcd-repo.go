package repository

import (
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
	//etcdClientOnce.Do(func() {
	//	//get config
	//	// 建立一个客户端
	//	if client, err = clientv3.New(config); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	client =
	//})
	return client
}
