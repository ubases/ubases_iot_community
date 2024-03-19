// Code generated by sgen.exe,2022-04-27 10:55:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"
	"time"

	"go-micro.dev/v4/metadata"

	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_open_system/model"
	"cloud_platform/iot_model/db_open_system/orm"
	"cloud_platform/iot_open_system_service/convert"
	proto "cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gen/field"
)

type OpenCompanySvc struct {
	Ctx context.Context
}

// 设置公共属性[创建时间,创建人,修改时间,修改人]新增 optertype =1 修改 =2
func (s *OpenCompanySvc) SetCommonFiled(req *proto.OpenCompany, userId int64, opterType int) {
	if opterType == 1 {
		req.CreatedBy = userId
		req.CreatedAt = timestamppb.New(time.Now())
	}
	req.UpdatedBy = userId
	req.UpdatedAt = timestamppb.New(time.Now())
}

// 创建OpenCompany
func (s *OpenCompanySvc) CreateOpenCompany(req *proto.OpenCompany) (*proto.OpenCompany, error) {
	// fixme 请在这里校验参数

	s.SetCommonFiled(req, req.CreatedBy, 1)

	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	dbObj := convert.OpenCompany_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OpenCompany
func (s *OpenCompanySvc) DeleteOpenCompany(req *proto.OpenCompany) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.UserId != 0 { //整数
		do = do.Where(t.UserId.Eq(req.UserId))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Eq(req.Name))
	}
	if req.Nature != "" { //字符串
		do = do.Where(t.Nature.Eq(req.Nature))
	}
	if req.LicenseNo != "" { //字符串
		do = do.Where(t.LicenseNo.Eq(req.LicenseNo))
	}
	if req.License != "" { //字符串
		do = do.Where(t.License.Eq(req.License))
	}
	if req.LegalPerson != "" { //字符串
		do = do.Where(t.LegalPerson.Eq(req.LegalPerson))
	}
	if req.ApplyPerson != "" { //字符串
		do = do.Where(t.ApplyPerson.Eq(req.ApplyPerson))
	}
	if req.Idcard != "" { //字符串
		do = do.Where(t.Idcard.Eq(req.Idcard))
	}
	if req.IdcardFrontImg != "" { //字符串
		do = do.Where(t.IdcardFrontImg.Eq(req.IdcardFrontImg))
	}
	if req.IdcardAfterImg != "" { //字符串
		do = do.Where(t.IdcardAfterImg.Eq(req.IdcardAfterImg))
	}
	if req.Address != "" { //字符串
		do = do.Where(t.Address.Eq(req.Address))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.AccountType != 0 { //整数
		do = do.Where(t.AccountType.Eq(req.AccountType))
	}
	if req.CaseRemak != "" { //字符串
		do = do.Where(t.CaseRemak.Eq(req.CaseRemak))
	}
	if req.Email != "" { //字符串
		do = do.Where(t.Email.Eq(req.Email))
	}
	if req.IsRealName != 0 { //整数
		do = do.Where(t.IsRealName.Eq(req.IsRealName))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OpenCompany
func (s *OpenCompanySvc) DeleteByIdOpenCompany(req *proto.OpenCompany) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OpenCompany
func (s *OpenCompanySvc) DeleteByIdsOpenCompany(req *proto.OpenCompanyBatchDeleteRequest) (*proto.OpenCompanyBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOpenCompany
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOpenCompany error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OpenCompany
func (s *OpenCompanySvc) UpdateOpenCompany(req *proto.OpenCompany) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.UserId != 0 { //整数
		updateField = append(updateField, t.UserId)
	}
	if req.Name != "" { //字符串
		updateField = append(updateField, t.Name)
	}
	if req.Nature != "" { //字符串
		updateField = append(updateField, t.Nature)
	}
	if req.LicenseNo != "" { //字符串
		updateField = append(updateField, t.LicenseNo)
	}
	if req.License != "" { //字符串
		updateField = append(updateField, t.License)
	}
	if req.LegalPerson != "" { //字符串
		updateField = append(updateField, t.LegalPerson)
	}
	if req.ApplyPerson != "" { //字符串
		updateField = append(updateField, t.ApplyPerson)
	}
	if req.Idcard != "" { //字符串
		updateField = append(updateField, t.Idcard)
	}
	if req.IdcardFrontImg != "" { //字符串
		updateField = append(updateField, t.IdcardFrontImg)
	}
	if req.IdcardAfterImg != "" { //字符串
		updateField = append(updateField, t.IdcardAfterImg)
	}
	if req.Address != "" { //字符串
		updateField = append(updateField, t.Address)
	}
	if req.Status != 0 { //整数
		updateField = append(updateField, t.Status)
	}
	if req.AccountType != 0 { //整数
		updateField = append(updateField, t.AccountType)
	}
	if req.CaseRemak != "" { //字符串
		updateField = append(updateField, t.CaseRemak)
	}
	if req.Email != "" { //字符串
		updateField = append(updateField, t.Email)
	}
	if req.IsRealName != 0 { //整数
		updateField = append(updateField, t.IsRealName)
	}
	if req.CreatedBy != 0 { //整数
		updateField = append(updateField, t.CreatedBy)
	}
	if req.UpdatedBy != 0 { //整数
		updateField = append(updateField, t.UpdatedBy)
	}
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateOpenCompany error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OpenCompany_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OpenCompany
func (s *OpenCompanySvc) UpdateAllOpenCompany(req *proto.OpenCompany) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.UserId)
	updateField = append(updateField, t.Name)
	updateField = append(updateField, t.Nature)
	updateField = append(updateField, t.LicenseNo)
	updateField = append(updateField, t.License)
	updateField = append(updateField, t.LegalPerson)
	updateField = append(updateField, t.ApplyPerson)
	updateField = append(updateField, t.Idcard)
	updateField = append(updateField, t.IdcardFrontImg)
	updateField = append(updateField, t.IdcardAfterImg)
	updateField = append(updateField, t.Address)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.AccountType)
	updateField = append(updateField, t.CaseRemak)
	updateField = append(updateField, t.Email)
	updateField = append(updateField, t.IsRealName)
	updateField = append(updateField, t.CreatedBy)
	updateField = append(updateField, t.UpdatedBy)
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false
	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateAllOpenCompany error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenCompany_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OpenCompanySvc) UpdateFieldsOpenCompany(req *proto.OpenCompanyUpdateFieldsRequest) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOpenCompany error : missing updateField")
		logger.Error(err)
		return nil, err
	}
	do = do.Select(updateField...)

	//主键条件
	HasPrimaryKey := false
	if req.Data.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Data.Id))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsOpenCompany error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenCompany_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 企业认证
func (s *OpenCompanySvc) CompanyAuth(req *proto.OpenCompanyUpdateFieldsRequest) (*proto.OpenCompany, error) {
	//1.检查状态
	company, errCompany := s.FindByIdOpenCompany(&proto.OpenCompanyFilter{
		Id: req.Data.Id,
	})
	if errCompany != nil {
		return nil, errCompany
	}
	//1=未提交,  2=认证中,   4=不通过,   3=已认证,   5=已撤销
	//1=未提交,  5=已撤销, 4=不通过 才可以修改
	if company.Status != 1 && company.Status != 5 && company.Status != 4 {
		return nil, errors.New("当前状态, 不可提交认证.")
	}
	req.Data.Status = 2
	req.Fields = append(req.Fields, "status")
	req.Data.RequestAuthAt = timestamppb.New(time.Now())
	req.Fields = append(req.Fields, "request_auth_at")

	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOpenCompany error : missing updateField")
		logger.Error(err)
		return nil, err
	}
	do = do.Select(updateField...)

	//主键条件
	HasPrimaryKey := false
	if req.Data.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Data.Id))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsOpenCompany error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenCompany_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOpenCompany error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OpenCompany
func (s *OpenCompanySvc) FindOpenCompany(req *proto.OpenCompanyFilter) (*proto.OpenCompany, error) {
	tenantId, _ := metadata.Get(s.Ctx, "tenantid")
	if tenantId != "" {
		req.TenantId = tenantId
	}

	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.UserId != 0 { //整数
		do = do.Where(t.UserId.Eq(req.UserId))
	}
	if req.TenantId != "" {
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Eq(req.Name))
	}
	if req.Nature != "" { //字符串
		do = do.Where(t.Nature.Eq(req.Nature))
	}
	if req.LicenseNo != "" { //字符串
		do = do.Where(t.LicenseNo.Eq(req.LicenseNo))
	}
	if req.License != "" { //字符串
		do = do.Where(t.License.Eq(req.License))
	}
	if req.LegalPerson != "" { //字符串
		do = do.Where(t.LegalPerson.Eq(req.LegalPerson))
	}
	if req.ApplyPerson != "" { //字符串
		do = do.Where(t.ApplyPerson.Eq(req.ApplyPerson))
	}
	if req.Idcard != "" { //字符串
		do = do.Where(t.Idcard.Eq(req.Idcard))
	}
	if req.IdcardFrontImg != "" { //字符串
		do = do.Where(t.IdcardFrontImg.Eq(req.IdcardFrontImg))
	}
	if req.IdcardAfterImg != "" { //字符串
		do = do.Where(t.IdcardAfterImg.Eq(req.IdcardAfterImg))
	}
	if req.Address != "" { //字符串
		do = do.Where(t.Address.Eq(req.Address))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.AccountType != 0 { //整数
		do = do.Where(t.AccountType.Eq(req.AccountType))
	}
	if req.CaseRemak != "" { //字符串
		do = do.Where(t.CaseRemak.Eq(req.CaseRemak))
	}
	if req.Email != "" { //字符串
		do = do.Where(t.Email.Eq(req.Email))
	}
	if req.IsRealName != 0 { //整数
		do = do.Where(t.IsRealName.Eq(req.IsRealName))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpenCompany error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenCompany_db2pb(dbObj)
	return res, err
}

// 根据非空条件查找OpenCompany
func (s *OpenCompanySvc) FindOpenCompanyNoCtx(req *proto.OpenCompanyFilter) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.UserId != 0 { //整数
		do = do.Where(t.UserId.Eq(req.UserId))
	}
	if req.TenantId != "" {
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Eq(req.Name))
	}
	if req.Nature != "" { //字符串
		do = do.Where(t.Nature.Eq(req.Nature))
	}
	if req.LicenseNo != "" { //字符串
		do = do.Where(t.LicenseNo.Eq(req.LicenseNo))
	}
	if req.License != "" { //字符串
		do = do.Where(t.License.Eq(req.License))
	}
	if req.LegalPerson != "" { //字符串
		do = do.Where(t.LegalPerson.Eq(req.LegalPerson))
	}
	if req.ApplyPerson != "" { //字符串
		do = do.Where(t.ApplyPerson.Eq(req.ApplyPerson))
	}
	if req.Idcard != "" { //字符串
		do = do.Where(t.Idcard.Eq(req.Idcard))
	}
	if req.IdcardFrontImg != "" { //字符串
		do = do.Where(t.IdcardFrontImg.Eq(req.IdcardFrontImg))
	}
	if req.IdcardAfterImg != "" { //字符串
		do = do.Where(t.IdcardAfterImg.Eq(req.IdcardAfterImg))
	}
	if req.Address != "" { //字符串
		do = do.Where(t.Address.Eq(req.Address))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.AccountType != 0 { //整数
		do = do.Where(t.AccountType.Eq(req.AccountType))
	}
	if req.CaseRemak != "" { //字符串
		do = do.Where(t.CaseRemak.Eq(req.CaseRemak))
	}
	if req.Email != "" { //字符串
		do = do.Where(t.Email.Eq(req.Email))
	}
	if req.IsRealName != 0 { //整数
		do = do.Where(t.IsRealName.Eq(req.IsRealName))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpenCompany error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenCompany_db2pb(dbObj)
	return res, err
}

// 查找公司记录.状态不等于的查询
func (s *OpenCompanySvc) FindOpenCompanyNot4(req *proto.OpenCompanyFilter) (*proto.OpenCompany, error) {
	tenantId, _ := metadata.Get(s.Ctx, "tenantid")
	if tenantId != "" {
		req.TenantId = tenantId
	}

	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.UserId != 0 { //整数
		do = do.Where(t.UserId.Eq(req.UserId))
	}
	if req.TenantId != "" {
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}

	//状态不等于
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Neq(req.Status))
	}

	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpenCompany error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenCompany_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找OpenCompany
func (s *OpenCompanySvc) FindByIdOpenCompany(req *proto.OpenCompanyFilter) (*proto.OpenCompany, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOpenCompany error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenCompany_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OpenCompany,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OpenCompanySvc) GetListOpenCompany(req *proto.OpenCompanyListRequest) ([]*proto.OpenCompany, int64, error) {
	var err error
	t := orm.Use(iotmodel.GetDB()).TOpenCompany
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.UserId != 0 { //整数
			do = do.Where(t.UserId.Eq(query.UserId))
		}
		if query.TenantId != "" {
			do = do.Where(t.TenantId.Eq(query.TenantId))
		}
		if query.Name != "" { //字符串
			do = do.Where(t.Name.Like("%" + query.Name + "%"))
		}
		if query.Nature != "" { //字符串
			do = do.Where(t.Nature.Like("%" + query.Nature + "%"))
		}
		if query.LicenseNo != "" { //字符串
			do = do.Where(t.LicenseNo.Like("%" + query.LicenseNo + "%"))
		}
		if query.License != "" { //字符串
			do = do.Where(t.License.Like("%" + query.License + "%"))
		}
		if query.LegalPerson != "" { //字符串
			do = do.Where(t.LegalPerson.Like("%" + query.LegalPerson + "%"))
		}
		if query.ApplyPerson != "" { //字符串
			do = do.Where(t.ApplyPerson.Like("%" + query.ApplyPerson + "%"))
		}
		if query.Idcard != "" { //字符串
			do = do.Where(t.Idcard.Like("%" + query.Idcard + "%"))
		}
		if query.IdcardFrontImg != "" { //字符串
			do = do.Where(t.IdcardFrontImg.Like("%" + query.IdcardFrontImg + "%"))
		}
		if query.IdcardAfterImg != "" { //字符串
			do = do.Where(t.IdcardAfterImg.Like("%" + query.IdcardAfterImg + "%"))
		}
		if query.Address != "" { //字符串
			do = do.Where(t.Address.Like("%" + query.Address + "%"))
		}
		if query.Status != 0 { //整数
			do = do.Where(t.Status.Eq(query.Status))
		}
		if query.AccountType != 0 { //整数
			do = do.Where(t.AccountType.Eq(query.AccountType))
		}
		if query.CaseRemak != "" { //字符串
			do = do.Where(t.CaseRemak.Like("%" + query.CaseRemak + "%"))
		}
		if query.Email != "" { //字符串
			do = do.Where(t.Email.Like("%" + query.Email + "%"))
		}
		if query.IsRealName != 0 { //整数
			do = do.Where(t.IsRealName.Eq(query.IsRealName))
		}
		if query.CreatedBy != 0 { //整数
			do = do.Where(t.CreatedBy.Eq(query.CreatedBy))
		}
		if query.UpdatedBy != 0 { //整数
			do = do.Where(t.UpdatedBy.Eq(query.UpdatedBy))
		}
		if query.UserName != "" {
			do = do.Where(t.Email.Like("%" + query.UserName + "%"))
		}
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = t.Id
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}

	var list []*model.TOpenCompany
	var total int64
	if req.PageSize > 0 {
		limit := req.PageSize
		if req.Page == 0 {
			req.Page = 1
		}
		offset := req.PageSize * (req.Page - 1)
		list, total, err = do.FindByPage(int(offset), int(limit))
	} else {
		list, err = do.Find()
		total = int64(len(list))
	}
	if err != nil {
		logger.Errorf("GetListOpenCompany error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OpenCompany, len(list))
	for i, v := range list {
		result[i] = convert.OpenCompany_db2pb(v)
	}
	return result, total, nil
}

func (s *OpenCompanySvc) OpenDevCompanyAuth(req *proto.OpenDevCompanyAuthRequest) error {
	userId, _ := metadata.Get(s.Ctx, "userid")
	if userId == "" {
		return errors.New("用户信息获取失败,请重新登录.")
	}
	AuthName := req.AuthName
	Id := iotutil.ToInt64(req.Id)
	status := req.Status // 1 通过, 2 不通过, 3 撤销
	why := req.Why

	//获取出数据库的状态
	company, errCompany := s.FindByIdOpenCompany(&proto.OpenCompanyFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if errCompany != nil {
		return errCompany
	}
	//(已弃用)旧版状态（=1 未提交 ,=2 认证中,   =3 已认证, =4 禁用)
	//(在用中)新版状态(1=未提交,  2=认证中,   4=不通过,   3=已认证,   5=已撤销)

	if status == 3 && company.Status != 3 {
		return errors.New("不是已认证状态, 不可以进行撤销操作.")
	}

	//如果状态不是认证中.则提示状态不对.
	if (status == 1 || status == 2) && company.Status != 2 {
		return errors.New("不是认证中状态, 不可进行审核操作.")
	}

	//操作数据库

	dataStatus := 1
	resultString := "通过"
	if status == 1 { //通过
		dataStatus = 3
		resultString = "通过"
	} else if status == 2 { //不通过
		dataStatus = 4
		resultString = "不通过"
	} else if status == 3 { //撤销
		dataStatus = 5
		resultString = "撤销"
	}

	_, errUp := s.UpdateFieldsOpenCompany(&proto.OpenCompanyUpdateFieldsRequest{
		Fields: []string{"status", "case_remak", "updated_at", "updated_by"},
		Data: &proto.OpenCompany{
			Id:        Id,
			Status:    int32(dataStatus),
			CaseRemak: why,
			UpdatedAt: timestamppb.New(time.Now()),
			UpdatedBy: iotutil.ToInt64(userId),
		},
	})
	if errUp != nil {
		return errUp
	}

	//插入审核记录
	var authLogs = OpenCompanyAuthLogsSvc{}
	authLogsId := iotutil.GetNextSeqInt64()
	authLogs.CreateOpenCompanyAuthLogs(&proto.OpenCompanyAuthLogs{
		Id:         authLogsId,
		CompanyId:  Id,
		AuthResult: resultString,
		AuthName:   AuthName,
		AuthDate:   timestamppb.New(time.Now()),
		Why:        why,
		CreatedBy:  iotutil.ToInt64(userId),
		CreatedAt:  timestamppb.New(time.Now()),
	})
	return nil
}
