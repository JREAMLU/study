package model

// type Model struct {
// 	ID        uint   `gorm:"primary_key;column:id"`
// 	CreatedAt uint64 `gorm:"column:added_time"`
// 	UpdateAt  uint64 `gorm:"column:updated_time"`
// }
//
// type Cron struct {
// 	Model
// 	Prehook           string `gorm:"column:prehook;size:255"`
// 	Spec              string `gorm:"column:spec;size:255"`
// 	Script            string `gorm:"column:script;size:255"`
// 	Posthook          string `gorm:"column:posthook;size:255"`
// 	Remark            string `gorm:"column:remark;size:255"`
// 	Name              string `gorm:"column:name;size:30"`
// 	Status            string `gorm:"column:status;size:255"`
// 	Pid               uint64 `gorm:"column:pid"`
// 	Last_execute_time uint64 `gorm:"column:last_execute_time"`
// 	Version           uint64 `gorm:"column:version"`
// 	Actions           uint64 `gorm:"column:actions"`
// }

type Cronlist struct {
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"column:name"`
	Type uint64 `gorm:"column:type"`
}
