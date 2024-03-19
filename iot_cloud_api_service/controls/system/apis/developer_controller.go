package apis

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls"
	services2 "cloud_platform/iot_cloud_api_service/controls/global"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"
	"strings"

	"cloud_platform/iot_cloud_api_service/rpc"

	"github.com/gin-gonic/gin"
)

var Developercontroller DeveloperController

var developerService services.DeveloperService

type DeveloperController struct {
} //用户操作控制器

// 添加开发者
func (DeveloperController) Add(c *gin.Context) {
	var req entitys.DeveloperEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if strings.Trim(req.Account, " ") == "" {
		iotgin.ResErrCli(c, errors.New("账号为空"))
		return
	}
	//添加账号：增加账号类型选项；
	//企业账号，企业名称必填，个人账号，企业名称选填
	if req.AccountType == iotconst.OPEN_USER_ENTERPRISE_ACCOUNT && req.CompanyName == "" {
		iotgin.ResBadRequest(c, "企业名称不能为空")
		return
	}

	//账号格式验证，账号必须是手机号码或者邮箱。
	if !iotutil.IsEmail(req.Account) && !iotutil.IsPhone(req.Account) {
		iotgin.ResErrCli(c, errors.New("账号必须是手机号或者邮箱"))
		return
	}

	req.AccountOrigin = iotconst.OPEN_USER_ACCOUNT_TYPE_ADD
	reqq := &protosService.DeveloperEntitys{Account: req.Account, Password: req.Password, CompanyName: req.CompanyName,
		Status: int32(req.Status), Phone: req.Phone, Email: req.Email, Address: req.Address, AccountType: req.AccountType, AccountOrigin: req.AccountOrigin}
	for _, v := range req.Contract {
		reqq.Contract = append(reqq.Contract, &protosService.Contract{Quantity: v.Quantity, ContractDate: v.ContractDate})
	}
	resp, err := rpc.ClientDeveloperService.Add(context.Background(), reqq)
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	if resp.Code != 200 {
		iotgin.ResFailCode(c, resp.Message, 500)
		return
	}
	//刷新缓存
	services2.RefreshDevelopCache()
	iotgin.ResSuccess(c, nil)
}

// 修改开发者
func (DeveloperController) Edit(c *gin.Context) {
	var req entitys.DeveloperEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := req.ValidId()
	if err != nil {
		iotgin.ResErrCli(c, err)
	}

	reqq := &protosService.DeveloperEntitys{Id: id, CompanyName: req.CompanyName, CompanyId: iotutil.ToInt64(req.CompanyId),
		Status: int32(req.Status), Phone: req.Phone, Email: req.Email, Address: req.Address, AccountType: req.AccountType}
	for _, v := range req.Contract {
		if v.Id == "" {
			reqq.Contract = append(reqq.Contract, &protosService.Contract{Quantity: v.Quantity, ContractDate: v.ContractDate})
		}
	}
	resp, err := rpc.ClientDeveloperService.Edit(context.Background(), reqq)
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	if resp.Code != 200 {
		iotgin.ResFailCode(c, resp.Message, 500)
		return
	}
	iotgin.ResSuccess(c, nil)
}

func GetQueryId(c *gin.Context) int64 {
	idStr := c.Query("id")
	if idStr == "" {
		return 0
	}
	if id, err := strconv.Atoi(idStr); err == nil {
		return int64(id)
	}
	return 0
}

// 查看详情
func (DeveloperController) Detail(c *gin.Context) {
	id := GetQueryId(c)
	if id == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}
	resp, err := rpc.ClientDeveloperService.Detail(context.Background(), &protosService.DeveloperFilterReq{Id: int64(id)})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	ret := entitys.DeveloperEntitys_pbToEntitys(resp)

	iotgin.ResSuccess(c, ret)
}

// 删除开发者
func (DeveloperController) Delete(c *gin.Context) {
	id := GetQueryId(c)
	if id == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}

	_, err := rpc.ClientDeveloperService.Delete(context.Background(), &protosService.DeveloperFilterReq{Id: int64(id)})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// 开发者启用禁用
func (DeveloperController) SetStatus(c *gin.Context) {
	var req entitys.StatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var id int
	if id, err = strconv.Atoi(req.Id); err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	_, err = rpc.ClientDeveloperService.SetStatus(context.Background(), &protosService.DeveloperStatusReq{Id: int64(id), Status: req.Status})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// 查询开发者列表
func (DeveloperController) List(c *gin.Context) {
	var req entitys.DeveloperEntitysListReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	reqq := entitys.DeveloperEntitysListReq_entitysToPb(&req)

	resp, err := rpc.ClientDeveloperService.List(context.Background(), reqq)
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	if resp.Code != 200 {
		iotgin.ResFailCode(c, resp.Message, 500)
		return
	}
	var resultList []*entitys.DeveloperEntitys
	for _, item := range resp.Data {
		resultList = append(resultList, entitys.DeveloperEntitys_pbToEntitys(item))
	}
	iotgin.ResPageSuccess(c, resultList, resp.Total, int(req.PageNum))
}

// 查询开发者列表
func (DeveloperController) BasicList(c *gin.Context) {
	var req entitys.DeveloperEntitysListReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	reqq := entitys.DeveloperEntitysListReq_entitysToPb(&req)

	resp, err := rpc.ClientDeveloperService.BasicList(context.Background(), reqq)
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	if resp.Code != 200 {
		iotgin.ResFailCode(c, resp.Message, 500)
		return
	}
	var resultList []*entitys.DeveloperEntitys
	for _, item := range resp.Data {
		resultList = append(resultList, entitys.DeveloperEntitys_pbToEntitys(item))
	}
	iotgin.ResPageSuccess(c, resultList, resp.Total, int(req.PageNum))
}

// 查询所有开发者公司列表
func (DeveloperController) ListCompany(c *gin.Context) {
	companyName := c.Query("companyName")
	resp, err := rpc.ClientOpenCompanyService.Lists(context.Background(), &protosService.OpenCompanyListRequest{
		Query: &protosService.OpenCompany{Name: companyName},
	})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ent []entitys.CompanyInfo
	for _, v := range resp.Data {
		ent = append(ent, entitys.CompanyInfo{CompanyId: strconv.Itoa(int(v.Id)), CompanyName: v.Name})
	}
	iotgin.ResSuccess(c, ent)
}

// 重置密码
func (DeveloperController) ResetPassword(c *gin.Context) {
	id := GetQueryId(c)
	if id == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}
	_, err := rpc.ClientDeveloperService.ResetPassword(context.Background(), &protosService.DeveloperResetPasswordReq{Id: id, DefaultPassword: config.Global.Service.DefaultPassword})
	if err != nil {
		iotgin.ResErrSrv(c)
		return
	}
	//清除该用户的token
	controls.ClearTokenByUserId(id)

	iotgin.ResSuccess(c, nil)
}
