// Code generated by sgen.exe,2022-04-19 09:58:34. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_basic_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type ConfigOssHandler struct{}

// 创建
func (h *ConfigOssHandler) Create(ctx context.Context, req *proto.ConfigOss, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.CreateConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *ConfigOssHandler) Delete(ctx context.Context, req *proto.ConfigOss, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.DeleteConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *ConfigOssHandler) DeleteById(ctx context.Context, req *proto.ConfigOss, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.DeleteByIdConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *ConfigOssHandler) DeleteByIds(ctx context.Context, req *proto.ConfigOssBatchDeleteRequest, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.DeleteByIdsConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *ConfigOssHandler) Update(ctx context.Context, req *proto.ConfigOss, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.UpdateConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *ConfigOssHandler) UpdateAll(ctx context.Context, req *proto.ConfigOss, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.UpdateAllConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *ConfigOssHandler) UpdateFields(ctx context.Context, req *proto.ConfigOssUpdateFieldsRequest, resp *proto.Response) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	_, err := s.UpdateFieldsConfigOss(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *ConfigOssHandler) Find(ctx context.Context, req *proto.ConfigOssFilter, resp *proto.ConfigOssResponse) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	data, err := s.FindConfigOss(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *ConfigOssHandler) FindById(ctx context.Context, req *proto.ConfigOssFilter, resp *proto.ConfigOssResponse) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	data, err := s.FindByIdConfigOss(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *ConfigOssHandler) Lists(ctx context.Context, req *proto.ConfigOssListRequest, resp *proto.ConfigOssResponse) error {
	s := service.ConfigOssSvc{Ctx: ctx}
	data, total, err := s.GetListConfigOss(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *ConfigOssHandler) SetResponse(resp *proto.ConfigOssResponse, data *proto.ConfigOss, err error) {
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

func (h *ConfigOssHandler) SetPageResponse(resp *proto.ConfigOssResponse, list []*proto.ConfigOss, total int64, err error) {
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
