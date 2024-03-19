// Code generated by sgen.exe,2022-05-31 12:02:53. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"
)

// 新增和修改
type OemAppDefMenuEntitys struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	MenuKey  string `json:"menuKey"`
	Position int32  `json:"position"`
	DefImage string `json:"defImage"`
	SelImage string `json:"selImage"`
	Required int32  `json:"required"`
}

// //增、删、改及查询返回
// type OemAppDefMenuEntitys struct {
//     Id int64  `json:"id,omitempty"`
//     Name string `json:"name,omitempty"`
//     MenuKey string `json:"menuKey,omitempty"`
//     Position int32  `json:"position,omitempty"`
//     DefImage string `json:"defImage,omitempty"`
//     SelImage string `json:"selImage,omitempty"`
//     Required int32  `json:"required,omitempty"`
//     Status int32  `json:"status,omitempty"`
//     CreatedBy int64  `json:"createdBy,omitempty"`
//     UpdatedBy int64  `json:"updatedBy,omitempty"`
//     CreatedAt time.Time `json:"createdAt,omitempty"`
//     UpdatedAt time.Time `json:"updatedAt,omitempty"`
//     DeletedAt time.Time `json:"deletedAt,omitempty"`
// }

// 新增参数非空检查
func (s *OemAppDefMenuEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *OemAppDefMenuEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*OemAppDefMenuQuery) QueryCheck() error {
	return nil
}

// 查询条件
type OemAppDefMenuQuery struct {
	Page  int64  `json:"pageNum" form:"pageNum"`
	Limit int64  `json:"pageSize" form:"pageSize"`
	Name  string `json:"name" form:"name"`
}

// OemAppDefMenuFilter，查询条件，字段请根据需要自行增减
type OemAppDefMenuFilter struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	MenuKey   string    `json:"menuKey,omitempty"`
	Position  int32     `json:"position,omitempty"`
	DefImage  string    `json:"defImage,omitempty"`
	SelImage  string    `json:"selImage,omitempty"`
	Required  int32     `json:"required,omitempty"`
	Status    int32     `json:"status,omitempty"`
	CreatedBy int64     `json:"createdBy,omitempty"`
	UpdatedBy int64     `json:"updatedBy,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

// 实体转pb对象
func OemAppDefMenu_e2pb(src *OemAppDefMenuEntitys) *proto.OemAppDefMenu {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppDefMenu{
		Id:       iotutil.ToInt64(src.Id),
		Name:     src.Name,
		MenuKey:  src.MenuKey,
		Position: src.Position,
		DefImage: src.DefImage,
		SelImage: src.SelImage,
		Required: src.Required,
	}
	return &pbObj
}

// pb对象转实体
func OemAppDefMenu_pb2e(src *proto.OemAppDefMenu) *OemAppDefMenuEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OemAppDefMenuEntitys{
		Id:       iotutil.ToString(src.Id),
		Name:     src.Name,
		MenuKey:  src.MenuKey,
		Position: src.Position,
		DefImage: src.DefImage,
		SelImage: src.SelImage,
		Required: src.Required,
	}
	return &entitysObj
}
