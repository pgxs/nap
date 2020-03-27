package repository

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEtcdRepository_EtcdClient(t *testing.T) {
	er := EtcdRepository{}
	cli := er.EtcdClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Put(ctx, "/test2", "test000077",clientv3.WithPrevKV())
	cancel()
	if err != nil {
		// handle error!

	}
	assert.NoError(t, err)
	fmt.Println(resp.Header.Revision)
	fmt.Printf("%+v\n", resp.PrevKv)
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	resp1,err:=cli.Get(ctx1,"/test",clientv3.WithPrefix(),clientv3.WithSort(clientv3.SortByKey,clientv3.SortDescend))
	cancel1()
	assert.NoError(t,err)
	fmt.Printf("%+v\n", resp1.Header.Revision)
	fmt.Printf("%+v\n",resp1.Kvs)

}
