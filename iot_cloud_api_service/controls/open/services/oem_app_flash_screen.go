package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_app_oem/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"time"

	goerrors "go-micro.dev/v4/errors"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OemAppFlashScreenService struct {
	Ctx context.Context
}

func (s OemAppFlashScreenService) SetContext(ctx context.Context) OemAppFlashScreenService {
	s.Ctx = ctx
	return s
}

func (s OemAppFlashScreenService) CreateFlashScreen(obj *entitys.OemAppFlashScreenEntitys) error {
	obj.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	obj.TenantId, _ = metadata.Get(s.Ctx, "tenantId")
	req := entitys.OemAppFlashScreen_e2pb(obj)
	req.CreatedAt = timestamppb.New(time.Now())
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err := rpc.ClientOemAppFlashScreenService.Create(s.Ctx, req)
	if err != nil {
		return err
	}
	reqBatch := &protosService.OemAppFlashScreenUserBatch{
		OemAppFlashScreenUsers: make([]*protosService.OemAppFlashScreenUser, 0),
	}
	for i := range obj.Accounts {
		reqBatch.OemAppFlashScreenUsers = append(reqBatch.OemAppFlashScreenUsers, &protosService.OemAppFlashScreenUser{
			Id:            iotutil.GetNextSeqInt64(),
			FlashScreenId: iotutil.ToInt64(obj.Id),
			Account:       obj.Accounts[i],
		})
	}
	_, err = rpc.ClientOemAppFlashScreenUserService.CreateBatch(s.Ctx, reqBatch)
	if err != nil {
		return err
	}

	if obj.PutinImgUrls != nil && len(obj.PutinImgUrls) > 0 {
		var imgs = []string{}
		for _, p := range obj.PutinImgUrls {
			imgs = append(imgs, p.ImageUrl)
		}
		commonGlobal.SetAttachmentStatus(model.TableNameTOemAppFlashScreen+"_prod", iotutil.ToString(req.Id), imgs...)
	}
	return nil
}

func (s OemAppFlashScreenService) UpdateFlashScreen(obj *entitys.OemAppFlashScreenEntitys) error {
	req := entitys.OemAppFlashScreen_e2pb(obj)
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err := rpc.ClientOemAppFlashScreenService.Update(s.Ctx, req)
	if err != nil {
		return err
	}
	_, err = rpc.ClientOemAppFlashScreenUserService.Delete(s.Ctx, &protosService.OemAppFlashScreenUser{
		FlashScreenId: req.Id,
	})
	if err != nil {
		return err
	}
	reqBatch := &protosService.OemAppFlashScreenUserBatch{
		OemAppFlashScreenUsers: make([]*protosService.OemAppFlashScreenUser, 0),
	}
	for i := range obj.Accounts {
		reqBatch.OemAppFlashScreenUsers = append(reqBatch.OemAppFlashScreenUsers, &protosService.OemAppFlashScreenUser{
			Id:            iotutil.GetNextSeqInt64(),
			FlashScreenId: iotutil.ToInt64(obj.Id),
			Account:       obj.Accounts[i],
		})
	}
	_, err = rpc.ClientOemAppFlashScreenUserService.CreateBatch(s.Ctx, reqBatch)
	if err != nil {
		return err
	}
	if obj.PutinImgUrls != nil && len(obj.PutinImgUrls) > 0 {
		var imgs = []string{}
		for _, p := range obj.PutinImgUrls {
			imgs = append(imgs, p.ImageUrl)
		}
		commonGlobal.SetAttachmentStatus(model.TableNameTOemAppFlashScreen+"_prod", iotutil.ToString(req.Id), imgs...)
	}
	return nil
}

func (s OemAppFlashScreenService) SetFlashScreen(obj *entitys.OemAppFlashScreenEntitys) error {
	// 状态从草稿->未开始时, 需要校验下目前是否有未开始或者投放中的闪屏记录
	resp, err := rpc.ClientOemAppFlashScreenService.FindById(s.Ctx, &protosService.OemAppFlashScreenFilter{
		Id: iotutil.ToInt64(obj.Id),
	})
	if err != nil {
		return err
	}
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	if obj.Status == 2 {
		respF, err := rpc.ClientOemAppFlashScreenService.GetFlashScreen(s.Ctx, &protosService.OemAppFlashScreenRequest{
			AppKey:   resp.Data[0].AppKey,
			TenantId: tenantId,
			Version:  resp.Data[0].AppVersion,
		})
		if err != nil {
			return err
		}
		if len(respF.Data) != 0 {
			return goerrors.New("", "已存在未开始或者投放中的闪屏记录", ioterrs.ErrFlashScreenAlreadyExist)
		}
		if resp.Data[0].EndTime.AsTime().Before(time.Now()) {
			return goerrors.New("", "闪屏有效期已过，不允许投放", ioterrs.ErrFlashScreenAlreadyExpired)
		}
	}
	req := &protosService.OemAppFlashScreenUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OemAppFlashScreen{
			Id:     iotutil.ToInt64(obj.Id),
			Status: obj.Status,
		},
	}
	_, err = rpc.ClientOemAppFlashScreenService.UpdateFields(s.Ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s OemAppFlashScreenService) GetFlashScreen(id string) (*entitys.OemAppFlashScreenEntitys, error) {
	req := &protosService.OemAppFlashScreenFilter{
		Id: iotutil.ToInt64(id),
	}
	resp, err := rpc.ClientOemAppFlashScreenService.FindById(s.Ctx, req)
	if err != nil {
		return nil, err
	}
	data := entitys.OemAppFlashScreen_pb2e(resp.Data[0])
	if data.Status == 2 {
		now := time.Now()
		if resp.Data[0].EndTime.AsTime().Before(now) {
			data.Status = 5
		} else if resp.Data[0].StartTime.AsTime().Before(now) && resp.Data[0].EndTime.AsTime().After(now) {
			data.Status = 3
		}
	}
	data.Accounts = []string{}
	// 如果是指定用户，则需要查询出指定得用户列表
	if data.PutinUser == 2 {
		respUser, err := rpc.ClientOemAppFlashScreenUserService.Lists(s.Ctx, &protosService.OemAppFlashScreenUserListRequest{
			Query: &protosService.OemAppFlashScreenUser{
				FlashScreenId: req.Id,
			},
		})
		if err != nil {
			return nil, err
		}
		for i := range respUser.Data {
			data.Accounts = append(data.Accounts, respUser.Data[i].Account)
		}
	}
	return data, nil
}

func (s OemAppFlashScreenService) GetFlashScreenList(obj *entitys.OemAppFlashScreenQuery) ([]*entitys.OemAppFlashScreenEntitys, int64, error) {
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")

	//将状态转换为时间

	req := &protosService.OemAppFlashScreenListRequest{
		Page:     int64(obj.Page),
		PageSize: int64(obj.Limit),
		Query: &protosService.OemAppFlashScreen{
			Name:     obj.Query.Name,
			TenantId: tenantId,
			AppKey:   obj.Query.AppKey,
			Status:   obj.Query.Status,
		},
		OrderKey:  "updated_at",
		OrderDesc: "desc",
	}
	resp, err := rpc.ClientOemAppFlashScreenService.Lists(s.Ctx, req)
	if err != nil {
		return nil, 0, err
	}
	items := []*entitys.OemAppFlashScreenEntitys{}
	for i := range resp.Data {
		item := entitys.OemAppFlashScreen_pb2e(resp.Data[i])
		if item.Status == 2 {
			now := time.Now()
			if resp.Data[i].EndTime.AsTime().Before(now) {
				item.Status = 5
			} else if resp.Data[i].StartTime.AsTime().Before(now) && resp.Data[i].EndTime.AsTime().After(now) {
				item.Status = 3
			}
		}
		item.Accounts = []string{}
		// 如果是指定用户，则需要查询出指定得用户列表
		if item.PutinUser == 2 {
			respUser, err := rpc.ClientOemAppFlashScreenUserService.Lists(s.Ctx, &protosService.OemAppFlashScreenUserListRequest{
				Query: &protosService.OemAppFlashScreenUser{
					FlashScreenId: resp.Data[i].Id,
				},
			})
			if err != nil {
				return nil, 0, err
			}
			for i := range respUser.Data {
				item.Accounts = append(item.Accounts, respUser.Data[i].Account)
			}
		}
		items = append(items, item)
	}
	return items, resp.Total, nil
}
