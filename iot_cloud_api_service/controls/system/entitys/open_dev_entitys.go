package entitys

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
)

type OpenDevListReq struct {
	CompanyName string `json:"companyName" form:"companyName"` //公司名称
	Status      int32  `json:"status" form:"status"`           //认证状态
	PageSize    int32  `json:"pageSize" form:"pageSize"`
	PageNum     int32  `json:"pageNum" form:"pageNum"`
}

// type OpenDevListRes struct{
// 	Code    int32             `json:"code"`
// 	Message string            `json:"message"`
// 	Total   int64             `json:"total"`
// 	Data    []*OpenDevListEntity `json:"data"`
// }

type OpenDevListEntityRes struct {
	Id            string `json:"id"`
	CompanyName   string `json:"companyName"`     //公司名称
	UserName      string `json:"userName"`        //用户账号 TODO 需要在company中加入冗余
	Status        int32  `json:"status"`          //认证状态
	RequestAuthAt int64  `json:"requestDateTime"` //申请时间
}

// 审核列表
type OpenDevAuthEntity struct {
	Result    string `json:"result"`
	Opter     string `json:"opter"`
	OpterTime int64  `json:"opterTime"`
	Why       string `json:"why"`
}

// 认证详细响应
type OpenDevDetailRes struct {
	ID             string               `json:"id"`
	Nature         string               `json:"nature"`
	Status         int32                `json:"status"`
	PostName       string               `json:"postName"`
	CompanyName    string               `json:"companyName"`
	LicenseNo      string               `json:"licenseNo"`
	License        string               `json:"license"`
	LegalPerson    string               `json:"legalPerson"`
	ApplyPerson    string               `json:"applyPerson"`
	Idcard         string               `json:"idcard"`
	IdcardFrontImg string               `json:"idcardFrontImg"`
	IdcardAfterImg string               `json:"idcardAfterImg"`
	Phone          string               `json:"phone"`
	Address        string               `json:"address"`
	AuthList       []*OpenDevAuthEntity `json:"authList"`
}

type OpenDevCompanyAuthReq struct {
	Id     string `json:"id"`
	Why    string `json:"why"`
	Status int32  `json:"status"`
}

func OpenCompanyToListReq(src *protosService.OpenCompany) *OpenDevListEntityRes {
	if src == nil {
		return nil
	}
	var authAt int64
	authAt = 0
	if src.RequestAuthAt.AsTime().Unix() > 1 {
		authAt = src.RequestAuthAt.AsTime().Unix()
	}
	entitysObj := OpenDevListEntityRes{
		Id:            iotutil.ToString(src.Id),
		CompanyName:   src.Name,
		Status:        src.Status,
		RequestAuthAt: int64(authAt),
		UserName:      src.UserName, //TODO   需要在company中加入冗余
	}
	return &entitysObj

}
