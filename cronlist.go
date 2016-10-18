package main

import (
	"encoding/json"
	"fmt"
	"plucron-server/model"
)

type Money struct {
	Spec   string `json:"spec"`
	Script string `json:"script"`
	Data   model.Cronlist
}

func main() {
	params := `{"spec":"*/2 * * * * *","script":"ping 127.0.0.1 >> text1.txt","data":{"name":"lbj","type":13}}`
	var money Money
	if err := json.Unmarshal([]byte(params), &money); err != nil {
		fmt.Println("err: ", err)
	}
	var cronlist model.Cronlist
	id, err := cronlist.AddList(money.Data)
	if err != nil {
		fmt.Println("add failed:", err)
	}
	fmt.Println("id: ", id)
}
