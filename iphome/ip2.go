package main

import (
	"encoding/json"
	"fmt"

	"git.corp.plu.cn/liutaihua/ipHome/ipquery"
)

func init() {
	ipquery.LoadIpdat("./ip.dat")
}

func main() {
	ips := []string{"115.231.106.232", "119.75.218.70"}
	res, _ := ipquery.Query(ips)
	result, _ := json.Marshal(res)
	fmt.Println(string(result))
}

//IpToJson
func IpToJson() {

}
