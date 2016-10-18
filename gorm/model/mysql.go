package model

import (
	"study/core"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var core.X *gorm.DB
//
// func init() {
// 	var err error
// 	core.X, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/plucron?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 	}
// 	core.X.SingularTable(true)
// 	core.X.LogMode(true)
// }

func Insert(cron Cronlist) (uint64, error) {
	res := core.X.Create(&cron)
	if res.Error != nil {
		return 0, res.Error
	}
	return cron.ID, nil
}

func Update(cron Cronlist, id []uint64) error {
	res := core.X.Table("cronlist").Where("id IN (?)", id).Updates(cron)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func Delete(id []uint64) error {
	res := core.X.Delete(Cronlist{}, "id IN (?)", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func Select(id []uint64) (cronlist Cronlist, err error) {
	sql := `
SELECT	 id, name, type
FROM  	  cronlist 
WHERE 	id IN (?)
`
	res := core.X.Raw(sql, id).Scan(&cronlist)
	if res.Error != nil {
		return cronlist, res.Error
	}
	return cronlist, nil
}

//transaction
func Transact() error {
	tx := core.X.Begin()
	cronlist := Cronlist{
		Name: "Iversion",
		Type: 2,
	}
	res := core.X.Create(&cronlist)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	res = core.X.Delete(Cronlist{}, "id IN (?)", []uint64{361})
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	tx.Commit()
	return nil
}
