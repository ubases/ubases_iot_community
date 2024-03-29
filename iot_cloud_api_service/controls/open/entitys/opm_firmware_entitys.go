// Code generated by sgen.exe,2022-05-03 09:22:14. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"errors"
)

// 增、删、改及查询返回
type OpmFirmwareEntitys struct {
	Id              int64       `json:"id,string"`
	TenantId        string      `json:"tenantId"`
	Name            string      `json:"name"`
	Flag            string      `json:"flag"`
	Type            int32       `json:"type"`
	FlashSize       interface{} `json:"flashSize"`
	UpgradeChannel  int32       `json:"upgradeChannel"`
	UpgradeOvertime int32       `json:"upgradeOvertime"`
	Status          *int32      `json:"status"`
	Remark          string      `json:"remark"`
	CreatedAt       int64       `json:"createdAt"`
	UpdatedBy       int64       `json:"updatedBy"`
	UpdatedAt       int64       `json:"updatedAt"`
	FirmwareKey     string      `json:"firmwareKey"`
	Version         string      `json:"version"`
	VersionDesc     string      `json:"versionDesc"`      //版本文件描述
	IsMust          int32       `json:"isMust"`           //是否必须
	UpgradeMode     int32       `json:"upgradeMode"`      //升级方式
	UpgradeFileName string      `json:"upgradeFileName"`  //升级文件名称
	UpgradeFilePath string      `json:"upgradeFilePath"`  //升级文件地址
	UpgradeFileSize int64       `json:"upgradeFileSize"`  //升级文件尺寸
	UpgradeFileKey  string      `json:"upgradeFileKey"`   //升级文件key
	ProdFilePath    string      `json:"prodFilePath"`     //生成文件地址
	ProdFileSize    int64       `json:"prodFileSize"`     //生成文件尺寸
	ProdFileKey     string      `json:"prodFileKey"`      //生成文件Key
	ProdFileName    string      `json:"prodFileName"`     //生成文件名称
	ProductId       int64       `json:"productId,string"` //产品Id，如果设置了该值，则固件新增之后默认绑定到产品
}

// 新增参数非空检查
func (s *OpmFirmwareEntitys) AddCheck() error {
	if s.Name == "" {
		return errors.New("固件名称不能为空")
	}
	if s.Flag == "" {
		return errors.New("固件标识不能为空")
	}
	if s.Type == 0 {
		return errors.New("固件类型不能为空")
	}
	if s.FlashSize == 0 {
		return errors.New("Flash尺寸不能为空")
	}
	if s.FlashSize == 0 {
		return errors.New("Flash尺寸不能为空")
	}
	if s.TenantId == "" {
		return errors.New("当前用户未指定租户Id")
	}
	return nil
}

// 修改参数非空检查
func (s *OpmFirmwareEntitys) UpdateCheck() error {
	if s.Name == "" {
		return errors.New("固件名称不能为空")
	}
	if s.Flag == "" {
		return errors.New("固件标识不能为空")
	}
	if s.Type == 0 {
		return errors.New("固件类型不能为空")
	}
	if s.FlashSize == 0 {
		return errors.New("Flash尺寸不能为空")
	}
	if s.FlashSize == 0 {
		return errors.New("Flash尺寸不能为空")
	}
	if s.TenantId == "" {
		return errors.New("当前用户未指定租户Id")
	}
	return nil
}

// 查询参数必填检查
func (s *OpmFirmwareQuery) QueryCheck() error {
	if s.Query.TenantId == "" {
		return errors.New("当前用户未指定租户Id")
	}
	return nil
}

// 查询条件
type OpmFirmwareQuery struct {
	Page      int64              `json:"page,omitempty"`
	Limit     int64              `json:"limit,omitempty"`
	Sort      string             `json:"sort,omitempty"`
	SortField string             `json:"sortField,omitempty"`
	SearchKey string             `json:"searchKey,omitempty"`
	Query     *OpmFirmwareFilter `json:"query,omitempty"`
}

// OpmFirmwareFilter，查询条件，字段请根据需要自行增减
type OpmFirmwareFilter struct {
	Id              int64       `json:"id,string,omitempty"`
	TenantId        string      `json:"tenantId,omitempty"`
	Name            string      `json:"name,omitempty"`
	Flag            string      `json:"flag,omitempty"`
	Type            interface{} `json:"type,omitempty"`
	FirmwareKey     string      `json:"firmwareKey"`
	FlashSize       int32       `json:"flashSize,omitempty"`
	UpgradeChannel  int32       `json:"upgradeChannel,omitempty"`
	UpgradeOvertime int32       `json:"upgradeOvertime,omitempty"`
	Status          *int32      `json:"status,omitempty"`
	Remark          string      `json:"remark,omitempty"`
	Version         string      `json:"version,omitempty"`
	CreatedAt       int64       `json:"createdAt,omitempty"`
	UpdatedBy       int64       `json:"updatedBy,omitempty"`
	UpdatedAt       int64       `json:"updatedAt,omitempty"`
}

// 实体转pb对象
func OpmFirmwareFilter_e2pb(src *OpmFirmwareFilter) *proto.OpmFirmware {
	if src == nil {
		return nil
	}
	pbObj := new(proto.OpmFirmware)
	pbObj.TenantId = src.TenantId
	if src.Status != nil {
		pbObj.Status = *src.Status
	} else {
		pbObj.Status = -1
	}
	if src.Name != "" {
		pbObj.Name = src.Name
	}
	if src.FirmwareKey != "" {
		pbObj.FirmwareKey = iotutil.ToString(src.FirmwareKey)
	}
	if src.Type != 0 {
		pbObj.Type = iotutil.ToString(src.Type)
	}
	return pbObj
}

// 实体转pb对象
func OpmFirmware_e2pb(src *OpmFirmwareEntitys) *proto.OpmFirmware {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmFirmware{
		Id:              src.Id,
		Name:            src.Name,
		Flag:            src.Flag,
		Type:            iotutil.ToString(src.Type),
		FlashSize:       iotutil.ToInt32(src.FlashSize),
		UpgradeChannel:  src.UpgradeChannel,
		UpgradeOvertime: src.UpgradeOvertime,
		Status:          1, //*src.Status,
		Remark:          src.Remark,
		Version:         src.Version,
		UpdatedBy:       src.UpdatedBy,
		VersionDesc:     src.VersionDesc,
		IsMust:          src.IsMust,
		UpgradeMode:     src.UpgradeMode,
		UpgradeFileName: src.UpgradeFileName,
		UpgradeFilePath: src.UpgradeFilePath,
		UpgradeFileSize: iotutil.ToInt64(src.UpgradeFileSize),
		UpgradeFileKey:  src.UpgradeFileKey,
		ProdFilePath:    src.ProdFilePath,
		ProdFileSize:    iotutil.ToInt64(src.ProdFileSize),
		ProdFileKey:     src.ProdFileKey,
		ProdFileName:    src.ProdFileName,
		FirmwareKey:     src.FirmwareKey,
	}
	return &pbObj
}

// pb对象转实体
func OpmFirmware_pb2e(src *proto.OpmFirmware) *OpmFirmwareEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OpmFirmwareEntitys{
		Id:              src.Id,
		Name:            src.Name,
		Flag:            src.Flag,
		Type:            iotutil.ToInt32(src.Type),
		FlashSize:       src.FlashSize,
		FirmwareKey:     src.FirmwareKey,
		UpgradeChannel:  src.UpgradeChannel,
		UpgradeOvertime: src.UpgradeOvertime,
		Status:          &src.Status,
		Remark:          src.Remark,
		Version:         src.Version,
		CreatedAt:       src.CreatedAt.AsTime().Unix(),
		UpdatedBy:       src.UpdatedBy,
		UpdatedAt:       src.UpdatedAt.AsTime().Unix(),
		IsMust:          src.IsMust,
	}
	return &entitysObj
}
