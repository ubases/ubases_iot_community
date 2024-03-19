package services

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4/metadata"
)

type ProductTypeService struct {
	Ctx context.Context
}

func (s ProductTypeService) SetContext(ctx context.Context) ProductTypeService {
	s.Ctx = ctx
	return s
}

// GetProductTypeByApp get ProductType list  data
func (s ProductTypeService) GetProductTypeByApp(filter entitys.AppQueryProductTypeForm) (rets []*entitys.TPmProductTypeVo, total int, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	ret, err := rpc.ProductTypeService.ListTPmProductType(context.Background(), &protosService.TPmProductTypeFilterPage{
		Page:      int64(filter.Page),
		Limit:     int64(filter.Limit),
		SearchKey: filter.SearchKey,
	})

	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Msg)
	}
	var (
		dataList = make([]*entitys.TPmProductTypeVo, len(ret.List))
	)
	mapstructure.WeakDecode(ret.List, &dataList)

	langMap := make(map[string]string)
	if lang != "" {
		sourceRowIds := []string{}
		for _, data := range dataList {
			sourceRowIds = append(sourceRowIds, fmt.Sprintf("%s_%s_name", lang, data.Id))
		}
		slice, err := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_PM_PRODUCT_TYPE, sourceRowIds...).Result()
		if err == nil {
			langMap = iotutil.ArrayUnionInterfaces(sourceRowIds, slice)
		}
	}
	// parentId "0" -> ""（前端要求）
	for i, data := range dataList {
		dataList[i].Name = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_name", lang, data.Id)], dataList[i].Name)
		if dataList[i].ParentId == "0" {
			dataList[i].ParentId = ""
		}
	}
	//获取链表树
	tree := getTreeIterative(dataList, "")
	//返回结果t
	var (
		limit    = filter.Limit
		offset   = limit * (filter.Page - 1)
		indexEnd = limit + offset
	)
	total = len(tree)
	//组装返回结果
	hasPage := filter.Page != 0 || filter.Limit != 0
	if hasPage {
		if offset > total {
			return nil, total, err
		} else {
			if (offset + limit) > total {
				indexEnd = total
			}
			return tree[offset:indexEnd], total, err
		}
	} else {
		return tree, total, err
	}
}

// 获取链表树
func getTreeIterative(list []*entitys.TPmProductTypeVo, parentId string) []*entitys.TPmProductTypeVo {
	memo := make(map[string]*entitys.TPmProductTypeVo)
	for _, v := range list {
		if _, ok := memo[v.Id]; ok {
			v.Children = memo[v.Id].Children
			memo[v.Id] = v
		} else {
			v.Children = make([]*entitys.TPmProductTypeVo, 0)
			memo[v.Id] = v
		}
		if _, ok := memo[v.ParentId]; ok {
			memo[v.ParentId].Children = append(memo[v.ParentId].Children, memo[v.Id])
		} else {
			memo[v.ParentId] = &entitys.TPmProductTypeVo{Children: []*entitys.TPmProductTypeVo{memo[v.Id]}}
		}
	}
	return memo[parentId].Children
}

// GetProductTree get ProductType list  data
func (s ProductTypeService) GetProductTree(filter entitys.AppQueryProductTypeForm) (rets []*entitys.TPmProductTypeVo, total int, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	ret, err := rpc.ProductTypeService.ListTPmProductType(context.Background(), &protosService.TPmProductTypeFilterPage{})
	if err != nil {
		return nil, 0, err
	}
	if ret.Code != 200 {
		return nil, 0, errors.New(ret.Msg)
	}
	//获取已发布产品列表
	productRet, err := rpc.ProductBaseService.ListTPmProduct(context.Background(), &protosService.TPmProductFilterPage{
		Page:     1,
		Limit:    1000,
		QueryObj: &protosService.TPmProductFilter{Status: 1},
	})
	if err != nil {
		return nil, 0, err
	}
	if productRet.Code != 200 {
		return nil, 0, errors.New(productRet.Msg)
	}

	//查询翻译
	typeLangMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_PM_PRODUCT_TYPE).Result()
	if typeLangMap == nil {
		typeLangMap = map[string]string{}
	}
	//分类翻译
	for i, v := range ret.List {
		langKey := fmt.Sprintf("%s_%v_name", lang, v.Id)
		ret.List[i].Name = iotutil.MapGetStringVal(typeLangMap[langKey], v.Name)
	}
	//查询翻译
	proLangMap, _ := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_NAME).Result()
	if proLangMap == nil {
		proLangMap = map[string]string{}
	}

	pVos := make(map[int64][]*entitys.TPmProductVo)
	if len(productRet.List) > 0 {
		list := productRet.List
		for _, request := range list {
			if request.ProductTypeId == 0 {
				continue
			}
			data := &entitys.TPmProductVo{
				Id:               iotutil.ToString(request.Id),
				ProductTypeId:    iotutil.ToString(request.ProductTypeId),
				ProductKey:       request.ProductKey,
				ImageURL:         request.ImageUrl,
				Model:            request.Model,
				WifiFlag:         request.WifiFlag,
				NetworkType:      iotutil.ToString(request.NetworkType),
				AttributeType:    request.AttributeType,
				Status:           request.Status,
				IsVirtualTest:    request.IsVirtualTest,
				Desc:             request.Desc,
				ProductTypeName:  request.ProductTypeName,
				PowerConsumeType: request.PowerConsumeType,
			}
			data.ImageURL = controls.ConvertProImg(data.ImageURL)
			langKey := fmt.Sprintf("%s_%v_name", lang, request.Id)
			request.Name = iotutil.MapGetStringVal(proLangMap[langKey], request.Name)
			mapstructure.WeakDecode(request, data)
			if _, ok := pVos[request.ProductTypeId]; ok {
				pVos[request.ProductTypeId] = append(pVos[request.ProductTypeId], data)
			} else {
				pVos[request.ProductTypeId] = []*entitys.TPmProductVo{data}
			}
		}
	}
	//获取产品分类链表树
	tree := s.convertTreeData(lang, "0", ret.List, pVos)
	return tree, total, err
}

// GetProductTreeV3 get ProductType list  data
func (s ProductTypeService) GetProductTreeV3(filter entitys.AppQueryProductTypeForm) (rets []*entitys.TPmProductTypeVo, total int, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	ret, err := rpc.ProductTypeService.ListTPmProductType(context.Background(), &protosService.TPmProductTypeFilterPage{})
	if err != nil {
		return nil, 0, err
	}
	if ret.Code != 200 {
		return nil, 0, errors.New(ret.Msg)
	}
	//获取已发布产品列表
	productRet, err := rpc.ProductBaseService.ListTPmProduct(context.Background(), &protosService.TPmProductFilterPage{
		Page:     1,
		Limit:    1000,
		QueryObj: &protosService.TPmProductFilter{Status: 1},
	})
	if err != nil {
		return nil, 0, err
	}
	if productRet.Code != 200 {
		return nil, 0, errors.New(productRet.Msg)
	}
	//查询翻译
	typeLangMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_PM_PRODUCT_TYPE).Result()
	if typeLangMap == nil {
		typeLangMap = map[string]string{}
	}
	//分类翻译
	for i, v := range ret.List {
		langKey := fmt.Sprintf("%s_%v_name", lang, v.Id)
		ret.List[i].Name = iotutil.MapGetStringVal(typeLangMap[langKey], v.Name)
	}
	//查询翻译
	proLangMap, _ := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_NAME).Result()
	if proLangMap == nil {
		proLangMap = map[string]string{}
	}
	pVos := make(map[int64][]*entitys.TPmProductVo)
	if len(productRet.List) > 0 {
		list := productRet.List
		for _, request := range list {
			if request.ProductTypeId == 0 {
				continue
			}
			data := &entitys.TPmProductVo{
				Id:               iotutil.ToString(request.Id),
				ProductTypeId:    iotutil.ToString(request.ProductTypeId),
				ProductKey:       request.ProductKey,
				ImageURL:         request.ImageUrl,
				Model:            request.Model,
				WifiFlag:         request.WifiFlag,
				NetworkType:      iotutil.ToString(request.NetworkType),
				AttributeType:    request.AttributeType,
				Status:           request.Status,
				IsVirtualTest:    request.IsVirtualTest,
				Desc:             request.Desc,
				ProductTypeName:  request.ProductTypeName,
				PowerConsumeType: request.PowerConsumeType,
			}
			langKey := fmt.Sprintf("%s_%v_name", lang, request.Id)
			request.Name = iotutil.MapGetStringVal(proLangMap[langKey], request.Name)
			mapstructure.WeakDecode(request, data)
			if _, ok := pVos[request.ProductTypeId]; ok {
				pVos[request.ProductTypeId] = append(pVos[request.ProductTypeId], data)
			} else {
				pVos[request.ProductTypeId] = []*entitys.TPmProductVo{data}
			}
		}
	}
	//获取产品分类链表树
	tree := s.convertTreeData(lang, "0", ret.List, pVos)
	return tree, total, err
}

func (s ProductTypeService) convertTreeData(lang string, pid string, areaList []*protosService.TPmProductTypeRequest, pVos map[int64][]*entitys.TPmProductVo) []*entitys.TPmProductTypeVo {

	treeList := func() []*entitys.TPmProductTypeVo {
		treeList := []*entitys.TPmProductTypeVo{}
		flatPtr := []*entitys.TPmProductTypeVo{}
		for _, src := range areaList {
			data := &entitys.TPmProductTypeVo{
				Id:          iotutil.ToString(src.Id),
				Identifier:  src.Identifier,
				Sort:        src.Sort,
				Name:        src.Name,
				ParentId:    iotutil.ToString(src.ParentId),
				Desc:        src.Desc,
				ImgFullPath: src.ImgFullPath,
				ParentName:  src.ParentName,
			}
			if pVos[iotutil.ToInt64(data.Id)] != nil {
				data.Products = pVos[iotutil.ToInt64(data.Id)]
			}
			flatPtr = append(flatPtr, data)
		}
		for m := range flatPtr {
			for n := range flatPtr {
				if flatPtr[m].Id == flatPtr[n].ParentId {
					if flatPtr[m].Children == nil {
						flatPtr[m].Children = []*entitys.TPmProductTypeVo{}
					}
					flatPtr[m].Children = append(flatPtr[m].Children, flatPtr[n])
				}
			}
		}
		//这里只考虑两级
		for _, j := range flatPtr {
			if j.ParentId == pid {
				if j.Children != nil && len(j.Children) > 0 {
					newChildren := []*entitys.TPmProductTypeVo{}
					for _, c := range j.Children {
						if c.Products == nil || len(c.Products) == 0 {
							continue
						}
						newChildren = append(newChildren, c)
					}
					treeList = append(treeList, newChildren...)
				}
			}
		}
		return treeList
	}()
	return treeList
}
