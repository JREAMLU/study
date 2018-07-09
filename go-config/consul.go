package main

import (
	"fmt"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/encoder/toml"

	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/consul"
)

type Config struct {
	InstanceName string
	DBName       string
	ReadWrite    readwrite
	ReadOnly     readonly
}

type readwrite struct {
	Server   string
	Password string
	Port     string
	UserID   string
	CharSet  string
}

type readonly struct {
	Server   string
	Password string
	Port     string
	UserID   string
	CharSet  string
}

type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

var host Host

func main() {
	enc := toml.NewEncoder()

	consulSource := consul.NewSource(
		consul.WithAddress("10.200.202.35:8500"),
		consul.WithPrefix("/conn/v1/mysql/BGCollector"),
		consul.StripPrefix(true),
		source.WithEncoder(enc),
	)

	// conf := config.NewConfig()
	err := config.Load(consulSource)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	config.Watch()
	// conf.Get("conn", "v1", "mysql", "test").Scan(&host)
	// fmt.Println("++++++++++++: ", host)

	// s := conf.Get("conn", "v1", "mysql", "test").Bytes()
	// fmt.Println("++++++++++++: ", string(s))

	s := config.Get("conn", "v1", "mysql", "BGCollector").Bytes()
	fmt.Println("++++++++++++: ", string(s))
}
