package main

import (
	"fmt"
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

func pl()  {
	//cache.Pailie([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20})
	cache.Pailie([]int{1,2,3,4,5,6,7,8,9,10,11})
}

func lock()  {
	fmt.Println(cache.Test([]string{"0201","0101","0102","1212","2002"}, "0202"))
	fmt.Println(cache.Test([]string{"8888"}, "0009"))
	fmt.Println(cache.Test([]string{"0000"}, "8888"))
	fmt.Println(cache.Test([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
}

func main() {
	//pl()
	//fmt.Println(cache.Test2(100))
	//fmt.Println(cache.Test2(2))
	//fmt.Println(cache.Test2(4))
	cache.TestCoin()
}
