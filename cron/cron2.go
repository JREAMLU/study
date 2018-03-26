package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("每一分钟执行:", time.Now().Format("2006-01-02 15:04:05"))
		//ps
	})
	c.Start()
	fmt.Println("检查: ", c.Entries())
	for {
	}

}
