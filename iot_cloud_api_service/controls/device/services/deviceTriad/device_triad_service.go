package deviceTriad

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	"cloud_platform/iot_cloud_api_service/controls/device/services/extract"
	"cloud_platform/iot_cloud_api_service/controls/oem/services/openData"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type IotDeviceTriadService struct {
	Ctx context.Context
}

func (s IotDeviceTriadService) SetContext(ctx context.Context) IotDeviceTriadService {
	s.Ctx = ctx
	return s
}

// 设备三元组详细
func (s IotDeviceTriadService) GetIotDeviceTriadDetail(id string) (*entitys.IotDeviceTriadEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientIotDeviceServer.FindById(context.Background(), &protosService.IotDeviceTriadFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]

	//查询产品信息
	//rpc.ClientOpmProductService.

	return entitys.IotDeviceTriad_pb2e(data), err
}

// QueryIotDeviceTriadList 设备三元组列表
func (s IotDeviceTriadService) QueryIotDeviceTriadList(filter entitys.IotDeviceTriadQuery) ([]*entitys.IotDeviceTriadEntitys, int64, error) {
	rep, err := rpc.ClientIotDeviceServer.Lists(context.Background(), &protosService.IotDeviceTriadListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.IotDeviceTriad{
			ProductId:  filter.Query.ProductId,
			ProductKey: filter.Query.ProductKey,
			IsTest:     filter.Query.IsTest,
			UseType:    filter.Query.UseType,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.IotDeviceTriadEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.IotDeviceTriad_pb2e(item))
	}
	return resultList, rep.Total, err
}

// QueryVirtualDeviceTriadList 查询虚拟设备
func (s IotDeviceTriadService) QueryVirtualDeviceTriadList(filter entitys.VirtualDeviceQuery) ([]*entitys.VirtualDeviceItem, int64, error) {
	var (
		useType int32 = 0
	)
	//如果是返回虚拟设备列表，则查询条件指定虚拟设备的userType
	if filter.IsVirtual == 1 {
		useType = iotconst.Use_Type_Device_Real_Test
	}
	//增加默认分页数量（这里默认最多返回100条数据
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = 100
	}
	rep, err := rpc.ClientIotDeviceServer.Lists(s.Ctx, &protosService.IotDeviceTriadListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.IotDeviceTriad{
			IsQueryTriadData: true,
			ProductId:        filter.ProductId,
			IsTest:           1,
			Status:           -1,
			UseType:          useType,
			//TenantId:         filter.TenantId,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	appKeyMap, _ := openData.GetAppMaps(s.Ctx, filter.TenantId)

	var resultList = []*entitys.VirtualDeviceItem{}
	for _, item := range rep.Data {
		device := &entitys.VirtualDeviceItem{
			Id:             item.Id,
			AppName:        item.AppName,
			Did:            item.Did,
			UserAccount:    item.UserAccount,
			DeviceUserName: item.UserName,
			AppKey:         item.AppKey,
			DevicePassword: item.Passward,
			RegionServerId: item.RegionServerId,
		}
		if item.UserId != 0 {
			device.UserId = iotutil.ToString(item.UserId)
		}
		//如果APPKey等于配置文件公版的APPKey,那么直接使用配置文件中的名称
		if item.AppKey == config.Global.DefaultApp.AppKey {
			device.AppName = config.Global.DefaultApp.AppName
		} else {
			device.AppName = appKeyMap[item.AppKey]
		}
		resultList = append(resultList, device)
	}
	return resultList, rep.Total, err
}

// CheckHasVirtualDeviceTriad 检查是否有虚拟设备
func (s IotDeviceTriadService) CheckHasVirtualDeviceTriad(productId int64, tenantId string, account string) (bool, error) {
	rep, err := rpc.ClientIotDeviceServer.Lists(s.Ctx, &protosService.IotDeviceTriadListRequest{
		Page:     int64(1),
		PageSize: int64(1),
		Query: &protosService.IotDeviceTriad{
			IsQueryTriadData: true,
			ProductId:        productId,
			IsTest:           1,
			Status:           -1,
			UseType:          iotconst.Use_Type_Device_Real_Test,
			TenantId:         tenantId,
			UserAccount:      account,
		},
	})
	if err != nil {
		return false, err
	}
	if rep.Code != 200 {
		return false, errors.New(rep.Message)
	}
	return len(rep.Data) > 0, nil
}

// AddIotDeviceTriad 新增设备三元组
func (s IotDeviceTriadService) AddIotDeviceTriad(req entitys.IotDeviceTriadEntitys) (string, error) {
	saveObj := entitys.IotDeviceTriad_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.CreatedAt = timestamppb.Now()
	res, err := rpc.ClientIotDeviceServer.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// GeneratorIotDeviceTriad 批量生成三元组
func (s IotDeviceTriadService) GeneratorIotDeviceTriad(req entitys.GenerateDeviceTriad, isImport bool) error {
	authQuantity, err := rpc.ClientAuthQuantityService.GetCountByTenantId(s.Ctx, &protosService.OpenAuthQuantityFilter{})
	if err != nil {
		return err
	}
	//if authQuantity.Total == 0 {
	//	return errors.New("当前账号未授权设备数量")
	//}
	//如是导入数据，不需要限制授权数量，生成三元组需要限制授权数量
	if req.PlatformCode == "" && !isImport {
		limitCount, err := s.CheckDeviceTriadNumber(req.TenantId, req.AccountType)
		if err != nil {
			return err
		}
		total := authQuantity.Total + limitCount
		currentDeviceCount, err := rpc.ClientIotDeviceServer.GetDeviceTriadCountByTenantId(s.Ctx, &protosService.IotDeviceTriadFilter{
			UseType: req.UseType,
		})
		if err != nil {
			return err
		}
		if currentDeviceCount.Data+int64(req.Number) > total {
			return errors.New("当前账号授权设备数量不足")
		}
	}
	productIdInt, _ := iotutil.ToInt64AndErr(req.ProductId)
	generate := &protosService.IotDeviceTriadGenerateRequest{
		ProductKey:       req.ProductKey,
		ProductId:        productIdInt,
		Number:           req.Number,
		Batch:            req.Batch,
		SerialNumbers:    req.SerialNumbers,
		GeneratorChannel: req.GeneratorChannel,
		UseType:          req.UseType,
		DeviceNatureKey:  req.DeviceNatureKey,
		DeviceImport:     entitys.DeviceImportDatasToDp(req.Devices),
		PlatformCode:     req.PlatformCode,
	}
	res, err := rpc.ClientIotDeviceServer.GeneratorDeviceTriad(s.Ctx, generate)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}

// CheckDeviceTriadNumber 获取设备三元组限制数量
func (s IotDeviceTriadService) CheckDeviceTriadNumber(tenantId string, accountType int32) (int64, error) {
	var limitCount int64 = 0 //10  //不需要限制了，统一有授权数量限制
	////查询用户的类型是否企业认证，和账号类型
	//resComp, err := rpc.ClientOpenCompanyService.Find(context.Background(), &protosService.OpenCompanyFilter{
	//	TenantId: tenantId,
	//})
	//if err != nil {
	//	return 0, err
	//}
	//if resComp.Code != 200 {
	//	return 0, errors.New(resComp.Message)
	//}
	//if resComp.Data != nil && len(resComp.Data) == 0 {
	//	return 0, errors.New("用户数据异常")
	//}
	////个人账号最多生成10个，
	////企业账号30个，企业认证账号100个
	////=1 企业 =2 个人
	//if accountType == 1 {
	//	limitCount = 30
	//}
	////=1 未提交 ,=2 认证中 =3 已认证
	//if resComp.Data[0].Status == 3 {
	//	limitCount = 100
	//}
	return limitCount, nil
}

// GeneratorTestIotDeviceTriad 批量测试生成三元组
func (s IotDeviceTriadService) GeneratorTestIotDeviceTriad(req entitys.GenerateDeviceTriad) error {
	var (
		limitCount int64 = 0 //10
		err        error
		tenantId   string = req.TenantId
	)
	limitCount, err = s.CheckDeviceTriadNumber(req.TenantId, req.AccountType)
	if err != nil {
		return err
	}
	//如果不是虚拟设备
	if req.UseType != iotconst.Use_Type_Device_Real_Test {
		authQuantity, err := rpc.ClientAuthQuantityService.GetCountByTenantId(s.Ctx, &protosService.OpenAuthQuantityFilter{})
		if err != nil {
			return err
		}
		limitCount = authQuantity.Total + limitCount
	} else {
		//如果是虚拟设备，限制虚拟设备增加数量限制（默认3个），实际根据配置文件配置
		limitCount = int64(iotutil.IfInt32(config.Global.Service.VirtualDeviceNumber == 0, 3, config.Global.Service.VirtualDeviceNumber))
	}
	if req.ProductId == "" {
		return errors.New("产品Id不能为空")
	}
	productIdInt := iotutil.ToInt64(req.ProductId)

	//当前账号授权设备数量不足
	countQuery := &protosService.IotDeviceTriadFilter{
		UseType:   req.UseType,
		ProductId: productIdInt,
		IsTest:    req.IsTest,
	}
	//如果是虚拟设备，增加每个appkey+产品只能有3个虚拟设备。
	if req.UseType == iotconst.Use_Type_Device_Real_Test {
		countQuery.AppKey = req.AppKey
	}
	currentDeviceCount, err := rpc.ClientIotDeviceServer.GetDeviceTriadCountByTenantId(s.Ctx, countQuery)
	if err != nil {
		return err
	}
	if currentDeviceCount.Data+int64(req.Number) > limitCount {
		if req.UseType != iotconst.Use_Type_Device_Real_Test {
			return errors.New("当前账号授权设备数量不足")
		} else {
			return errors.New("当前账号授权虚拟设备数量不足")
		}
	}

	//如果是公版APP租户Id是公版的租户Id，否则是当前用户的租户Id
	if req.AppKey == config.Global.DefaultApp.AppKey {
		tenantId = config.Global.DefaultApp.TenantId
	}

	productRes, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{Id: productIdInt})
	if err != nil {
		return err
	}
	if productRes.Code != 200 {
		return errors.New(productRes.Message)
	}
	if len(productRes.Data) == 0 {
		return errors.New("未找到产品信息")
	}

	productKey := productRes.Data[0].ProductKey
	productName := productRes.Data[0].Name
	//TODO AttributeType、DeviceNatureKey 重复了，需要清理
	req.DeviceNatureKey = productRes.Data[0].AttributeType
	//新增模式 1 批量生成新设备 2 添加已存在设备 3 虚拟设备新增
	switch req.AddMode {
	case 1:
		err = s.InsertDevice(productIdInt, productKey, &req)
	case 3:
		err := s.CreateAndBindDeviceTriad(entitys.AddAppAccountEntity{
			Account:        req.UserAccount,
			ProductId:      productIdInt,
			ProductName:    productName,
			AppKey:         iotutil.IfString(req.AppKey == "", config.Global.DefaultApp.AppKey, req.AppKey),
			DeviceId:       req.DeviceId,
			TenantId:       tenantId,
			RegionServerId: req.RegionServerId,
		}, &req)
		if err != nil {
			return err
		}
		//缓存延后刷新,不影响数据保存
		go extract.SetUserDeviceCached(s.Ctx, req)
	case 2:
		err = s.SetTestDeviceTriad(productKey, req)
	default:
		err = errors.New("操作模式错误")
	}
	return err
}

// InsertDevice 插入设备
func (s IotDeviceTriadService) InsertDevice(productIdInt int64, productKey string, req *entitys.GenerateDeviceTriad) error {
	generate := &protosService.IotDeviceTriadGenerateRequest{
		ProductKey:       productKey,
		ProductId:        productIdInt,
		Number:           req.Number,
		Batch:            req.Batch,
		SerialNumbers:    req.SerialNumbers,
		GeneratorChannel: req.GeneratorChannel,
		UseType:          req.UseType,
		DeviceNatureKey:  req.DeviceNatureKey,
		IsTest:           req.IsTest,
	}
	res, err := rpc.ClientIotDeviceServer.GeneratorDeviceTriad(s.Ctx, generate)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// SetTestDeviceTriad 修改测试设备
func (s IotDeviceTriadService) SetTestDeviceTriad(productKey string, req entitys.GenerateDeviceTriad) error {
	res, err := rpc.ClientIotDeviceServer.SetTestDeviceTriad(s.Ctx, &protosService.SetTestTriadRequest{
		IsTest:     1,
		DeviceId:   req.DeviceId,
		ProductKey: productKey,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// UpdateIotDeviceTriad 修改设备三元组
func (s IotDeviceTriadService) UpdateIotDeviceTriad(req entitys.IotDeviceTriadEntitys) (string, error) {
	res, err := rpc.ClientIotDeviceServer.Update(context.Background(), entitys.IotDeviceTriad_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteIotDeviceTriad 删除设备三元组
func (s IotDeviceTriadService) DeleteIotDeviceTriad(req entitys.IotDeviceTriadFilter) error {
	triadResp, err := rpc.ClientIotDeviceServer.FindById(context.Background(), &protosService.IotDeviceTriadFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if triadResp.Code != 200 {
		return errors.New(triadResp.Message)
	}

	homeDevResp, err := rpc.IotDeviceHomeService.Lists(s.Ctx, &protosService.IotDeviceHomeListRequest{
		Query: &protosService.IotDeviceHome{
			DeviceId: triadResp.Data[0].Did,
		},
	})
	if err != nil {
		return err
	}
	if homeDevResp.Code != 200 {
		return errors.New(homeDevResp.Message)
	}

	//如果是真是设备，则不容许直接删除，只是解除测试标识
	if req.UseType == 1 {
		//如果是真实设备，则移除设备得测试标识
		rep, err := rpc.ClientIotDeviceServer.SetTestDeviceTriad(s.Ctx, &protosService.SetTestTriadRequest{
			DeviceId:     req.Did,
			DeviceDataId: req.Id,
			ProductKey:   req.ProductKey,
			IsTest:       0,
		})
		if err != nil {
			return err
		}
		if rep.Code != 200 {
			return errors.New(rep.Message)
		}
	} else {
		rep, err := rpc.ClientIotDeviceServer.Delete(context.Background(), &protosService.IotDeviceTriad{
			Id:  iotutil.ToInt64(req.Id),
			Did: req.Did,
		})
		if err != nil {
			return err
		}
		if rep.Code != 200 {
			return errors.New(rep.Message)
		}
	}

	// 删除家庭详情缓存
	keys := []string{}
	homeIds := []int64{}
	for i := range homeDevResp.Data {
		homeIds = append(homeIds, homeDevResp.Data[i].HomeId)
	}

	for i := range homeIds {
		resp, err := rpc.UcHomeUserService.Lists(s.Ctx, &protosService.UcHomeUserListRequest{
			Query: &protosService.UcHomeUser{
				HomeId: homeIds[i],
			},
		})
		if err != nil {
			return err
		}
		if resp.Code != 200 {
			return errors.New(resp.Message)
		}
		for j := range resp.Data {
			keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(resp.Data[j].UserId)),
				fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeIds[i]), iotutil.ToString(resp.Data[j].UserId)))
		}
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(s.Ctx, keys...)
		if _, err := pipe.Exec(s.Ctx); err != nil {
			return err
		}
	}
	return nil
}

// SetStatusIotDeviceTriad 禁用/启用设备三元组
func (s IotDeviceTriadService) SetStatusIotDeviceTriad(req entitys.IotDeviceTriadFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientIotDeviceServer.UpdateFields(context.Background(), &protosService.IotDeviceTriadUpdateFieldsRequest{
		Fields: []string{"activeStatus"},
		Data: &protosService.IotDeviceTriad{
			Id:     iotutil.ToInt64(req.Id),
			Status: req.Status,
		},
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// AddAppAccount  添加APP账号
func (s IotDeviceTriadService) AddAppAccount(req entitys.AddAppAccountEntity, triadReq *entitys.GenerateDeviceTriad) error {
	//查询用户信息
	userRes, err := rpc.UcUserService.Find(s.Ctx, &protosService.UcUserFilter{
		UserName:       req.Account,
		TenantId:       req.TenantId,
		AppKey:         req.AppKey,
		RegionServerId: req.RegionServerId, //默认中国地区
	})
	if err != nil {
		return err
	}
	if userRes.Code != 200 {
		return errors.New("该账号尚未在APP中注册")
	}
	user := userRes.Data[0]
	homeId, err := iotutil.ToInt64AndErr(user.DefaultHomeId)
	if err != nil {
		return err
	}
	//查询产品信息
	proRes, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{
		Id: req.ProductId,
	})
	if err != nil {
		return err
	}
	if proRes.Code != 200 {
		return errors.New(proRes.Message)
	}
	pro := proRes.Data[0]

	appInfo, err := extract.GetAppInfo(s.Ctx, req.AppKey)
	if err != nil {
		return err
	}

	if triadReq != nil {
		err = s.InsertDevice(pro.Id, pro.ProductKey, triadReq)
		if err != nil {
			return err
		}
	}

	homeDevResp, err := rpc.IotDeviceHomeService.Lists(s.Ctx, &protosService.IotDeviceHomeListRequest{
		Query: &protosService.IotDeviceHome{
			DeviceId: req.DeviceId,
		},
	})
	if err != nil {
		return err
	}
	if homeDevResp.Code != 200 {
		return errors.New(homeDevResp.Message)
	}

	rep, err := rpc.ClientIotDeviceServer.BindTestAccount(s.Ctx, &protosService.BindTestAccountRequest{
		AppKey:         appInfo.AppKey,
		UserId:         user.Id,
		ProductKey:     pro.ProductKey,
		DeviceId:       req.DeviceId,
		HomeId:         homeId,
		ProductId:      pro.Id,
		ProductName:    pro.Name,
		UserAccount:    req.Account,
		AppName:        appInfo.Name,
		TenantId:       req.TenantId,
		RegionServerId: req.RegionServerId,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}

	// 删除家庭详情缓存
	keys := []string{}
	homeIds := []int64{homeId}
	for i := range homeDevResp.Data {
		homeIds = append(homeIds, homeDevResp.Data[i].HomeId)
	}

	for i := range homeIds {
		resp, err := rpc.UcHomeUserService.Lists(s.Ctx, &protosService.UcHomeUserListRequest{
			Query: &protosService.UcHomeUser{
				HomeId: homeIds[i],
			},
		})
		if err != nil {
			return err
		}
		if resp.Code != 200 {
			return errors.New(resp.Message)
		}
		for j := range resp.Data {
			keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(resp.Data[j].UserId)),
				fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeIds[i]), iotutil.ToString(resp.Data[j].UserId)))
		}
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(s.Ctx, keys...)
		if _, err := pipe.Exec(s.Ctx); err != nil {
			return err
		}
	}
	return nil
}

// CreateAndBindDeviceTriad  添加APP账号并设置虚拟设备
func (s IotDeviceTriadService) CreateAndBindDeviceTriad(req entitys.AddAppAccountEntity, triadReq *entitys.GenerateDeviceTriad) error {
	//查询用户信息
	userRes, err := rpc.UcUserService.Find(s.Ctx, &protosService.UcUserFilter{
		UserName:       req.Account,
		TenantId:       req.TenantId,
		AppKey:         req.AppKey,
		RegionServerId: req.RegionServerId, //默认中国地区
	})
	if err != nil {
		return err
	}
	if userRes.Code != 200 {
		return errors.New("该账号尚未在APP中注册")
	}
	user := userRes.Data[0]
	homeId, err := iotutil.ToInt64AndErr(user.DefaultHomeId)
	if err != nil {
		return err
	}
	//查询产品信息
	proRes, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{
		Id: req.ProductId,
	})
	if err != nil {
		return err
	}
	if proRes.Code != 200 {
		return errors.New(proRes.Message)
	}
	pro := proRes.Data[0]

	appInfo, err := extract.GetAppInfo(s.Ctx, req.AppKey)
	if err != nil {
		return err
	}
	generate := &protosService.IotDeviceTriadGenerateRequest{
		ProductId:  pro.Id,
		ProductKey: pro.ProductKey,
		UseType:    iotconst.Use_Type_Device_Real_Test,
		IsTest:     1,
		BindInfo: &protosService.BindTestAccountRequest{
			AppKey:         req.AppKey,
			UserId:         user.Id,
			ProductKey:     pro.ProductKey,
			DeviceId:       req.DeviceId,
			HomeId:         homeId,
			ProductId:      pro.Id,
			ProductName:    pro.Name,
			UserAccount:    req.Account,
			AppName:        appInfo.Name,
			TenantId:       req.TenantId,
			RegionServerId: req.RegionServerId,
		},
	}
	rep, err := rpc.ClientIotDeviceServer.CreateAndBindDeviceTriad(s.Ctx, generate)
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}

	return nil
}

// QueryVirtualDeviceList 设备三元组列表
func (s IotDeviceTriadService) QueryVirtualDeviceList(filter entitys.IotDeviceTriadQuery) ([]*entitys.VirtualDeviceItem, int64, error) {
	rep, err := rpc.ClientIotDeviceServer.Lists(context.Background(), &protosService.IotDeviceTriadListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.IotDeviceTriad{
			ProductId:  filter.Query.ProductId,
			ProductKey: filter.Query.ProductKey,
			IsTest:     1,
			UseType:    filter.Query.UseType,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.VirtualDeviceItem{}
	for _, item := range rep.Data {
		device := &entitys.VirtualDeviceItem{
			Id: item.Id,
			//AppName:     "",
			Did: item.Did,
			//UserAccount: item.UserId,
			DevicePassword: item.Passward,
			DeviceUserName: item.UserName,
		}
		if item.UserId != 0 {
			device.UserId = iotutil.ToString(item.UserId)
		}
		resultList = append(resultList, device)
	}
	return resultList, rep.Total, err
}
