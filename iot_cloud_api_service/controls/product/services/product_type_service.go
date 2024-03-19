package services

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

type ProductTypeService struct {
}

// CreateProductType create one record
func (s ProductTypeService) CreateProductType(req *entitys.CreateProductTypeForm) (ret int64, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	data, err := req.ToPB()
	if err != nil {
		return 0, err
	}
	//参数填充，生成主键ID/标识符
	data.Identifier = iotutil.Uuid()
	data.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
	data.UpdatedTime = time.Now().Format("2006-01-02 15:04:05")
	res, err := rpc.ClientProductTypeService.CreateTPmProductType(context.Background(), data)

	if err != nil {
		return 0, errors.New(err.Error())
	}
	if res.Code != 200 {
		return 0, errors.New(res.Msg)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_product_type", res.Data.Id, "name", req.Name, req.NameEn)
	if err := cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA); err != nil {
		return 0, err
	}
	return res.Data.Id, err
}

// UpdateProductType edit ProductType one record
func (s ProductTypeService) UpdateProductType(req *entitys.UpProductTypeForm) (err error) {
	if err = req.Valid(); err != nil {
		return
	}
	data, err := req.ToPB()
	if err != nil {
		return err
	}
	//data.Id = iotutil.ToInt64(req.Id)
	data.UpdatedTime = time.Now().Format("2006-01-02 15:04:05")

	res, err := rpc.ClientProductTypeService.UpdateTPmProductType(context.Background(), data)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Msg)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_product_type", req.Id, "name", req.Name, req.NameEn)
	if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.PRODUCT_TYPE_ID_DATA, req.Id)); err != nil {
		return err
	}
	if err := cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA); err != nil {
		return err
	}
	return err
}

// GetProductTypeBatch get ProductType list  data
func (s ProductTypeService) GetProductTypeList(filter *entitys.QueryProductTypeForm) (rets []*entitys.TPmProductTypeVo, total int, err error) {
	ret, err := rpc.ClientProductTypeService.ListTPmProductType(context.Background(), &protosService.TPmProductTypeFilterPage{
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

	// parentId "0" -> ""（前端要求）
	for _, data := range dataList {
		if data.ParentId == "0" {
			data.ParentId = ""
		}
	}
	//如果传入查询条件，则不对数据进行转tree处理
	if filter.SearchKey == "" {
		tree := getTreeIterative(dataList, "")
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
	} else {
		return dataList, total, err
	}
}

// GetProductTypeBatch get ProductType list  data
func (s ProductTypeService) GetTypeAndProductList(filter *entitys.QueryProductTypeForm) (rets []*entitys.TPmProductTypeVo, total int, err error) {
	ret, err := rpc.ClientProductTypeService.ListTPmProductType(context.Background(), &protosService.TPmProductTypeFilterPage{
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

	//获取已发布产品列表
	productRet, err := rpc.ClientProductService.ListTPmProduct(context.Background(), &protosService.TPmProductFilterPage{
		Page:     1,
		Limit:    1000,
		QueryObj: &protosService.TPmProductFilter{Status: 1},
	})
	if err != nil {
		return nil, 0, err
	}
	//产品分类-产品归类
	var (
		pVos map[int64][]*entitys.TPmProductVo = make(map[int64][]*entitys.TPmProductVo)
		//pCaps map[int64]int64                   = make(map[int64]int64)
	)

	if len(productRet.List) > 0 {
		list := productRet.List
		for _, request := range list {
			if request.ProductTypeId != 0 {
				data := &entitys.TPmProductVo{}
				mapstructure.WeakDecode(request, data)
				if _, ok := pVos[request.ProductTypeId]; ok {
					pVos[request.ProductTypeId] = append(pVos[request.ProductTypeId], data)
				} else {
					pVos[request.ProductTypeId] = []*entitys.TPmProductVo{data}
				}
			}
		}
	}
	//productId空值处理 & 产品列表填充
	for _, data := range dataList {
		if data.ParentId == "0" {
			data.ParentId = ""
		}
		if pVos[iotutil.ToInt64(data.Id)] != nil {
			data.Products = pVos[iotutil.ToInt64(data.Id)]
		}
	}
	//获取产品分类链表树
	tree := getTreeIterative(dataList, "")
	println(tree)
	return tree, total, err
}

// GetProductType get ProductType one record
func (s ProductTypeService) GetProductType(id string) (res *entitys.TPmProductTypeVo, err error) {
	var (
		ret *protosService.TPmProductTypeResponseObject
	)
	resv := &entitys.TPmProductTypeVo{}
	if err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.PRODUCT_TYPE_ID_DATA, id), resv); err == nil {
		return resv, nil
	}
	ret, err = rpc.ClientProductTypeService.GetByIdTPmProductType(context.Background(), &protosService.TPmProductTypeFilterById{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Msg)
	}
	if err = mapstructure.WeakDecode(ret.GetData(), &res); err != nil {
		return
	}
	if res != nil && res.ParentId == "0" {
		res.ParentId = ""
	}

	for _, v := range ret.Data.ModelsItems {
		mi := entitys.ModelsItem{
			Id:         strconv.Itoa(int(v.Id)),
			Dpid:       int(v.Dpid),
			Identifier: v.Identifier,
			Name:       v.Name,
			RwFlag:     v.RwFlag,
			DataType:   v.DataType,
			Properties: v.Properties,
			Mark:       v.Mark,
			Required:   int(v.Required),
		}
		res.ModelItems = append(res.ModelItems, &mi)
	}
	if err := cached.RedisStore.Set(persist.GetRedisKey(iotconst.PRODUCT_TYPE_ID_DATA, id), res, 0); err != nil {
		return nil, err
	}
	return res, nil
}

// DelProductType delete ProductType one record
func (s ProductTypeService) DelProductType(id int64) (err error) {
	var (
		data = protosService.TPmProductTypeRequest{}
	)
	data.Id = id
	ret, err := rpc.ClientProductTypeService.DeleteTPmProductType(context.Background(), &data)
	if err != nil {
		return err
	}
	if ret != nil && ret.Code != 200 {
		return errors.New(ret.Msg)
	}
	if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.PRODUCT_TYPE_ID_DATA, iotutil.ToString(id))); err != nil {
		return err
	}
	if err := cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA); err != nil {
		return err
	}
	return nil
}

// private func
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
			//memo[v.ParentId].Count += memo[v.Id].Count
		} else {
			memo[v.ParentId] = &entitys.TPmProductTypeVo{Children: []*entitys.TPmProductTypeVo{memo[v.Id]}}
			//memo[v.ParentId].Count += memo[v.Id].Count
		}
	}
	return memo[parentId].Children
}

// GetProductTree get ProductType list  data
func (s ProductTypeService) GetProductTree(filter *entitys.QueryProductTypeForm) (rets []*entitys.TPmProductTypeVo, total int, err error) {
	var retCache []*entitys.TPmProductTypeVo
	if err := cached.RedisStore.Get(iotconst.OPEN_PRODUCT_TREE_DATA, &retCache); err == nil {
		return retCache, 0, nil
	}
	ret, err := rpc.ClientProductTypeService.ListTPmProductType(context.Background(), &protosService.TPmProductTypeFilterPage{})
	if err != nil {
		return nil, 0, err
	}
	if ret.Code != 200 {
		return nil, 0, errors.New(ret.Msg)
	}
	//获取已发布产品列表
	productRet, err := rpc.ClientProductService.ListTPmProduct(context.Background(), &protosService.TPmProductFilterPage{
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
				Name:             request.Name,
				NameEn:           request.NameEn,
				ImageURL:         request.ImageUrl,
				Model:            request.Model,
				WifiFlag:         request.WifiFlag,
				NetworkType:      request.NetworkType,
				AttributeType:    request.AttributeType,
				Status:           request.Status,
				IsVirtualTest:    request.IsVirtualTest,
				Remark:           request.Desc,
				ProductTypeName:  request.ProductTypeName,
				PowerConsumeType: request.PowerConsumeType,
			}
			mapstructure.WeakDecode(request, data)
			if _, ok := pVos[request.ProductTypeId]; ok {
				pVos[request.ProductTypeId] = append(pVos[request.ProductTypeId], data)
			} else {
				pVos[request.ProductTypeId] = []*entitys.TPmProductVo{data}
			}
		}
	}
	//获取产品分类链表树
	tree := s.convertTreeData("0", ret.List, pVos)
	if err := cached.RedisStore.Set(iotconst.OPEN_PRODUCT_TREE_DATA, &tree, 600*time.Second); err != nil {
		return tree, 0, err
	}
	return tree, total, err
}

func (s ProductTypeService) convertTreeData(pid string, areaList []*protosService.TPmProductTypeRequest, pVos map[int64][]*entitys.TPmProductVo) []*entitys.TPmProductTypeVo {
	treeList := func() []*entitys.TPmProductTypeVo {
		treeList := []*entitys.TPmProductTypeVo{}
		flatPtr := []*entitys.TPmProductTypeVo{}
		for _, src := range areaList {
			area := &entitys.TPmProductTypeVo{
				Id:          iotutil.ToString(src.Id),
				Name:        src.Name,
				NameEn:      src.NameEn,
				Identifier:  src.Identifier,
				Sort:        src.Sort,
				ParentId:    iotutil.ToString(src.ParentId),
				Desc:        src.Desc,
				ImgFullPath: src.ImgFullPath,
				ParentName:  src.ParentName,
			}
			if pVos[iotutil.ToInt64(area.Id)] != nil {
				area.Products = pVos[iotutil.ToInt64(area.Id)]
			}
			flatPtr = append(flatPtr, area)
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
					j.Children = newChildren
				}
				j.Products = nil
				treeList = append(treeList, j)
			}
		}
		return treeList
	}()
	return treeList
}
