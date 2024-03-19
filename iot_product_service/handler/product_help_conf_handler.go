// Code generated by sgen.exe,2022-08-18 20:09:05. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	goerrors "go-micro.dev/v4/errors"

	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_product_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type ProductHelpConfHandler struct{}

// 创建
func (h *ProductHelpConfHandler) Create(ctx context.Context, req *proto.ProductHelpConf, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.CreateProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfCreate)
	}
	return nil
}

// 匹配多条件删除
func (h *ProductHelpConfHandler) Delete(ctx context.Context, req *proto.ProductHelpConf, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.DeleteProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfDelete)
	}
	return nil
}

// 匹配ID删除
func (h *ProductHelpConfHandler) DeleteById(ctx context.Context, req *proto.ProductHelpConf, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.DeleteByIdProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfDelete)
	}
	return nil
}

// 根据主键批量删除
func (h *ProductHelpConfHandler) DeleteByIds(ctx context.Context, req *proto.ProductHelpConfBatchDeleteRequest, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.DeleteByIdsProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfDelete)
	}
	return nil
}

// 更新
func (h *ProductHelpConfHandler) Update(ctx context.Context, req *proto.ProductHelpConf, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.UpdateProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfUpdate)
	}
	return nil
}

// 根据主键更新所有字段
func (h *ProductHelpConfHandler) UpdateAll(ctx context.Context, req *proto.ProductHelpConf, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.UpdateAllProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfUpdate)
	}
	return nil
}

// 根据主键更新指定列
func (h *ProductHelpConfHandler) UpdateFields(ctx context.Context, req *proto.ProductHelpConfUpdateFieldsRequest, resp *proto.Response) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	_, err := s.UpdateFieldsProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfUpdate)
	}
	return nil
}

// 多条件查找，返回单条数据
func (h *ProductHelpConfHandler) Find(ctx context.Context, req *proto.ProductHelpConfFilter, resp *proto.ProductHelpConfResponse) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	data, err := s.FindProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfGet)
	}
	h.SetResponse(resp, data)
	return nil
}

// 根据ID查找，返回单条数据
func (h *ProductHelpConfHandler) FindById(ctx context.Context, req *proto.ProductHelpConfFilter, resp *proto.ProductHelpConfResponse) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	data, err := s.FindByIdProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfGet)
	}
	h.SetResponse(resp, data)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *ProductHelpConfHandler) Lists(ctx context.Context, req *proto.ProductHelpConfListRequest, resp *proto.ProductHelpConfResponse) error {
	s := service.ProductHelpConfSvc{Ctx: ctx}
	data, total, err := s.GetListProductHelpConf(req)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBProductHelpConfList)
	}
	h.SetPageResponse(resp, data, total)
	return nil
}

func (h *ProductHelpConfHandler) SetResponse(resp *proto.ProductHelpConfResponse, data *proto.ProductHelpConf) {
	resp.Code = ioterrs.Success
	resp.Message = "success"
	if data != nil {
		resp.Total = 1
		resp.Data = append(resp.Data, data)
	}
}

func (h *ProductHelpConfHandler) SetPageResponse(resp *proto.ProductHelpConfResponse, list []*proto.ProductHelpConf, total int64) {
	resp.Code = ioterrs.Success
	resp.Message = "success"
	resp.Total = total
	resp.Data = list
}
