package main

import (
	"github.com/heroyan/container/cache"
	"log"

	"github.com/heroyan/container/uts"
	nebula "github.com/vesoft-inc/nebula-go"
)

func test()  {
	uts.Uts2()

	client, err := nebula.NewClient("127.0.0.1:3699")
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Connect("username", "password"); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	resp, err := client.Execute("SHOW HOSTS;")
	if err != nil {
		log.Fatal(err)
	}

	if nebula.IsError(resp) {
		log.Printf("ErrorCode: %v, ErrorMsg: %s", resp.GetErrorCode(), resp.GetErrorMsg())
	}
}

func test2()  {
	lru := cache.Constructor(2)
	lru.Put(1, 1)
	log.Println(lru.Result())
	lru.Put(2, 2)
	log.Println(lru.Result())
	log.Println(lru.Get(1))
	lru.Put(3, 3)
	log.Println(lru.Result())
	log.Println(lru.Get(2))
	lru.Put(4, 4)
	log.Println(lru.Result())
	log.Println(lru.Get(1))
	log.Println(lru.Get(3))
	log.Println(lru.Get(4))
}

func main() {
	test2()
}
