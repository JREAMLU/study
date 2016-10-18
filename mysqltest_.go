package main

import (
	"encoding/json"
	"fmt"
	"os"
	"plucron-server/config"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	alias    = config.DB_ALIAS_M
	driver   = config.DB_DRIVER_M
	username = config.DB_USERNAME_M
	passowrd = config.DB_PASSWORD_M
	dbname   = config.DB_NAME_M
	charset  = config.DB_CHARSET_M
	maxidle  = config.DB_MAXIDLE_M
	debug    = config.DB_DEBUG
)

type Crontest struct {
	Id   int
	Name string
	Type int `orm:"column(type)"`
}

func init() {
	// set default database
	orm.RegisterDataBase(alias, driver, username+":"+passowrd+"@/"+dbname+"?charset="+charset, 30)
	// orm.RegisterDataBase(alias, driver, username+":@/"+database+"?charset="+charset, maxIdle)

	orm.Debug = debug

	// register model
	orm.RegisterModel(new(Crontest))

	// create table
	// orm.RunSyncdb("default", false, true)
}

/**
 *	@auther		jream.lu
 *	@intro		封装原生查询语句
 *	@return 	slice maps
 */
func SelectMap(params []interface{}, sql string) (maps []orm.Params, total int64) {
	o := orm.NewOrm()
	num, err := o.Raw(sql, params).Values(&maps)
	if err != nil {
		//log
		fmt.Println(err)
	}
	return maps, num
}

/**
 *	@auther		jream.lu
 *	@intro		封装原生查询语句
 *	@return 	slice lists
 */
func SelectSlice(params []interface{}, sql string) (lists []orm.ParamsList, total int64) {
	o := orm.NewOrm()
	num, err := o.Raw(sql, params).ValuesList(&lists)
	if err != nil {
		//log
		fmt.Println(err)
	}
	return lists, num
}

/**
 *	@auther		jream.lu
 *	@intro		封装原生更新语句
 *	@return 	int64, err
 */
func Update(params []interface{}, sql string) int64 {
	o := orm.NewOrm()
	res, err := o.Raw(sql, params).Exec()
	if err != nil {
		fmt.Println(err)
	}
	num, _ := res.RowsAffected()
	fmt.Println("mysql row affected nums: ", num)
	return num
}

/**
 *	@auther		jream.lu
 *	@intro		封装原生批量更新语句
 *	@return 	slice lists
 */
func UpdateBatch() {}

//exapmle=====================================================
//SELECT
func GetCronList() {
	params := []interface{}{1, 2, 3, 1}
	sql := `
SELECT	id, name, type as ptype 
FROM 	 crontest WHERE id IN (?, ?, ?) 
AND 	   type = ? 
`
	// result, _ := SelectMap(params, sql)
	result, num := SelectMap(params, sql)
	for _, v := range result {
		fmt.Println("id:", v["id"])
		fmt.Println("name:", v["name"])
		fmt.Println("type:", v["ptype"])
	}
	fmt.Println(result, num)
	fmt.Println("================map 到json str=====================")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(result)
}

func Add() {
	o := orm.NewOrm()
	var crontest Crontest
	crontest.Name = "哈登"
	crontest.Type = 3
	id, err := o.Insert(&crontest)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("id: ", id)
}

func UpdateCron() {
	params := []interface{}{4, 5}
	sql := `
UPDATE 	crontest
SET 		 type = ?
WHERE 	 id = ?
`
	num := Update(params, sql)
	fmt.Println("num: ", num)
}

func DeleteCron() {
	params := []interface{}{5}
	sql := `
DELETE FROM crontest
WHERE id = ?
`
	num := Update(params, sql)
	fmt.Println("num: ", num)
}

func main() {
	DeleteCron()
}
