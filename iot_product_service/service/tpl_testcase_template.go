package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	"cloud_platform/iot_product_service/convert"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type TplTestcaseTemplateSvc struct {
	Ctx context.Context
}

// 创建TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) CreateTplTestcaseTemplate(req *proto.TplTestcaseTemplate) (*proto.TplTestcaseTemplate, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	dbObj := convert.TplTestcaseTemplate_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) DeleteTplTestcaseTemplate(req *proto.TplTestcaseTemplate) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ProductTypeId != 0 { //整数
		do = do.Where(t.ProductTypeId.Eq(req.ProductTypeId))
	}
	if req.ProductId != 0 { //整数
		do = do.Where(t.ProductId.Eq(req.ProductId))
	}
	if req.Lang != "" { //字符串
		do = do.Where(t.Lang.Eq(req.Lang))
	}
	if req.TplCode != "" { //字符串
		do = do.Where(t.TplCode.Eq(req.TplCode))
	}
	if req.TplName != "" { //字符串
		do = do.Where(t.TplName.Eq(req.TplName))
	}
	if req.TplFile != "" { //字符串
		do = do.Where(t.TplFile.Eq(req.TplFile))
	}
	if req.TplDesc != "" { //字符串
		do = do.Where(t.TplDesc.Eq(req.TplDesc))
	}
	if req.Status != -1 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) DeleteByIdTplTestcaseTemplate(req *proto.TplTestcaseTemplate) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) DeleteByIdsTplTestcaseTemplate(req *proto.TplTestcaseTemplateBatchDeleteRequest) (*proto.TplTestcaseTemplateBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsTplTestcaseTemplate error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) UpdateTplTestcaseTemplate(req *proto.TplTestcaseTemplate) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ProductTypeId != 0 { //整数
		updateField = append(updateField, t.ProductTypeId)
	}
	if req.ProductId != 0 { //整数
		updateField = append(updateField, t.ProductId)
	}
	if req.Lang != "" { //字符串
		updateField = append(updateField, t.Lang)
	}
	if req.TplCode != "" { //字符串
		updateField = append(updateField, t.TplCode)
	}
	if req.TplName != "" { //字符串
		updateField = append(updateField, t.TplName)
	}
	if req.TplFile != "" { //字符串
		updateField = append(updateField, t.TplFile)
	}
	if req.TplFileName != "" { //字符串
		updateField = append(updateField, t.TplFileName)
	}
	if req.TplFileSize != 0 { //字符串
		updateField = append(updateField, t.TplFileSize)
	}
	if req.TplDesc != "" { //字符串
		updateField = append(updateField, t.TplDesc)
	}
	if req.Status != 0 { //整数
		updateField = append(updateField, t.Status)
	}
	if req.Version != "" {
		updateField = append(updateField, t.Version)
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
		logger.Error("UpdateTplTestcaseTemplate error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.TplTestcaseTemplate_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) UpdateAllTplTestcaseTemplate(req *proto.TplTestcaseTemplate) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ProductTypeId)
	updateField = append(updateField, t.ProductId)
	updateField = append(updateField, t.Lang)
	updateField = append(updateField, t.TplCode)
	updateField = append(updateField, t.TplName)
	updateField = append(updateField, t.TplFile)
	updateField = append(updateField, t.TplDesc)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.CreatedBy)
	updateField = append(updateField, t.UpdatedBy)
	updateField = append(updateField, t.TplFileName)
	updateField = append(updateField, t.TplFileSize)

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
		logger.Error("UpdateAllTplTestcaseTemplate error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.TplTestcaseTemplate_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *TplTestcaseTemplateSvc) UpdateFieldsTplTestcaseTemplate(req *proto.TplTestcaseTemplateUpdateFieldsRequest) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsTplTestcaseTemplate error : missing updateField")
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
		logger.Error("UpdateFieldsTplTestcaseTemplate error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.TplTestcaseTemplate_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) FindTplTestcaseTemplate(req *proto.TplTestcaseTemplateFilter) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ProductTypeId != 0 { //整数
		do = do.Where(t.ProductTypeId.Eq(req.ProductTypeId))
	}
	if req.ProductId != 0 { //整数
		do = do.Where(t.ProductId.Eq(req.ProductId))
	}
	if req.Lang != "" { //字符串
		do = do.Where(t.Lang.Eq(req.Lang))
	}
	if req.TplCode != "" { //字符串
		do = do.Where(t.TplCode.Eq(req.TplCode))
	}
	if req.TplName != "" { //字符串
		do = do.Where(t.TplName.Eq(req.TplName))
	}
	if req.TplFile != "" { //字符串
		do = do.Where(t.TplFile.Eq(req.TplFile))
	}
	if req.TplDesc != "" { //字符串
		do = do.Where(t.TplDesc.Eq(req.TplDesc))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	res := convert.TplTestcaseTemplate_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找TplTestcaseTemplate
func (s *TplTestcaseTemplateSvc) FindByIdTplTestcaseTemplate(req *proto.TplTestcaseTemplateFilter) (*proto.TplTestcaseTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TTplTestcaseTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdTplTestcaseTemplate error : %s", err.Error())
		return nil, err
	}
	res := convert.TplTestcaseTemplate_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找TplTestcaseTemplate,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *TplTestcaseTemplateSvc) GetListTplTestcaseTemplate(req *proto.TplTestcaseTemplateListRequest) ([]*proto.TplTestcaseTemplate, int64, error) {
	var err error
	ormQuery := orm.Use(iotmodel.GetDB())
	t := ormQuery.TTplTestcaseTemplate
	productType := ormQuery.TPmProductType
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {
		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.ProductTypeId != 0 { //整数
			do = do.Where(t.ProductTypeId.Eq(query.ProductTypeId))
		}
		if query.ProductId != 0 { //整数
			do = do.Where(t.ProductId.Eq(query.ProductId))
		}
		if query.Lang != "" { //字符串
			do = do.Where(t.Lang.Eq(query.Lang))
		}
		if query.TplCode != "" { //字符串
			do = do.Where(t.TplCode.Eq(query.TplCode))
		}
		if query.TplName != "" { //字符串
			do = do.Where(t.TplName.Like("%" + query.TplName + "%"))
		}
		if query.Status != -1 { //整数
			do = do.Where(t.Status.Eq(query.Status))
		}
	}

	//关联产品分类查询
	do = do.LeftJoin(productType, productType.Id.EqCol(t.ProductTypeId))
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = t.CreatedAt
	}
	if req.OrderDesc != "asc" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}
	do = do.Select(t.ALL, productType.Name.As("productTypeName"))
	var list []*struct {
		model.TTplTestcaseTemplate
		ProductTypeName string `gorm:"column:productTypeName" json:"productTypeName"`
	}
	var total int64
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
		logger.Errorf("GetListTplTestcaseTemplate error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.TplTestcaseTemplate, len(list))
	for i, v := range list {
		result[i] = convert.TplTestcaseTemplate_db2pb(&v.TTplTestcaseTemplate)
		result[i].ProductTypeName = v.ProductTypeName
	}
	return result, total, nil
}
