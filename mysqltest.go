package main

import (
	"flag"
	"fmt"

	"github.com/astaxie/beego/orm"

	"git.corp.plu.cn/plugo/infrastructure/mysql"
)

var (
	mysqlConfig string
)

type Entity struct {
	Id int `orm:"column(Id);pk;auto"`
}
type CronModel struct {
	Entity
	Prehook           string `orm:"column(prehook)"`
	Spec              string `orm:"column(spec)"`
	Script            string `orm:"column(script)"`
	Posthook          string `orm:"column(posthook)"`
	Remark            string `orm:"column(remark)"`
	Name              string `orm:"column(name)"`
	Status            string `orm:"column(status)"`
	Pid               int    `orm:"column(pid)"`
	Last_execute_time int    `orm:"column(last_execute_time)"`
	Added_time        int    `orm:"column(added_time)"`
	Updated_time      int    `orm:"column(updated_time)"`
}

func init() {
	flag.StringVar(&mysqlConfig, "mysql", "./connectionStrings.test.config", "redis config")
	flag.Parse()
	err := mysql.LoadMysqlSettings(mysqlConfig)
	if err != nil {
		fmt.Println("mysql connect err: ", err)
		panic(err)
	}

	orm.RegisterModel(new(CronModel))
}

func GetCronByPid(id int) ([]CronModel, error) {
	o := mysql.GetMysqlDefaultDb()
	var results []CronModel
	_, err := o.Raw("select * from cron where pid=?", id).QueryRows(&results)
	return results, err
}

func main() {
	result, err := GetCronByPid(1)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("result: ", result)
}

// type Money struct {
// 	Spec   string `json:"spec"`
// 	Script string `json:"script"`
// 	Data   model.Cronlist
// }

// func main() {
// 	params := `{"spec":"*/2 * * * * *","script":"ping 127.0.0.1 >> text1.txt","data":{"name":"lbj","type":13}}`
// 	var money Money
// 	if err := json.Unmarshal([]byte(params), &money); err != nil {
// 		fmt.Println("err: ", err)
// 	}
// 	var cronlist model.Cronlist
// 	id, err := cronlist.AddList(money.Data)
// 	if err != nil {
// 		fmt.Println("add failed:", err)
// 	}
// 	fmt.Println("id: ", id)
// }
