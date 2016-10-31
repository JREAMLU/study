package main

import (
	"fmt"
	"net/url"
)

func main() {
	eg2()
}

func eg2() {
	s := `v=1&_v=mithril&cid=1028169706.1473345397&tid=100001-1&t=event&dl=http%3A%2F%2Fwww.longzhu.com%2F&dr=https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DFyq4ek-wTrBRP0Fumr4Qo4zKFRtKZF1A4iA5Aio2cdpLNm72zP5ewhTf4MSMlupg%26wd%3D%26eqid%3Dc35ca6a700041e5e0000000558108d0f&ul=zh-cn&de=UTF-8&dt=%E9%BE%99%E7%8F%A0%E7%9B%B4%E6%92%ADlongzhu.com-%E6%B8%B8%E6%88%8F%E7%9B%B4%E6%92%AD%E5%B9%B3%E5%8F%B0&sd=24-bit&sr=1920x1080&vp=1903x944&je=0&fl=21.9%20r9&uid=49088152&ec=Page%20Visibility&ea=change&el=hidden%20%3D%3E%20visible&ev=7&_s=4&cd10=2&_=1445495818`
	m, _ := url.ParseQuery(s)
	fmt.Println(m)
}

func eg1() {
	s := "http://mattock.plu.cn/test?name=kk&age=14&word=aa#ff"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("host: ", u.Host)
	fmt.Println("path: ", u.Path)
	fmt.Println("RawPath: ", u.RawPath)
	fmt.Println("rawQuery: ", u.RawQuery)

}
