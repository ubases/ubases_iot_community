// Code generated by sgen.exe,2022-05-06 14:01:20. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"context"
	"errors"

	"go-micro.dev/v4/logger"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/orm"
	"cloud_platform/iot_product_service/convert"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpmNetworkGuideSvc struct {
	Ctx context.Context
}

// 创建OpmNetworkGuide
func (s *OpmNetworkGuideSvc) SaveOpmNetworkGuide(req *proto.OpmNetworkGuide) (*proto.OpmNetworkGuide, error) {
	var err error
	tenantId, err := CheckTenantId(s.Ctx)
	if err != nil {
		return nil, err
	}
	userIdStr, err := GetUserId(s.Ctx)
	if err != nil {
		return nil, err
	}
	userId := iotutil.ToInt64(userIdStr)
	if userId == 0 {
		return nil, errors.New("未获取到登录用户编号")
	}

	//校验参数
	if req.ProductId == 0 {
		return nil, errors.New("产品编号不能为空")
	}
	if req.Steps == nil || len(req.Steps) == 0 {
		return nil, errors.New("配网步骤不能为空")
	}
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		t := q.TOpmNetworkGuide
		tStep := q.TOpmNetworkGuideStep
		do := t.WithContext(context.Background())
		doStep := tStep.WithContext(context.Background())

		//删除之前的配置
		networkGuids, err := do.Where(t.ProductId.Eq(req.ProductId), t.Type.Eq(req.Type)).Find()
		//for _, guid := range networkGuids {
		//	_, _ = do.Where(t.Id.Eq(guid.Id)).Delete()
		//	_, _ = doStep.Where(tStep.NetworkGuideId.Eq(guid.Id)).Delete()
		//}
		networGuide := &model.TOpmNetworkGuide{
			ProductId: req.ProductId,
			Type:      req.Type,
			CreatedBy: userId,
			UpdatedBy: userId,
			TenantId:  tenantId,
		}
		if len(networkGuids) > 0 {
			networGuide.Id = networkGuids[0].Id
		} else {
			networGuide.Id = iotutil.GetNextSeqInt64()
		}
		err = do.Save(networGuide)
		if err != nil {
			return err
		}

		notDeleteStepIds := []int64{}
		for _, step := range req.Steps {
			if step.Id != 0 {
				notDeleteStepIds = append(notDeleteStepIds, step.Id)
			}
		}
		//删除没包含在本次提交的数据
		if len(notDeleteStepIds) > 0 {
			doStep.Where(tStep.NetworkGuideId.Eq(networGuide.Id), tStep.Id.NotIn(notDeleteStepIds...)).Delete()
		}
		//保存步骤数据（新增/修改）
		for _, step := range req.Steps {
			tempStep := &model.TOpmNetworkGuideStep{
				NetworkGuideId: networGuide.Id,
				Instruction:    step.Instruction,
				InstructionEn:  step.InstructionEn,
				ImageUrl:       step.ImageUrl,
				VideoUrl:       step.VideoUrl,
				Sort:           step.Sort,
				CreatedBy:      userId,
				UpdatedBy:      userId,
				TenantId:       tenantId,
			}
			if step.Id == 0 {
				tempStep.Id = iotutil.GetNextSeqInt64()
			} else {
				tempStep.Id = step.Id
			}
			err = doStep.Save(tempStep)
			if err != nil {
				break
			}
		}
		return err
	})
	if err != nil {
		logger.Errorf("CreateOpmNetworkGuide error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 创建OpmNetworkGuide
func (s *OpmNetworkGuideSvc) SetNetworkGuideTypes(req *proto.SetNetworkGuideTypeRequest) error {
	var err error
	tenantId, err := CheckTenantId(s.Ctx)
	if err != nil {
		return err
	}
	userIdStr, err := GetUserId(s.Ctx)
	if err != nil {
		return err
	}
	userId := iotutil.ToInt64(userIdStr)
	if userId == 0 {
		return errors.New("未获取到登录用户编号")
	}

	//校验参数
	if req.ProductId == 0 {
		return errors.New("产品编号不能为空")
	}
	if req.Type == nil || len(req.Type) == 0 {
		return errors.New("配网步骤不能为空")
	}
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		t := q.TOpmNetworkGuide
		tStep := q.TOpmNetworkGuideStep
		do := t.WithContext(context.Background())
		doStep := tStep.WithContext(context.Background())
		//删除之前的配置
		productInfo, err := q.TOpmProduct.WithContext(context.Background()).Where(q.TOpmProduct.Id.Eq(req.ProductId)).First()
		if err != nil {
			return err
		}

		resMap := make(map[int32]*proto.OpmNetworkGuide)
		do = do.Where(t.ProductId.Eq(req.ProductId))
		//do = do.Where(t.TenantId.Eq(tenantId))
		networkGuideList, err := do.Find()
		if err != nil {
			logger.Errorf("SetNetworkGuideTypes error : %s", err.Error())
			return err
		}
		networkGuidIds := make([]int64, 0)
		for _, guide := range networkGuideList {
			resMap[guide.Type] = &proto.OpmNetworkGuide{
				Id:        guide.Id,
				ProductId: guide.ProductId,
				Type:      guide.Type,
				Steps:     []*proto.OpmNetworkGuideStep{},
			}
			networkGuidIds = append(networkGuidIds, guide.Id)
			if !iotutil.ArraysExistsInt32(req.Type, guide.Type) {
				//doBase.Where(tBase.Id.Eq(guide.Id)).Delete()
				//删除不存在的配网引导类型
				do.Where(t.Type.Eq(guide.Type)).Delete()
			} else {
				//将需要存在的配网引导类型移除
				req.Type = iotutil.DeleteInt32Element(req.Type, guide.Type)
			}
		}

		for _, t := range req.Type {
			if _, ok := resMap[t]; !ok {
				networGuide := &model.TOpmNetworkGuide{
					Id:        iotutil.GetNextSeqInt64(),
					ProductId: req.ProductId, //通过基础品类编号获取配网引导
					Type:      t,
					CreatedBy: userId,
					UpdatedBy: userId,
					TenantId:  tenantId,
				}
				err = do.Create(networGuide)
				if err != nil {
					return err
				}

				defaultNetworkGuide, err := s.FindDefaultNetworkGuideByBaseProductId(&proto.OpmNetworkGuideFilter{
					ProductId: productInfo.Id,
					Type:      t,
					//ProductKey: productInfo,
				})
				if err != nil {
					return err
				}
				if len(defaultNetworkGuide) == 0 {
					continue
				}
				saveNetworkGuide := defaultNetworkGuide[0]
				saveStepList := []*model.TOpmNetworkGuideStep{}
				for _, step := range saveNetworkGuide.Steps {
					saveStepList = append(saveStepList, &model.TOpmNetworkGuideStep{
						Id:             iotutil.GetNextSeqInt64(),
						NetworkGuideId: networGuide.Id,
						Instruction:    step.Instruction,
						InstructionEn:  step.InstructionEn,
						ImageUrl:       step.ImageUrl,
						VideoUrl:       step.VideoUrl,
						Sort:           step.Sort,
						CreatedBy:      userId,
						UpdatedBy:      userId,
						TenantId:       tenantId,
					})
				}
				if len(saveStepList) > 0 {
					err = doStep.Create(saveStepList...)
				}
			}
		}
		return err
	})
	if err != nil {
		logger.Errorf("CreateOpmNetworkGuide error : %s", err.Error())
		return err
	}
	return nil
}

// 通过产品编号进行删除
func (s *OpmNetworkGuideSvc) DeleteOpmNetworkGuideByProductId(req *proto.OpmNetworkGuideFilter) error {
	tenantId, err := CheckTenantId(s.Ctx)
	if err != nil {
		return err
	}
	if req.ProductId == 0 {
		return errors.New("产品编号不能为空")
	}

	t := orm.Use(iotmodel.GetDB()).TOpmNetworkGuide
	do := t.WithContext(context.Background())
	do = do.Where(t.ProductId.Eq(req.ProductId))
	do = do.Where(t.TenantId.Eq(tenantId))

	_, err = do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOpmNetworkGuide error : %s", err.Error())
		return err
	}
	return err
}

// 通过产品编号获取配网引导信息
func (s *OpmNetworkGuideSvc) FindOpmNetworkGuideByProductId(req *proto.OpmNetworkGuideFilter) ([]*proto.OpmNetworkGuide, error) {
	//tenantId, err := CheckTenantId(s.Ctx)
	//if err != nil {
	//	return nil, err
	//}

	q := orm.Use(iotmodel.GetDB())

	product, err := q.TOpmProduct.WithContext(context.Background()).Where(q.TOpmProduct.Id.Eq(req.Id)).FirstOrInit()
	if err != nil || product.ProductKey == "" {
		return nil, errors.New("未获取到产品信息")
	}
	if product.BaseProductId == 0 {
		return nil, errors.New("产品编号不能为空")
	}
	t := q.TOpmNetworkGuide
	tStep := q.TOpmNetworkGuideStep
	do := t.WithContext(context.Background())
	doStep := tStep.WithContext(context.Background())

	do = do.Where(t.ProductId.Eq(product.Id))
	//do = do.Where(t.TenantId.Eq(tenantId))
	//TODO 可优化为直接关联步骤表查询
	list, err := do.Order(t.Type).Find()
	networkGuidIds := make([]int64, 0)
	resMap := make(map[int64]*proto.OpmNetworkGuide)
	resList := []*proto.OpmNetworkGuide{}
	if err == nil {
		//logger.Errorf("FindByIdOpmNetworkGuide error : %s", err.Error())
		//return nil, err

		for _, guide := range list {
			resMap[guide.Id] = convert.OpmNetworkGuide_db2pb(guide)
			networkGuidIds = append(networkGuidIds, guide.Id)
		}
		if len(networkGuidIds) > 0 {
			steps, err := doStep.LeftJoin(t, t.Id.EqCol(tStep.NetworkGuideId)).
				Where(tStep.NetworkGuideId.In(networkGuidIds...)).
				Order(tStep.Sort).Find()
			if err != nil {
				logger.Errorf("FindByIdOpmNetworkGuide steps error : %s", err.Error())
				return nil, err
			}
			for _, step := range steps {
				if _, ok := resMap[step.NetworkGuideId]; ok {
					resMap[step.NetworkGuideId].Steps = append(resMap[step.NetworkGuideId].Steps, convert.OpmNetworkGuideStep_db2pb(step))
				}
			}
			for _, row := range resMap {
				resList = append(resList, row)
			}
		}
	}
	if len(resList) == 0 {
		//todo 迁移到单独方法
		tBase := q.TPmNetworkGuide
		tStepBase := q.TPmNetworkGuideStep
		doBase := tBase.WithContext(context.Background())
		doStepBase := tStepBase.WithContext(context.Background())

		doBase = doBase.Where(tBase.ProductId.Eq(product.BaseProductId))
		//do = do.Where(t.TenantId.Eq(tenantId))
		list, err := doBase.Find()
		if err != nil {
			logger.Errorf("FindByIdOpmNetworkGuide error : %s", err.Error())
			return nil, err
		}

		networkGuidIds := make([]int64, 0)
		for _, guide := range list {
			resMap[guide.Id] = &proto.OpmNetworkGuide{
				Id:        guide.Id,
				ProductId: guide.ProductId,
				Type:      guide.Type,
				Steps:     []*proto.OpmNetworkGuideStep{},
			}
			networkGuidIds = append(networkGuidIds, guide.Id)
		}
		if len(networkGuidIds) > 0 {
			steps, err := doStepBase.Where(tStepBase.NetworkGuideId.In(networkGuidIds...)).Order(tStepBase.Sort).Find()
			if err != nil {
				logger.Errorf("FindByIdOpmNetworkGuide steps error : %s", err.Error())
				return nil, err
			}
			for _, step := range steps {
				if _, ok := resMap[step.NetworkGuideId]; ok {
					resMap[step.NetworkGuideId].Steps = append(resMap[step.NetworkGuideId].Steps, &proto.OpmNetworkGuideStep{
						Id:             step.Id,
						NetworkGuideId: step.NetworkGuideId,
						Instruction:    step.Instruction,
						InstructionEn:  step.InstructionEn,
						ImageUrl:       step.ImageUrl,
						VideoUrl:       step.VideoUrl,
						Sort:           step.Sort,
					})
				}
			}
			for _, row := range resMap {
				resList = append(resList, row)
			}
		}
	}
	return resList, err
}

// 通过产品编号获取配网引导信息
func (s *OpmNetworkGuideSvc) FindDefaultNetworkGuideByBaseProductId(req *proto.OpmNetworkGuideFilter) ([]*proto.OpmNetworkGuide, error) {
	//tenantId, err := CheckTenantId(s.Ctx)
	//if err != nil {
	//	return nil, err
	//}

	q := orm.Use(iotmodel.GetDB())
	productInfo, err := q.TOpmProduct.WithContext(context.Background()).Where(q.TOpmProduct.Id.Eq(req.ProductId)).FirstOrInit()
	if err != nil {
		return nil, errors.New("未获取到产品信息")
	}
	if productInfo.Id == 0 {
		return nil, errors.New("产品编号不能为空")
	}

	resList := []*proto.OpmNetworkGuide{}
	resMap := make(map[int64]*proto.OpmNetworkGuide)
	tBase := q.TPmNetworkGuide
	tStepBase := q.TPmNetworkGuideStep
	doBase := tBase.WithContext(context.Background())
	doStepBase := tStepBase.WithContext(context.Background())

	doBase = doBase.Where(tBase.ProductId.Eq(productInfo.BaseProductId), tBase.Type.Eq(req.Type))
	//do = do.Where(t.TenantId.Eq(tenantId))
	list, err := doBase.Find()
	if err != nil {
		logger.Errorf("FindByIdOpmNetworkGuide error : %s", err.Error())
		return nil, err
	}

	networkGuidIds := make([]int64, 0)
	for _, guide := range list {
		resMap[guide.Id] = &proto.OpmNetworkGuide{
			Id:        iotutil.GetNextSeqInt64(),
			ProductId: productInfo.Id,
			Type:      guide.Type,
			Steps:     []*proto.OpmNetworkGuideStep{},
		}
		networkGuidIds = append(networkGuidIds, guide.Id)
	}
	if len(networkGuidIds) > 0 {
		steps, err := doStepBase.Where(tStepBase.NetworkGuideId.In(networkGuidIds...)).Order(tStepBase.Sort).Find()
		if err != nil {
			logger.Errorf("FindByIdOpmNetworkGuide steps error : %s", err.Error())
			return nil, err
		}
		for _, step := range steps {
			if _, ok := resMap[step.NetworkGuideId]; ok {
				resMap[step.NetworkGuideId].Steps = append(resMap[step.NetworkGuideId].Steps, &proto.OpmNetworkGuideStep{
					Id:             iotutil.GetNextSeqInt64(),
					NetworkGuideId: step.NetworkGuideId,
					Instruction:    step.Instruction,
					InstructionEn:  step.InstructionEn,
					ImageUrl:       step.ImageUrl,
					VideoUrl:       step.VideoUrl,
					Sort:           step.Sort,
				})
			}
		}
	}
	for _, row := range resMap {
		resList = append(resList, row)
	}
	return resList, err
}

// 通过产品编号获取配网引导信息
func (s *OpmNetworkGuideSvc) SetDefaultNetworkGuideByBaseProductId(tx *orm.Query, productId int64, baseProductId int64, tenantId, proImg string) error {
	resList := []*proto.OpmNetworkGuide{}
	resMap := make(map[int64]*proto.OpmNetworkGuide)
	tBase := tx.TPmNetworkGuide
	tStepBase := tx.TPmNetworkGuideStep
	doBase := tBase.WithContext(context.Background())
	doStepBase := tStepBase.WithContext(context.Background())

	doBase = doBase.Where(tBase.ProductId.Eq(baseProductId))
	//do = do.Where(t.TenantId.Eq(tenantId))
	list, err := doBase.Find()
	if err != nil {
		logger.Errorf("FindByIdOpmNetworkGuide error : %s", err.Error())
		return err
	}

	networkGuidIds := make([]int64, 0)
	for _, guide := range list {
		resMap[guide.Id] = &proto.OpmNetworkGuide{
			Id:        iotutil.GetNextSeqInt64(),
			ProductId: productId,
			Type:      guide.Type,
			Steps:     []*proto.OpmNetworkGuideStep{},
		}
		networkGuidIds = append(networkGuidIds, guide.Id)
	}
	if len(networkGuidIds) > 0 {
		steps, err := doStepBase.Where(tStepBase.NetworkGuideId.In(networkGuidIds...)).Order(tStepBase.Sort).Find()
		if err != nil {
			logger.Errorf("FindByIdOpmNetworkGuide steps error : %s", err.Error())
			return err
		}
		for _, step := range steps {
			if _, ok := resMap[step.NetworkGuideId]; ok {
				resMap[step.NetworkGuideId].Steps = append(resMap[step.NetworkGuideId].Steps, &proto.OpmNetworkGuideStep{
					Id:             iotutil.GetNextSeqInt64(),
					NetworkGuideId: step.NetworkGuideId,
					Instruction:    step.Instruction,
					InstructionEn:  step.InstructionEn,
					ImageUrl:       step.ImageUrl,
					VideoUrl:       step.VideoUrl,
					Sort:           step.Sort,
				})
			}
		}
	}

	guides := make([]*model.TOpmNetworkGuide, 0)
	guideSteps := make([]*model.TOpmNetworkGuideStep, 0)
	for _, row := range resMap {
		resList = append(resList, row)
		guideId := iotutil.GetNextSeqInt64()
		guides = append(guides, &model.TOpmNetworkGuide{
			Id:        guideId,
			ProductId: row.ProductId,
			Type:      row.Type,
			TenantId:  tenantId,
		})
		for _, step := range row.Steps {
			guideSteps = append(guideSteps, &model.TOpmNetworkGuideStep{
				Id:             iotutil.GetNextSeqInt64(),
				NetworkGuideId: guideId,
				ProductId:      productId,
				Instruction:    step.Instruction,
				InstructionEn:  step.InstructionEn,
				ImageUrl:       step.ImageUrl,
				VideoUrl:       step.VideoUrl,
				Sort:           step.Sort,
				TenantId:       tenantId,
			})
		}
	}
	if len(guides) > 0 {
		tOpm := tx.TOpmNetworkGuide
		tOpmStep := tx.TOpmNetworkGuideStep
		doOpm := tOpm.WithContext(context.Background())
		doOpmStep := tOpmStep.WithContext(context.Background())
		err := doOpm.Create(guides...)
		if err != nil {
			logger.Errorf("SetDefaultNetworkGuideByBaseProductId error : %s", err.Error())
			return err
		}
		err = doOpmStep.Create(guideSteps...)
		if err != nil {
			logger.Errorf("SetDefaultNetworkGuideByBaseProductId error : %s", err.Error())
			return err
		}
	}
	return err
}
