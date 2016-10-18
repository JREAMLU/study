package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/wangtuanjie/ip17mon"
)

var ipfile string

func init() {
	flag.StringVar(&ipfile, "file", "./ip.dat", "port for listener")

	flag.Parse()
	if err := ip17mon.Init(ipfile); err != nil {
		panic(err)
	}
}

func main() {
	ips := "127.0.0.1,192.168.1.1"
	ip := GetinIphone(ips)
	fmt.Println(ip)
}

func GetinIphone(ips string) string {
	var (
		err        error
		res        = map[string]interface{}{"code": 0}
		ipList     []string
		ipinfoList []map[string]string
		location   *ip17mon.LocationInfo
		loc_info   string
	)

	if ips == "" {
		res["code"] = 11111
		res["data"] = map[string]string{}
		res["message"] = "not found ips argument"

		result, _ := json.Marshal(res)
		return string(result)
	}

	ipinfoList = make([]map[string]string, 0)
	ipList = strings.Split(ips, ",")
	for _, ip := range ipList {
		if location, err = ip17mon.Find(ip); err != nil {
			//log
		}
		if location != nil {
			loc_info = fmt.Sprintf("%s||%s||%s||%s", location.Country, location.Region, location.City, location.Isp)
		} else {
			loc_info = "NA||NA||NA||NA"
		}
		ipinfoList = append(ipinfoList, map[string]string{ip: loc_info})
	}
	res["data"] = ipinfoList
	res["message"] = "success"

	result, _ := json.Marshal(res)
	return string(result)

}
