// Code generated by sgen.exe,2022-05-05 19:33:56. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotstruct"
	"context"

	"cloud_platform/iot_product_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type PmControlPanelsHandler struct{}

// 创建
func (h *PmControlPanelsHandler) Create(ctx context.Context, req *proto.PmControlPanels, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	ret, err := s.CreatePmControlPanels(req)
	SetResponse(resp, err)
	if ret != nil && err == nil {
		resp.Data = ret.Id
		service.GetJsPublisherMgr().PushData(&service.NatsPubData{
			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_T_PM_CONTROL_PANEL, ret.Id, "name", req.Name, req.NameEn),
		})
	}
	return nil
}

// 匹配多条件删除
func (h *PmControlPanelsHandler) Delete(ctx context.Context, req *proto.PmControlPanels, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	_, err := s.DeletePmControlPanels(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *PmControlPanelsHandler) DeleteById(ctx context.Context, req *proto.PmControlPanels, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	_, err := s.DeleteByIdPmControlPanels(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *PmControlPanelsHandler) DeleteByIds(ctx context.Context, req *proto.PmControlPanelsBatchDeleteRequest, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	_, err := s.DeleteByIdsPmControlPanels(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *PmControlPanelsHandler) Update(ctx context.Context, req *proto.PmControlPanels, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	ret, err := s.UpdatePmControlPanels(req)
	SetResponse(resp, err)
	if ret != nil && err == nil {
		service.GetJsPublisherMgr().PushData(&service.NatsPubData{
			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_T_PM_CONTROL_PANEL, req.Id, "name", req.Name, req.NameEn),
		})
	}
	return nil
}

// 根据主键更新所有字段
func (h *PmControlPanelsHandler) UpdateAll(ctx context.Context, req *proto.PmControlPanels, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	_, err := s.UpdateAllPmControlPanels(req)
	SetResponse(resp, err)
	if err == nil {
		service.GetJsPublisherMgr().PushData(&service.NatsPubData{
			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_T_PM_CONTROL_PANEL, req.Id, "name", req.Name, req.NameEn),
		})
	}
	return nil
}

// 根据主键更新指定列
func (h *PmControlPanelsHandler) UpdateFields(ctx context.Context, req *proto.PmControlPanelsUpdateFieldsRequest, resp *proto.Response) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	_, err := s.UpdateFieldsPmControlPanels(req)
	SetResponse(resp, err)
	if err == nil {
		if req.Data.Name != "" || req.Data.NameEn != "" {
			service.GetJsPublisherMgr().PushData(&service.NatsPubData{
				Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
				Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_T_PM_CONTROL_PANEL, req.Data.Id, "name", req.Data.Name, req.Data.NameEn),
			})
		}
	}
	return nil
}

// 多条件查找，返回单条数据
func (h *PmControlPanelsHandler) Find(ctx context.Context, req *proto.PmControlPanelsFilter, resp *proto.PmControlPanelsResponse) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	data, err := s.FindPmControlPanels(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *PmControlPanelsHandler) FindById(ctx context.Context, req *proto.PmControlPanelsFilter, resp *proto.PmControlPanelsResponse) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	data, err := s.FindByIdPmControlPanels(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *PmControlPanelsHandler) Lists(ctx context.Context, req *proto.PmControlPanelsListRequest, resp *proto.PmControlPanelsResponse) error {
	s := service.PmControlPanelsSvc{Ctx: ctx}
	data, total, err := s.GetListPmControlPanels(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *PmControlPanelsHandler) SetResponse(resp *proto.PmControlPanelsResponse, data *proto.PmControlPanelsDetails, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
		resp.Total = 0
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		if data != nil {
			resp.Total = 1
			resp.Data = append(resp.Data, data)
		}
	}
}

func (h *PmControlPanelsHandler) SetPageResponse(resp *proto.PmControlPanelsResponse, list []*proto.PmControlPanelsDetails, total int64, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Total = total
		resp.Data = list
	}
}
