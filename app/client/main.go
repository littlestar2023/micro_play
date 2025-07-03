package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	// 配置etcd客户端连接参数
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints: []string{"192.168.1.182:2379"}, // etcd服务地址
		Endpoints:   []string{"47.236.140.65:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,                // 连接超时时间
		Username:    "root",
		Password:    "90!@#RCSdev",
	})
	if err != nil {
		log.Fatalf("连接etcd失败: %v", err)
	}
	defer cli.Close()

	fmt.Println("成功连接到etcd")

	// 创建上下文，设置操作超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 写入数据到etcd
	key := "/demo/key"
	value := "Hello, etcd!"
	_, err = cli.Put(ctx, key, value)
	if err != nil {
		log.Fatalf("写入数据失败: %v", err)
	}

	fmt.Printf("成功写入数据 - Key: %s, Value: %s\n", key, value)

	// 可选：读取数据验证写入结果
	resp, err := cli.Get(ctx, key)
	if err != nil {
		log.Fatalf("读取数据失败: %v", err)
	}

	if len(resp.Kvs) > 0 {
		fmt.Printf("读取数据成功 - Key: %s, Value: %s\n",
			string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	} else {
		fmt.Println("未找到对应Key的数据")
	}
}
