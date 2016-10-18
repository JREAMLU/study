package main

import (
	"fmt"
	"study/core"

	"gopkg.in/redis.v4"
)

func init() {
	core.InitRedis("172.16.9.221:6391", "", 3)
}

func main() {
	// var list = make(map[string]interface{})
	// list["name"] = "jream"
	// list["age"] = 25
	// x, err := json.Marshal(list)
	// fmt.Println(string(x), err)
	// err = core.R.Set("name", x, 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// val, err := client.Get("name").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("name", val)
	//
	// val2, err := client.Get("key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exists")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }

	// var list = make(map[string]string)
	// list["name"] = "jream"
	// list["age"] = "23"
	//
	// err := client.HMSet("user:3", list).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// val, err := client.HGetAll("user:3").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("name", val, err)

	var list = make(map[string]interface{})
	list["name"] = "jream"
	list["age"] = 25

	var rz redis.Z
	rz.Score = 1
	rz.Member = "abc"

	io := core.R.ZAdd("ayi", rz)
	fmt.Println(io)

	val := core.R.ZRange("w3ckey", 0, 10)
	fmt.Println("name", val)

}
