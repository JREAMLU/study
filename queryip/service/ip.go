package service

import (
	"fmt"
	"log"
	"time"

	"git.corp.plu.cn/plugo/infrastructure/extensions"

	"github.com/JREAMLU/study/queryip/data/redis"
	"github.com/wangtuanjie/ip17mon"
)

var (
	location     *ip17mon.LocationInfo
	combineInfo  string
	combineInfos []string
)

// SetIPS save ips
func SetIPS(ips []string) error {
	return redis.NewIPSHandler().SaveBatch(ips)
}

// BatchIPS batch save ips
func BatchIPS(ips []uint32) error {
	beginTime := time.Now().UnixNano()
	for _, ipUint := range ips {
		combineInfos = append(combineInfos, gos(ipUint))
	}
	endTime := time.Now().UnixNano()
	takeTime := endTime - beginTime
	log.Printf("IP append time: %vms , %vns", takeTime/1000000, takeTime)

	return SetIPS(combineInfos)
}

func gos(ipUint uint32) string {
	location = ip17mon.FindByUint(ipUint)
	combineInfo = fmt.Sprintf(
		"%s||%s||%s||%s||%s",
		extensions.Int2Ip(int64(ipUint)),
		location.Country,
		location.Region,
		location.City,
		location.Isp,
	)

	return combineInfo
}
