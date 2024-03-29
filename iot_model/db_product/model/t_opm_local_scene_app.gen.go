// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOpmLocalSceneApp = "t_opm_local_scene_app"

// TOpmLocalSceneApp mapped from table <t_opm_local_scene_app>
type TOpmLocalSceneApp struct {
	Id        int64     `gorm:"column:id;primaryKey" json:"id"`                        // 主键
	AppId     int64     `gorm:"column:app_id;not null" json:"appId"`                   // APPId
	AppKey    string    `gorm:"column:app_key" json:"appKey"`                          // APPKey
	SceneId   int64     `gorm:"column:scene_id;not null" json:"sceneId"`               // 场景Id
	CreatedBy int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
}

// TableName TOpmLocalSceneApp's table name
func (*TOpmLocalSceneApp) TableName() string {
	return TableNameTOpmLocalSceneApp
}
