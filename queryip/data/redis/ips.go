package redis

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"git.corp.plu.cn/plugo/infrastructure/redis"
)

const (
	// REDISINSTANCENAME server name
	REDISINSTANCENAME = "cache"
	// REDISKEYTPLIPS redis key tpl of ips
	REDISKEYTPLIPS = "ips"
)

// IPSHandler ips redis handler
type IPSHandler struct {
	ips redis.RedisSet
}

// NewIPSHandler new ips handler
func NewIPSHandler() *IPSHandler {
	return &IPSHandler{
		ips: redis.NewRedisSet(REDISINSTANCENAME, "%v"),
	}
}

// Save save ips
func (handler *IPSHandler) Save(ip string) error {
	spew.Dump(handler.ips.ServerName)
	ok, err := handler.ips.Add(REDISKEYTPLIPS, ip)
	fmt.Println(ok, err)
	if err != nil {
		return err
	}

	return nil
}

// SaveBatch save ips batch
func (handler *IPSHandler) SaveBatch(ips []string) error {
	ok, err := handler.ips.Add(REDISKEYTPLIPS, ips...)
	fmt.Println(ok, err)
	if err != nil {
		return err
	}

	return nil
}
