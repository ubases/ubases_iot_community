package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gen/field"

	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_open_system/model"
	"cloud_platform/iot_model/db_open_system/orm"
	proto "cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DeveloperSvc struct {
	Ctx context.Context
}
type ScanObject struct {
	Id       int64
	TenantId string
}

type ScanDeveloperObject struct {
	Id            int64
	UserName      string
	UserNickname  string
	Mobile        string
	UserEmail     string
	UserStatus    int
	Address       string
	CompanyId     int64
	CompanyName   string
	Quantity      int32
	TenantId      string
	AccountType   int32 //账号类型（=1 企业 =2 个人）
	AccountOrigin int32 //账号类型（=1 企业 =2 个人）
}

func (o ScanDeveloperObject) toPb() *proto.DeveloperEntitys {
	obj := proto.DeveloperEntitys{
		Id:            o.Id,
		Account:       o.UserName,
		Password:      "",
		CompanyId:     o.CompanyId,
		CompanyName:   o.CompanyName,
		Status:        int32(o.UserStatus),
		Phone:         o.Mobile,
		Email:         o.UserEmail,
		Address:       o.Address,
		Quantity:      o.Quantity,
		TenantId:      o.TenantId,
		AccountType:   o.AccountType,
		AccountOrigin: o.AccountOrigin,
	}
	return &obj
}

type ScanContractObject struct {
	Id           int64
	Quantity     int32
	ContractDate time.Time
}

func (o ScanContractObject) toPb() *proto.Contract {
	obj := proto.Contract{
		Id:           o.Id,
		Quantity:     o.Quantity,
		ContractDate: o.ContractDate.Unix(),
	}
	return &obj
}

func genAuthCode(TenantId string, Quantity int32) (string, error) {
	text := strings.Join([]string{TenantId, strconv.FormatInt(int64(Quantity), 10)}, "/")
	str, err := iotutil.AES_CBC_EncryptHex([]byte(text), []byte(TenantId))
	if err != nil {
		return "", err
	}
	return str, nil
}

func (d DeveloperSvc) AddAuthQuantity(userId int64, contract []*proto.Contract) error {
	//给用户增加模组数量
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	var scanObject ScanObject
	err := do.Select(t.Id, t.TenantId).Where(t.UserId.Eq(userId)).Scan(&scanObject)
	if err != nil {
		return err
	}
	if len(contract) == 0 {
		contract = append(contract, &proto.Contract{Quantity: 50, ContractDate: time.Now().Unix()})
	}
	for _, v := range contract {
		//过虑已添加的模组数量，web调用更新接口会将已添加过的和待添加的一起上传，授权数量已添加过的id>0，待添加的id=0
		if v.Id > 0 {
			continue
		}
		AuthCode, err := genAuthCode(scanObject.TenantId, v.Quantity)
		if err != nil {
			continue
		}
		var oaqs OpenAuthQuantitySvc
		_, err = oaqs.CreateOpenAuthQuantity(&proto.OpenAuthQuantity{Id: iotutil.GetNextSeqInt64(),
			UserId: userId, CompanyId: scanObject.Id,
			TenantId: scanObject.TenantId, AuthCode: AuthCode, AuthQuantity: v.Quantity,
			AuthDate: timestamppb.New(time.Unix(v.ContractDate, 0)), Status: 1})
		if err != nil {
			iotlogger.LogHelper.Errorf("添加授权数量失败:%s", err.Error())
			continue
		}
	}
	return nil
}

func (d DeveloperSvc) Add(req *proto.DeveloperEntitys) (*proto.Response, error) {
	//检查该账户是否已存在
	var ous OpenUserSvc
	userId, err := ous.AddUser(req)
	if err != nil {
		return nil, err
	}
	err = d.AddAuthQuantity(userId, req.Contract)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (d DeveloperSvc) Edit(req *proto.DeveloperEntitys) (*proto.Response, error) {
	if req.Id == 0 {
		return nil, errors.New("缺userId")
	}
	var ous OpenUserSvc
	req1 := proto.OpenUserUpdateFieldsRequest{
		Fields: []string{"mobile", "user_status", "user_email", "address", "company_name", "account_type"},
		Data: &proto.OpenUser{Id: req.Id, Mobile: req.Phone, UserStatus: req.Status,
			UserEmail: req.Email, Address: req.Address, CompanyName: req.CompanyName, AccountType: req.AccountType},
	}
	_, err := ous.UpdateFieldsOpenUser(&req1)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("更新用户[%d]数据失败:%s", req.Id, err.Error()))
	}
	var ocs OpenCompanySvc
	reqCompany := proto.OpenCompanyUpdateFieldsRequest{
		Fields: []string{"name", "account_type"},
		Data:   &proto.OpenCompany{Id: req.CompanyId, Name: req.CompanyName, AccountType: req.AccountType},
	}
	_, err = ocs.UpdateFieldsOpenCompany(&reqCompany)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("更新公司名称[%d]数据失败:%s", req.Id, err.Error()))
	}
	err = d.AddAuthQuantity(req.Id, req.Contract)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (d DeveloperSvc) Detail(req *proto.DeveloperFilterReq) (*proto.DeveloperEntitys, error) {
	if req.Id == 0 { //整数
		return nil, errors.New("开发者ID不能为空")
	}
	ctx := context.Background()
	query := orm.Use(iotmodel.GetDB())
	tOpenUser := query.TOpenUser
	tOpenCompany := query.TOpenCompany
	tOpenAuthQuantity := query.TOpenAuthQuantity
	subQuery := tOpenAuthQuantity.WithContext(ctx).Select(field.NewField(tOpenAuthQuantity.TableName(),
		"user_id"), tOpenAuthQuantity.AuthQuantity.Sum().As("Quantity")).
		Where(tOpenAuthQuantity.Status.Gt(0)).Group(field.NewField(tOpenAuthQuantity.TableName(),
		"user_id")).As("tmp")
	do := tOpenUser.WithContext(ctx).Select(tOpenUser.Id, tOpenUser.UserName, tOpenUser.UserNickname,
		tOpenUser.Mobile, tOpenUser.UserEmail, tOpenUser.UserStatus, tOpenUser.Address, tOpenUser.AccountOrigin, tOpenCompany.AccountType,
		tOpenCompany.Id.As("CompanyId"), tOpenCompany.Name.As("CompanyName"),
		field.NewField("tmp", "Quantity")).LeftJoin(tOpenCompany,
		tOpenUser.Id.EqCol(tOpenCompany.UserId))
	do = do.LeftJoin(subQuery, field.NewField("tmp", "user_id").EqCol(tOpenUser.Id))
	var scanObject ScanDeveloperObject
	err := do.Where(tOpenUser.Id.Eq(req.Id)).Scan(&scanObject)
	if err != nil {
		return nil, err
	}
	pb := scanObject.toPb()
	var list []*ScanContractObject
	t := query.TOpenAuthQuantity
	do2 := t.WithContext(context.Background()).Select(t.Id, t.AuthQuantity.As("Quantity"), t.AuthDate.As("ContractDate"))
	do2 = do2.Where(t.UserId.Eq(req.Id), t.Status.Gt(0)).Order(t.Id.Desc())
	err = do2.Scan(&list)
	if err == nil && len(list) > 0 {
		pb.Contract = make([]*proto.Contract, len(list))
		for i, v := range list {
			pb.Contract[i] = v.toPb()
		}
	}

	//根据公司id 获取用户角色
	roleName, _, errRole := d.GetCompnayIdFindRoleName(pb.CompanyId)
	if errRole != nil {
		return nil, errRole
	}
	pb.RoleName = roleName
	return pb, nil
}

// 根据公司id 获取用户角色
func (d DeveloperSvc) GetCompnayIdFindRoleName(companyId int64) (string, string, error) {
	//根据公司id 获取用户角色
	var companySvc = OpenCompanySvc{}
	res, errCompany := companySvc.FindByIdOpenCompany(&proto.OpenCompanyFilter{Id: companyId})
	if errCompany != nil {
		return "", "", errCompany
	}
	var roleSvc = OpenRoleSvc{}
	roleids := roleSvc.GetUserCompanyRoleId(res.UserId, res.TenantId)
	if len(roleids) > 0 {
		role, errRole := roleSvc.FindOpenRole(&proto.OpenRoleFilter{Id: iotutil.ToInt64(roleids[0])})
		if errRole != nil {
			return "", res.TenantId, errRole
		}
		return role.Name, res.TenantId, nil
	}
	return "", res.TenantId, nil
}

func (d DeveloperSvc) Delete(req *proto.DeveloperFilterReq) (*proto.Response, error) {
	if req.Id == 0 { //整数
		return nil, errors.New("开发者ID不能为空")
	}
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	_, err := do.Where(t.Id.Eq(req.Id)).Delete()
	if err != nil {
		logger.Errorf("DeveloperSvc.Delete error : %s", err.Error())
		return nil, err
	}
	return nil, nil
}

func (d DeveloperSvc) SetStatus(req *proto.DeveloperStatusReq) (*proto.Response, error) {
	if req.Id == 0 { //整数
		return nil, errors.New("开发者ID不能为空")
	}
	if req.Status < 0 || req.Status > 2 {
		return nil, errors.New(fmt.Sprintf("Status取值范围:0,1,2,当前是:%d", req.Status))
	}
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	_, err := do.Where(t.Id.Eq(req.Id)).Update(t.UserStatus, req.Status)
	if err != nil {
		logger.Errorf("DeveloperSvc.SetStatus error : %s", err.Error())
		return nil, err
	}
	return nil, nil
}

func (d DeveloperSvc) List(req *proto.DeveloperListRequest) ([]*proto.DeveloperEntitys, int64, error) {
	ctx := context.Background()
	query := orm.Use(iotmodel.GetDB())
	tOpenUser := query.TOpenUser
	tOpenCompany := query.TOpenCompany
	tOpenAuthQuantity := query.TOpenAuthQuantity
	subQuery := tOpenAuthQuantity.WithContext(ctx).Select(field.NewField(tOpenAuthQuantity.TableName(),
		"user_id"), tOpenAuthQuantity.AuthQuantity.Sum().As("Quantity")).
		Where(tOpenAuthQuantity.Status.Gt(0)).Group(field.NewField(tOpenAuthQuantity.TableName(),
		"user_id")).As("tmp")
	do := tOpenUser.WithContext(ctx).Select(tOpenUser.Id, tOpenUser.UserName, tOpenUser.UserNickname,
		tOpenUser.Mobile, tOpenUser.UserEmail, tOpenUser.UserStatus, tOpenUser.Address, tOpenUser.AccountOrigin, tOpenUser.AccountType,
		tOpenCompany.TenantId.As("tenantId"), tOpenCompany.Id.As("CompanyId"), tOpenCompany.Name.As("CompanyName"),
		field.NewField("tmp", "Quantity")).LeftJoin(tOpenCompany,
		tOpenUser.Id.EqCol(tOpenCompany.UserId))
	do = do.LeftJoin(subQuery, field.NewField("tmp", "user_id").EqCol(tOpenUser.Id))
	if req != nil && req.Query != nil {
		if req.Query.Account != "" {
			do = do.Where(tOpenUser.UserName.Like("%" + req.Query.Account + "%"))
		}
		if req.Query.CompanyId > 0 {
			do = do.Where(tOpenCompany.Id.Eq(req.Query.CompanyId))
		}
		if req.Query.Status > 0 {
			do = do.Where(tOpenUser.UserStatus.Eq(req.Query.Status))
		}
		if req.Query.AccountType > 0 {
			do = do.Where(tOpenUser.AccountType.Eq(req.Query.AccountType))
		}
		if req.Query.AccountOrigin > 0 {
			do = do.Where(tOpenUser.AccountOrigin.Eq(req.Query.AccountOrigin))
		}
	}
	if req.SearchKey != "" {
		do = do.Where(tOpenCompany.Name.Like("%" + req.SearchKey + "%"))
	}
	orderCol, ok := tOpenUser.GetFieldByName(req.OrderKey)
	if !ok {
		do = do.Order(tOpenUser.CreatedAt.Desc())
	} else {
		do = do.Order(orderCol.Desc())
	}
	var list []*ScanDeveloperObject
	var total int64
	var err error
	if req.PageSize > 0 {
		limit := req.PageSize
		if req.Page == 0 {
			req.Page = 1
		}
		offset := req.PageSize * (req.Page - 1)
		total, err = do.ScanByPage(&list, int(offset), int(limit))
	} else {
		err = do.Scan(&list)
		total = int64(len(list))
	}
	if err != nil {
		logger.Errorf("DeveloperSvc.List error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.DeveloperEntitys, len(list))
	for i, v := range list {
		//TODO  这里很奇怪。需求不适合. 后期建议去掉. 暂时满足测试需求
		roleName, tenantId, _ := d.GetCompnayIdFindRoleName(v.CompanyId)
		result[i] = v.toPb()
		result[i].RoleName = roleName
		result[i].TenantId = tenantId
	}
	return result, total, nil
}

func (d DeveloperSvc) BasicList(req *proto.DeveloperListRequest) ([]*proto.DeveloperEntitys, int64, error) {
	ctx := context.Background()
	query := orm.Use(iotmodel.GetDB())
	tOpenUser := query.TOpenUser
	tOpenCompany := query.TOpenCompany

	do := tOpenUser.WithContext(ctx).Select(tOpenUser.Id, tOpenUser.UserName, tOpenUser.UserNickname,
		tOpenUser.Mobile, tOpenUser.UserEmail, tOpenUser.UserStatus, tOpenUser.Address,
		tOpenCompany.TenantId.As("TenantId"), tOpenCompany.Id.As("CompanyId"), tOpenCompany.Name.As("CompanyName")).
		LeftJoin(tOpenCompany, tOpenUser.Id.EqCol(tOpenCompany.UserId))

	if req != nil && req.Query != nil {
		if req.Query.Account != "" {
			do = do.Where(tOpenUser.UserName.Like("%" + req.Query.Account + "%"))
		}
		if req.Query.CompanyId > 0 {
			do = do.Where(tOpenCompany.Id.Eq(req.Query.CompanyId))
		}
		if req.Query.Status > 0 {
			do = do.Where(tOpenUser.UserStatus.Eq(req.Query.Status))
		}
	}
	if req.SearchKey != "" {
		do = do.Where(tOpenCompany.Name.Like("%" + req.SearchKey + "%"))
	}
	orderCol, ok := tOpenUser.GetFieldByName(req.OrderKey)
	if !ok {
		do = do.Order(tOpenUser.CreatedAt.Desc())
	} else {
		do = do.Order(orderCol.Desc())
	}
	var list []*ScanDeveloperObject
	var total int64
	var err error
	if req.PageSize > 0 {
		limit := req.PageSize
		if req.Page == 0 {
			req.Page = 1
		}
		offset := req.PageSize * (req.Page - 1)
		total, err = do.ScanByPage(&list, int(offset), int(limit))
	} else {
		err = do.Scan(&list)
		total = int64(len(list))
	}
	if err != nil {
		logger.Errorf("DeveloperSvc.List error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.DeveloperEntitys, len(list))
	for i, v := range list {
		result[i] = v.toPb()
	}
	return result, total, nil
}

func (d DeveloperSvc) ResetPassword(req *proto.DeveloperResetPasswordReq) (*proto.Response, error) {
	if req.Id == 0 { //整数
		return nil, errors.New("开发者ID不能为空")
	}
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	UserSalt := iotutil.GetSecret(8)
	newPassword := iotutil.Md5(req.DefaultPassword)
	password := iotutil.Md5(newPassword + UserSalt)
	user := &model.TOpenUser{Id: req.Id, UserPassword: password, UserSalt: UserSalt}
	_, err := do.Select(t.UserPassword, t.UserSalt).Where(t.Id.Eq(req.Id)).Updates(user)
	if err != nil {
		logger.Errorf("DeveloperSvc.ResetPassword error : %s", err.Error())
		return nil, err
	}
	return nil, nil

}
