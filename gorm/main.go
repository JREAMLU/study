package main

import (
	"study/core"
	"study/gorm/model"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	core.InitGorm("root:root@tcp(localhost:3306)/plucron?charset=utf8&parseTime=True&loc=Local")
}

func main() {
	cronlist, err := model.Select([]uint64{1, 2})
	if err != nil {

	}
	spew.Dump(cronlist)

	var mc model.Cronlist
	mc.Name = `i am "jream"`
	mc.Type = 1

	model.Insert(mc)
}
