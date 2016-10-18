package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// var data = `
// redisConfig: [
//     [
//         name: cache
//         poolSize: 1
//         connect: [
//             ip: 172.16.9.221:6391
//             db: 0
//             master: true
//         ]
//     ]
//     [
//         name: feed
//         poolSize: 1
//         connect: [
//             ip: 172.16.9.221:6391
//             db: 0
//             master: true
//         ]
//     ]
// ]
// `

// type T struct {
// 	A string
// 	B struct {
// 		RenamedC int   `yaml:"c"`
// 		D        []int `yaml:",flow"`
// 	}
// }

// type T struct {
// 	RedisConfig []struct {
// 		Name     string `yaml:"name"`
// 		PoolSize int    `yaml:"poolSize"`
// 		Connect  []struct {
// 			IP     string `yaml:"ip"`
// 			Db     int    `yaml:"db"`
// 			Master bool   `yaml:"master"`
// 		} `yaml:"connect"`
// 	} `yaml:"redisConfig"`
// }

type T struct {
	RedisConf []RedisConfig `yaml:"redisConfig"`
}

type RedisConfig struct {
	Name     string    `yaml:"name"`
	PoolSize int       `yaml:"poolSize"`
	Connects []Connect `yaml:"connect"`
}

type Connect struct {
	IP     string `yaml:"ip"`
	Db     int    `yaml:"db"`
	Master bool   `yaml:"master"`
}

func read(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func main() {
	t := T{}

	data := read("./redis.yml")

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

}
