package main

import (
	"fmt"
	"plucron-server/model"
)

func ta() {
	var cron model.Cron
	cron.Id = 3
	cron.Status = "procsessing"
	cron.Pid = 100
	cron.Last_execute_time = 132313
	num, err := cron.UpdateCronServer()

	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("num: ", num)
}

func tb() {
	var cron model.Cron
	cron.Id = 3
	cron.Status = "procsessing"
	cron.Pid = 123
	num, err := cron.UpdateCronServer()

	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("num: ", num)
}

func main() {
	tb()
}
