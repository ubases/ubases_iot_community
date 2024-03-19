package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
	"time"

	goerrors "go-micro.dev/v4/errors"
)

type ProductMaterialService struct {
	Ctx context.Context
}

func (s ProductMaterialService) SetContext(ctx context.Context) ProductMaterialService {
	s.Ctx = ctx
	return s
}

func (s ProductMaterialService) ClickProductMaterial(id int64) error {
	req := &protosService.OpmProductMaterials{
		Id: id,
	}
	_, err := rpc.ClientOpmProductMaterialsService.Click(s.Ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s ProductMaterialService) GetProductMaterial(uids []string, tenantId, productKey, lang string) ([]*entitys.OpmProductMaterialsEntitys, error) {
	// tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	items := []*entitys.OpmProductMaterialsEntitys{}
	checked := map[string]struct{}{}
	for _, uid := range uids {
		data := &entitys.OpmProductMaterialsEntitys{}
		// 通过解析耗材uid，来查询对应品牌和香型的耗材
		if len(uid) != 18 {
			return nil, goerrors.New("", "material uid length is 18", ioterrs.ErrAppRequestParam)
		}
		if _, ok := checked[uid[0:4]]; ok {
			continue
		}
		if err := cached.RedisStore.Get(fmt.Sprintf(iotconst.PRODUCT_MATERIAL_DATA, tenantId, uid[0:2], uid[2:4], lang), data); err == nil && data != nil {
			checked[uid[0:4]] = struct{}{}
			if len(productKey) != 0 {
				// 先通过productKey, 找到productId,然后再匹配
				respPro, err := rpc.ProductService.Find(s.Ctx, &protosService.OpmProductFilter{
					ProductKey: productKey,
				})
				if err != nil {
					return nil, err
				}
				var checked bool
				for i := range data.SuitableProduct {
					if iotutil.ToInt64(data.SuitableProduct[i]) == respPro.Data[0].Id {
						checked = true
					}
				}
				if checked {
					items = append(items, data)
				}
			} else {
				items = append(items, data)
			}
			continue
		}
		req := &protosService.OpmProductMaterialsFilter{
			TenantId:      tenantId,
			BrandCode:     uid[0:2],
			FragranceCode: uid[2:4],
			Lang:          lang,
		}
		resp, err := rpc.ClientOpmProductMaterialsService.Find(s.Ctx, req)
		if err != nil && goerrors.FromError(err).GetDetail() == ioterrs.ErrRecordNotFound {
			continue
		} else if err != nil {
			return nil, err
		}
		iotlogger.LogHelper.Helper.Debug(resp.Data[0])
		if resp.Data[0].Lang != lang {
			continue
		}
		data = entitys.OpmProductMaterials_pb2e(resp.Data[0])
		data.SuitableProduct = []string{}
		respRel, err := rpc.ClientOpmProductMaterialRelationService.Lists(s.Ctx, &protosService.OpmProductMaterialRelationListRequest{
			Query: &protosService.OpmProductMaterialRelation{
				MaterialId: resp.Data[0].Id,
			},
		})
		if err != nil {
			return nil, err
		}
		for i := range respRel.Data {
			data.SuitableProduct = append(data.SuitableProduct, iotutil.ToString(respRel.Data[i].ProductId))
		}
		if err := cached.RedisStore.Set(fmt.Sprintf(iotconst.PRODUCT_MATERIAL_DATA, tenantId, uid[0:2], uid[2:4], lang), *data, 10*time.Minute); err != nil {
			return nil, err
		}
		checked[uid[0:4]] = struct{}{}
		if len(productKey) != 0 {
			// 先通过productKey, 找到productId,然后再匹配
			respPro, err := rpc.ProductService.Find(s.Ctx, &protosService.OpmProductFilter{
				ProductKey: productKey,
			})
			if err != nil {
				return nil, err
			}
			var checked bool
			for i := range data.SuitableProduct {
				if iotutil.ToInt64(data.SuitableProduct[i]) == respPro.Data[0].Id {
					checked = true
				}
			}
			if checked {
				items = append(items, data)
			}
		} else {
			items = append(items, data)
		}
	}
	return items, nil
}
