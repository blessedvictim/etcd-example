package main

import (
	"context"
	"log"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
	myKey          = "lol-key"
)

func main() {
	cli, _ := client.New(client.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer func() {
		_ = cli.Close()
	}()
	kv := client.NewKV(cli)
	_, err := kv.Put(context.Background(), myKey, "XEXE")
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	watchCh := cli.Watch(ctx, myKey)

	log.Println("watch for", myKey)

	for val := range watchCh {
		if val.Err() != nil {
			panic(err)
		}
		for _, event := range val.Events {
			log.Printf("new event ||| key: %s | value: %s \n", event.Kv.Key, event.Kv.Value)
		}
	}
}
