// Code generated by sgen,2024-03-11 16:22:42. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_product_service/convert"
)

type OpmProductTestAccountSvc struct {
	Ctx context.Context
}

//创建OpmProductTestAccount
func (s *OpmProductTestAccountSvc) CreateOpmProductTestAccount(req *proto.OpmProductTestAccount) (*proto.OpmProductTestAccount, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
	do := t.WithContext(context.Background())
	dbObj := convert.OpmProductTestAccount_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOpmProductTestAccount error : %s", err.Error())
		return nil,err
	}
	return req, err
}

//根据条件删除OpmProductTestAccount
func (s *OpmProductTestAccountSvc) DeleteOpmProductTestAccount(req *proto.OpmProductTestAccount) (*proto.OpmProductTestAccount, error) {
    t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
    do := t.WithContext(context.Background())
    // fixme 请检查条件
    
    if req.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Id))
    }
    if req.ProductId != 0 {//整数
        do = do.Where(t.ProductId.Eq(req.ProductId))
    }
    if req.Account != "" {//字符串
        do = do.Where(t.Account.Eq(req.Account))
    }
	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	if req.RegionServerId != 0 {
		do = do.Where(t.RegionServerId.Eq(req.RegionServerId))
	}
    if req.UserId != 0 {//整数
        do = do.Where(t.UserId.Eq(req.UserId))
    }
    if req.CreatedBy != 0 {//整数
        do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
    }
    _,err := do.Delete()
    if err != nil {
        logger.Errorf("DeleteOpmProductTestAccount error : %s", err.Error())
		return nil,err
    }
    return req, err
}

//根据数据库表主键删除OpmProductTestAccount
func (s *OpmProductTestAccountSvc) DeleteByIdOpmProductTestAccount(req *proto.OpmProductTestAccount) (*proto.OpmProductTestAccount, error) {
    t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
    do := t.WithContext(context.Background())
    // fixme 请检查条件
    
    if req.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Id))
    }
   _,err := do.Delete()
    if err != nil {
        logger.Errorf("DeleteByIdOpmProductTestAccount error : %s", err.Error())
		return nil,err
    }
    return req, err
}

//根据数据库表主键批量删除OpmProductTestAccount
func (s *OpmProductTestAccountSvc) DeleteByIdsOpmProductTestAccount(req *proto.OpmProductTestAccountBatchDeleteRequest) (*proto.OpmProductTestAccountBatchDeleteRequest, error) {
    var err error
    for _,k := range req.Keys {
        t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
        do := t.WithContext(context.Background())
         
        do = do.Where(t.Id.Eq(k.Id))
              
         _,err = do.Delete()
         if err != nil {
             logger.Errorf("DeleteByIdsOpmProductTestAccount error : %s", err.Error())
             break
         }
    }
    return req, err
}

//根据主键更新OpmProductTestAccount
func (s *OpmProductTestAccountSvc) UpdateOpmProductTestAccount(req *proto.OpmProductTestAccount) (*proto.OpmProductTestAccount, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField  []field.Expr
    
    if req.ProductId != 0 {//整数
        updateField = append(updateField,t.ProductId)
    }
    if req.Account != "" {//字符串
        updateField = append(updateField,t.Account)
    }
	if req.TenantId != "" { //字符串
		updateField = append(updateField, t.TenantId)
	}
	if req.RegionServerId != 0 { //整数
		updateField = append(updateField, t.RegionServerId)
	}
    if req.UserId != 0 {//整数
        updateField = append(updateField,t.UserId)
    }
    if req.CreatedBy != 0 {//整数
        updateField = append(updateField,t.CreatedBy)
    }
    if len(updateField) > 0 {
        do = do.Select(updateField...)
    }
    //主键条件
    HasPrimaryKey := false
      
    if req.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Id))
        HasPrimaryKey = true
    }      

    if !HasPrimaryKey {
        logger.Error("UpdateOpmProductTestAccount error : Missing condition")
        return nil,errors.New("Missing condition")
    }

	dbObj := convert.OpmProductTestAccount_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOpmProductTestAccount error : %s", err.Error())
		return nil,err
	}
	return req, err
}

////根据主键更新所有字段OpmProductTestAccount
func (s *OpmProductTestAccountSvc) UpdateAllOpmProductTestAccount(req *proto.OpmProductTestAccount) (*proto.OpmProductTestAccount, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField  []field.Expr
      
	updateField = append(updateField,t.ProductId)
	updateField = append(updateField,t.Account)
	updateField = append(updateField,t.UserId)
	updateField = append(updateField,t.AppKey)
	updateField = append(updateField, t.RegionServerId)
	updateField = append(updateField, t.TenantId)
	updateField = append(updateField,t.CreatedBy)
    if len(updateField) > 0 {
        do = do.Select(updateField...)
    }
    //主键条件
    HasPrimaryKey := false   
    if req.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Id))
        HasPrimaryKey = true
    }      
    if !HasPrimaryKey {
        logger.Error("UpdateAllOpmProductTestAccount error : Missing condition")
        return nil,errors.New("Missing condition")
    }
	dbObj := convert.OpmProductTestAccount_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOpmProductTestAccount error : %s", err.Error())
		return nil,err
	}
	return req, err
}

func (s *OpmProductTestAccountSvc) UpdateFieldsOpmProductTestAccount(req *proto.OpmProductTestAccountUpdateFieldsRequest) (*proto.OpmProductTestAccount, error) {
    t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
	do := t.WithContext(context.Background())

	var updateField  []field.Expr
	for _,v:=range req.Fields {
	    col, ok := t.GetFieldByName(v)
        if ok {
          updateField = append(updateField,col)
        }
    }
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOpmProductTestAccount error : missing updateField")
		logger.Error(err)
		return nil,err
	}
	do = do.Select(updateField...)

    //主键条件
    HasPrimaryKey := false   
    if req.Data.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Data.Id))
        HasPrimaryKey = true
    }      
    if !HasPrimaryKey {
        logger.Error("UpdateFieldsOpmProductTestAccount error : Missing condition")
        return nil,errors.New("Missing condition")
    }
	dbObj := convert.OpmProductTestAccount_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOpmProductTestAccount error : %s", err.Error())
		return nil,err
	}
	return req.Data, nil
}

//根据非空条件查找OpmProductTestAccount
func (s *OpmProductTestAccountSvc) FindOpmProductTestAccount(req *proto.OpmProductTestAccountFilter) (*proto.OpmProductTestAccount, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
    
    if req.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Id))
    }
    if req.ProductId != 0 {//整数
        do = do.Where(t.ProductId.Eq(req.ProductId))
    }
    if req.Account != "" {//字符串
        do = do.Where(t.Account.Eq(req.Account))
    }
	if req.AppKey != "" {//字符串
		do = do.Where(t.AppKey.Eq(req.AppKey))
	}
    if req.UserId != 0 {//整数
        do = do.Where(t.UserId.Eq(req.UserId))
    }
    if req.CreatedBy != 0 {//整数
        do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
    }
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpmProductTestAccount error : %s", err.Error())
		return nil,err
	}
	res := convert.OpmProductTestAccount_db2pb(dbObj)
	return res, err
}

//根据数据库表主键查找OpmProductTestAccount
func (s *OpmProductTestAccountSvc) FindByIdOpmProductTestAccount(req *proto.OpmProductTestAccountFilter) (*proto.OpmProductTestAccount, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
	do := t.WithContext(context.Background())
    // fixme 请检查条件和校验参数
    
    if req.Id != 0 {//整数
        do = do.Where(t.Id.Eq(req.Id))
    }
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOpmProductTestAccount error : %s", err.Error())
		return nil,err
	}
	res := convert.OpmProductTestAccount_db2pb(dbObj)
	return res, err
}

//根据分页条件查找OpmProductTestAccount,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OpmProductTestAccountSvc) GetListOpmProductTestAccount(req *proto.OpmProductTestAccountListRequest) ([]*proto.OpmProductTestAccount, int64, error) {
    // fixme 请检查条件和校验参数
    var err error
    t := orm.Use(iotmodel.GetDB()).TOpmProductTestAccount
    do := t.WithContext(context.Background())
    query := req.Query
    if query != nil {
        
        if query.Id != 0 {//整数
            do = do.Where(t.Id.Eq(query.Id))
        }
        if query.ProductId != 0 {//整数
            do = do.Where(t.ProductId.Eq(query.ProductId))
        }
		if query.AppKey != "" {//字符串
			do = do.Where(t.AppKey.Eq(query.AppKey))
		}
		if query.TenantId != "" { //字符串
			do = do.Where(t.TenantId.Eq(query.TenantId))
		}
		if query.RegionServerId != 0 {
			do = do.Where(t.RegionServerId.Eq(query.RegionServerId))
		}
        if query.Account != "" {//字符串
            do = do.Where(t.Account.Like("%" + query.Account + "%"))
        }
        if query.UserId != 0 {//整数
            do = do.Where(t.UserId.Eq(query.UserId))
        }
        if query.CreatedBy != 0 {//整数
            do = do.Where(t.CreatedBy.Eq(query.CreatedBy))
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

	var list []*model.TOpmProductTestAccount
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
	    logger.Errorf("GetListOpmProductTestAccount error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
	     return nil, total, nil
	}
	result := make([]*proto.OpmProductTestAccount,len(list))
    for i, v := range list {
        result[i] = convert.OpmProductTestAccount_db2pb(v)
    }
    return result, total, nil
}


