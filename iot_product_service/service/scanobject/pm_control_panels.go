package scanobject

import (
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScanPmControlPanels struct {
	model.TPmControlPanels
	ProductTypeName     string
	ProductTypeFullName string
	ProductName         string
	ProductStatus       int32
}

func (o ScanPmControlPanels) ToPb() *proto.PmControlPanelsDetails {
	pbObj := proto.PmControlPanelsDetails{
		Id:                  o.Id,
		Name:                o.Name,
		NameEn:              o.NameEn,
		Lang:                o.Lang,
		Desc:                o.Desc,
		Url:                 o.Url,
		UrlName:             o.UrlName,
		PanelSize:           o.PanelSize,
		PreviewName:         o.PreviewName,
		PreviewUrl:          o.PreviewUrl,
		PreviewSize:         o.PreviewSize,
		ProductTypeId:       o.ProductTypeId,
		ProductTypeName:     o.ProductTypeName,
		ProductTypeFullName: o.ProductTypeFullName,
		ProductId:           o.ProductId,
		ProductName:         o.ProductName,
		ProductStatus:       iotutil.ToString(o.ProductStatus),
		Status:              o.Status,
		CreatedBy:           o.CreatedBy,
		UpdatedBy:           o.UpdatedBy,
		CreatedAt:           timestamppb.New(o.CreatedAt),
		UpdatedAt:           timestamppb.New(o.UpdatedAt),
		LangFileName:        o.LangFileName,
		PanelKey:            o.PanelKey,
	}
	return &pbObj
}
