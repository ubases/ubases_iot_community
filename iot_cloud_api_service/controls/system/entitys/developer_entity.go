package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
	"errors"
	"strconv"
)

// 列表查询条件
type DeveloperEntitysListReq struct {
	PageNum     uint64 `form:"pageNum,omitempty"`
	PageSize    uint64 `form:"pageSize,omitempty"`
	Account     string `form:"account,omitempty"`
	CompanyId   string `form:"companyId,omitempty"`
	CompanyName string `form:"companyName,omitempty"`
	Status      int32  `form:"status,omitempty"`
	//1.增加列表字段显示和筛选项：账号类型，账号来源
	AccountType   int32 `form:"accountType,omitempty"`   //账号类型（=1 企业 =2 个人）
	AccountOrigin int32 `form:"accountOrigin,omitempty"` //账号来源（=1 注册 =2 管理员创建）
}

func DeveloperEntitysListReq_entitysToPb(src *DeveloperEntitysListReq) *protosService.DeveloperListRequest {
	obj := protosService.DeveloperListRequest{
		Page:      int64(src.PageNum),
		PageSize:  int64(src.PageSize),
		SearchKey: src.CompanyName,
	}

	CompanyId, err := strconv.Atoi(src.CompanyId)
	if err != nil {
		CompanyId = 0
	}

	obj.Query = &protosService.DeveloperListSearchInfo{
		Account: src.Account, CompanyId: int64(CompanyId), Status: src.Status,
		AccountOrigin: src.AccountOrigin, AccountType: src.AccountType,
	}
	return &obj
}

type DeveloperEntitys struct {
	Id            string     `json:"id,omitempty"`
	Account       string     `json:"account,omitempty"`
	AccountType   int32      `json:"accountType"` //账号类型（=1 企业 =2 个人）
	Password      string     `json:"password,omitempty"`
	CompanyId     string     `json:"companyId,omitempty"`
	CompanyName   string     `json:"companyName,omitempty"`
	Status        int        `json:"status,omitempty"`
	Phone         string     `json:"phone,omitempty"`
	Email         string     `json:"email,omitempty"`
	Address       string     `json:"address,omitempty"`
	Quantity      int32      `json:"quantity,omitempty"`
	Contract      []Contract `json:"contract,omitempty"`
	RoleName      string     `json:"roleName"`
	TenantId      string     `json:"tenantId"`
	AccountOrigin int32      `json:"accountOrigin"` //账号类型（=1 企业 =2 个人）
}

func (o DeveloperEntitys) ValidId() (int64, error) {
	if o.Id == "" {
		return 0, errors.New("缺id")
	}
	var id int
	var err error
	if id, err = strconv.Atoi(o.Id); err != nil {
		return int64(id), err
	}
	return int64(id), nil
}

type Contract struct {
	Id           string `json:"id"`
	Quantity     int32  `json:"quantity"`
	ContractDate int64  `json:"contractDate"`
}

type CompanyInfo struct {
	CompanyId   string `json:"companyId"`
	CompanyName string `json:"companyName"`
}

type StatusReq struct {
	Id     string `json:"id"`
	Status int32  `json:"status"`
}

func DeveloperEntitys_pbToEntitys(src *protosService.DeveloperEntitys) *DeveloperEntitys {
	obj := DeveloperEntitys{
		Id:            strconv.Itoa(int(src.Id)),
		Account:       src.Account,
		Password:      src.Password,
		CompanyId:     strconv.Itoa(int(src.CompanyId)),
		CompanyName:   src.CompanyName,
		Status:        int(src.Status),
		Phone:         src.Phone,
		Email:         src.Email,
		Quantity:      src.Quantity,
		Address:       src.Address,
		RoleName:      src.RoleName,
		TenantId:      src.TenantId,
		AccountType:   src.AccountType,
		AccountOrigin: src.AccountOrigin,
	}
	for _, v := range src.Contract {
		obj.Contract = append(obj.Contract, Contract{Id: strconv.Itoa(int(v.Id)), Quantity: v.Quantity, ContractDate: v.ContractDate})
	}
	return &obj
}
