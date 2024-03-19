// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTSceneIntelligenceResultTask = "t_scene_intelligence_result_task"

// TSceneIntelligenceResultTask mapped from table <t_scene_intelligence_result_task>
type TSceneIntelligenceResultTask struct {
	Id             int64     `gorm:"column:id;primaryKey" json:"id"`     // 唯一主键
	StartTime      time.Time `gorm:"column:start_time" json:"startTime"` // 运行时间
	EndTime        time.Time `gorm:"column:end_time" json:"endTime"`
	IntelligenceId int64     `gorm:"column:intelligence_id" json:"intelligenceId"` // 智能场景编号
	ResultId       int64     `gorm:"column:result_id" json:"resultId"`             // 执行结果id
	IsSuccess      int32     `gorm:"column:is_success" json:"isSuccess"`           // 是否成功
	ResultMsg      string    `gorm:"column:result_msg" json:"resultMsg"`           // 结果描述
	TaskId         int64     `gorm:"column:task_id" json:"taskId"`                 // 任务编号
	TaskImg        string    `gorm:"column:task_img" json:"taskImg"`               // 任务图片(产品图片、智能图片、功能图标）
	TaskDesc       string    `gorm:"column:task_desc" json:"taskDesc"`             // 任务描述
	TaskType       int32     `gorm:"column:task_type" json:"taskType"`             // 任务类型（=1 延时 =2 设备执行 =3 场景开关）
	ObjectId       string    `gorm:"column:object_id" json:"objectId"`             // 对象ID（设备Id、场景Id）
	ObjectDesc     string    `gorm:"column:object_desc" json:"objectDesc"`         // 对象的标题或者描述（设备名称、场景名称）
	Functions      string    `gorm:"column:functions" json:"functions"`            // 功能集合json
	FuncKey        string    `gorm:"column:func_key" json:"funcKey"`               // 执行功能Key
	FuncDesc       string    `gorm:"column:func_desc" json:"funcDesc"`             // 冗余：功能描述
	FuncValue      string    `gorm:"column:func_value" json:"funcValue"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`   // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updatedAt"`   // 修改时间
	ProductKey     string    `gorm:"column:product_key" json:"productKey"` // 产品Key
}

// TableName TSceneIntelligenceResultTask's table name
func (*TSceneIntelligenceResultTask) TableName() string {
	return TableNameTSceneIntelligenceResultTask
}
