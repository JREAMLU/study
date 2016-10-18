package main

import (
	"fmt"
	"study/core"
)

func request1(str string) int {
	fmt.Println("1" + str)
	return 1
}

func request2() interface{} {
	fmt.Println("2")
	return nil
}

func main() {
	var addFunc core.MultiAddFunc
	addFunc = append(addFunc, core.AddFunc{Name: "a", Handler: request1, Params: []interface{}{"a"}})
	addFunc = append(addFunc, core.AddFunc{Name: "b", Handler: request2})

	res, err := core.GoAsyncRequest(addFunc, 2)
	fmt.Println(res, err)
	fmt.Println(res["a"][0])
	fmt.Println(res["b"][0])
}

//e.g
/*
func main() {
	// 新建一个async对象
	async := async.New()

	// 添加request1异步请求,第一个参数为该异步请求的唯一logo,第二个参数为异步完成后的回调函数,回调参数类型为func()interface{}
	async.Add("a", request1, "a")
	// 添加request2异步请求
	async.Add("b", request2)

	// 执行
	if chans, ok := async.Run(); ok {
		// 将数据从通道中取回,取回的值是一个map[string]interface{}类型,key为async.Add()时添加的logo,interface{}为该logo回调函数返回的结果
		res := <-chans
		// 这里最好判断下是否所有的异步请求都已经执行成功
		if len(res) == 2 {
			for k, v := range res {
				//do something
				fmt.Println("=======:", k, v)
			}
		} else {
			log.Println("async not execution all task")
		}
	}

	// 清除掉本次操作的所以数据,方便后续继续使用async对象
	async.Clean()
}
*/
