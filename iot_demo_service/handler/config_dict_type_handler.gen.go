// fixme 本文件是demo，展示Handler实现；正式服务可以删除或者替换成自己的Handler
// 文件中的proto包，就是google proto3文件生成的go文件所在的包名，可手动修改为自己生成的go文件包名
package handler

import (
	"context"

	"cloud_platform/iot_demo_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type ConfigDictTypeHandler struct{}

// 创建
func (h *ConfigDictTypeHandler) Create(ctx context.Context, req *proto.ConfigDictType, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.CreateConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *ConfigDictTypeHandler) Delete(ctx context.Context, req *proto.ConfigDictType, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.DeleteConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *ConfigDictTypeHandler) DeleteById(ctx context.Context, req *proto.ConfigDictType, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.DeleteByIdConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *ConfigDictTypeHandler) DeleteByIds(ctx context.Context, req *proto.ConfigDictTypeBatchDeleteRequest, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.DeleteByIdsConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *ConfigDictTypeHandler) Update(ctx context.Context, req *proto.ConfigDictType, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.UpdateConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *ConfigDictTypeHandler) UpdateAll(ctx context.Context, req *proto.ConfigDictType, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.UpdateAllConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *ConfigDictTypeHandler) UpdateFields(ctx context.Context, req *proto.ConfigDictTypeUpdateFieldsRequest, resp *proto.Response) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	_, err := s.UpdateFieldsConfigDictType(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *ConfigDictTypeHandler) Find(ctx context.Context, req *proto.ConfigDictTypeFilter, resp *proto.ConfigDictTypeResponse) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	data, err := s.FindConfigDictType(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *ConfigDictTypeHandler) FindById(ctx context.Context, req *proto.ConfigDictTypeFilter, resp *proto.ConfigDictTypeResponse) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	data, err := s.FindByIdConfigDictType(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *ConfigDictTypeHandler) Lists(ctx context.Context, req *proto.ConfigDictTypeListRequest, resp *proto.ConfigDictTypeResponse) error {
	s := service.ConfigDictTypeSvc{Ctx: ctx}
	data, total, err := s.GetListConfigDictType(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *ConfigDictTypeHandler) SetResponse(resp *proto.ConfigDictTypeResponse, data *proto.ConfigDictType, err error) {
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

func (h *ConfigDictTypeHandler) SetPageResponse(resp *proto.ConfigDictTypeResponse, list []*proto.ConfigDictType, total int64, err error) {
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
