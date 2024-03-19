package services

import (
	//langEntitys "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type SysAppHelpCenterService struct {
	Ctx context.Context
}

func (s SysAppHelpCenterService) SetContext(ctx context.Context) SysAppHelpCenterService {
	s.Ctx = ctx
	return s
}

func (s SysAppHelpCenterService) CreateHelpCenter(req *entitys.SysAppHelpCenterEntitys) error {
	req.Id = iotutil.GetNextSeqInt64()
	reqHelp := entitys.SysAppHelpCenter_e2pb(req)
	reqHelp.CreatedAt = timestamppb.Now()
	reqHelp.UpdatedAt = timestamppb.Now()
	_, err := rpc.ClientSysAppHelpCenterService.Create(s.Ctx, reqHelp)
	if err != nil {
		return err
	}
	return nil
}

func (s SysAppHelpCenterService) CopyHelpCenter(req *entitys.SysAppHelpCenterEntitys) error {
	// 先通过helpId查询相应的目录列表, 然后级联复制目录及词条
	respDir, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
		Query: &protosService.SysAppDocDir{
			HelpId:   req.Id,
			ParentId: -1,
		},
	})
	if err != nil {
		return err
	}
	newHelpId := iotutil.GetNextSeqInt64()
	if len(respDir.Data) != 0 {
		if err := CopyAppEntryByRecurse(respDir.Data, newHelpId, 0); err != nil {
			return err
		}
	}
	reqHelp := entitys.SysAppHelpCenter_e2pb(req)
	reqHelp.Id = newHelpId
	reqHelp.CreatedAt = timestamppb.Now()
	reqHelp.UpdatedAt = timestamppb.Now()
	_, err = rpc.ClientSysAppHelpCenterService.Create(s.Ctx, reqHelp)
	if err != nil {
		return err
	}
	return nil
}

func (s SysAppHelpCenterService) UpdateHelpCenter(req *entitys.SysAppHelpCenterEntitys) error {
	reqHelp := entitys.SysAppHelpCenter_e2pb(req)
	reqHelp.UpdatedAt = timestamppb.Now()
	_, err := rpc.ClientSysAppHelpCenterService.Update(s.Ctx, reqHelp)
	if err != nil {
		return err
	}
	return nil
}

func (s SysAppHelpCenterService) DeleteHelpCenter(helpId int64) error {
	// 先通过helpId查询相应的目录列表, 然后级联删除目录及词条
	respDir, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
		Query: &protosService.SysAppDocDir{
			HelpId:   helpId,
			ParentId: -1,
		},
	})
	if err != nil {
		return err
	}
	if len(respDir.Data) != 0 {
		if err := DelAppEntryByRecurse(respDir.Data); err != nil {
			return err
		}
	}
	_, err = rpc.ClientSysAppHelpCenterService.DeleteById(s.Ctx, &protosService.SysAppHelpCenter{
		Id: helpId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s SysAppHelpCenterService) GetHelpCenter(helpId int64) (*entitys.SysAppHelpCenterEntitys, error) {
	resp, err := rpc.ClientSysAppHelpCenterService.FindById(s.Ctx, &protosService.SysAppHelpCenterFilter{
		Id: helpId,
	})
	if err != nil {
		return nil, err
	}
	data := entitys.SysAppHelpCenter_pb2e(resp.Data[0])
	// 查询帮助中心下的词条数
	respDir, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
		Query: &protosService.SysAppDocDir{
			HelpId:   resp.Data[0].Id,
			ParentId: -1,
		},
	})
	if err != nil {
		return nil, err
	}
	var count int
	if len(respDir.Data) != 0 {
		c, err := GetAppEntryCount(respDir.Data)
		if err != nil {
			return nil, err
		}
		count += c
	}
	data.Count = count
	return data, nil
}

func (s SysAppHelpCenterService) GetHelpCenterList(req *entitys.SysAppHelpCenterQuery) ([]*entitys.SysAppHelpCenterEntitys, error) {
	resp, err := rpc.ClientSysAppHelpCenterService.Lists(s.Ctx, &protosService.SysAppHelpCenterListRequest{
		Page:     int64(req.Page),
		PageSize: int64(req.Limit),
		Query: &protosService.SysAppHelpCenter{
			Name:         req.Query.Name,
			TemplateType: req.Query.TemplateType,
			Version:      req.Query.Version,
		},
	})
	if err != nil {
		return nil, err
	}
	data := []*entitys.SysAppHelpCenterEntitys{}
	for i := range resp.Data {
		item := entitys.SysAppHelpCenter_pb2e(resp.Data[i])
		// 查询每个帮助中心下的词条数
		respDir, err := rpc.ClientSysAppDocDirService.Lists(s.Ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				HelpId:   resp.Data[i].Id,
				ParentId: -1,
			},
		})
		if err != nil {
			return nil, err
		}
		var count int
		if len(respDir.Data) != 0 {
			c, err := GetAppEntryCount(respDir.Data)
			if err != nil {
				return nil, err
			}
			count += c
		}
		item.Count = count
		data = append(data, item)
	}
	return data, nil
}

func (s SysAppHelpCenterService) GetHelpCenterListForOpen(req *entitys.SysAppHelpCenterQuery) ([]*entitys.SysAppHelpCenterEntitys, error) {
	respDict, err := rpc.TConfigDictDataServerService.Lists(s.Ctx, &protosService.ConfigDictDataListRequest{
		Query: &protosService.ConfigDictData{
			DictType: "app_template_type",
		},
	})
	if err != nil {
		return nil, err
	}
	data := []*entitys.SysAppHelpCenterEntitys{}
	for i := range respDict.Data {
		resp, err := rpc.ClientSysAppHelpCenterService.Lists(s.Ctx, &protosService.SysAppHelpCenterListRequest{
			Page:      1,
			PageSize:  1,
			OrderKey:  "version",
			OrderDesc: "desc",
			Query: &protosService.SysAppHelpCenter{
				TemplateType: iotutil.ToInt32(respDict.Data[i].DictValue),
			},
		})
		if err != nil {
			return nil, err
		}
		for i := range resp.Data {
			item := entitys.SysAppHelpCenter_pb2e(resp.Data[i])
			data = append(data, item)
		}
	}
	return data, nil
}

// 递归获取app帮助中心下文档数
func GetAppEntryCount(data []*protosService.SysAppDocDir) (int, error) {
	var count int
	ctx := context.Background()
	for i := range data {
		respEntrySeting, err := rpc.ClientSysAppEntrySetingService.Lists(ctx, &protosService.SysAppEntrySetingListRequest{
			Query: &protosService.SysAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return 0, err
		}
		count += len(respEntrySeting.Data)
		respDir, err := rpc.ClientSysAppDocDirService.Lists(ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return 0, err
		}
		if len(respDir.Data) != 0 {
			c, err := GetAppEntryCount(respDir.Data)
			if err != nil {
				return 0, err
			}
			count += c
		}
	}
	return count, nil
}

// 递归删除app帮助中心下目录和文档
func DelAppEntryByRecurse(data []*protosService.SysAppDocDir) error {
	ctx := context.Background()
	for i := range data {
		// 先查询app目录设置id关联表
		respEntrySeting, err := rpc.ClientSysAppEntrySetingService.Lists(ctx, &protosService.SysAppEntrySetingListRequest{
			Query: &protosService.SysAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		// 遍历关联表列表，通过settingId来删除文档
		for j := range respEntrySeting.Data {
			// 先删除关联表的目录和settingId之间的关联关系
			_, err = rpc.ClientSysAppEntrySetingService.Delete(ctx, &protosService.SysAppEntrySeting{
				Id: respEntrySeting.Data[j].Id,
			})
			if err != nil {
				return err
			}
			// 然后通过settingId来删除文档
			_, err = rpc.ClientSysAppEntryService.Delete(ctx, &protosService.SysAppEntry{
				SetingId: respEntrySeting.Data[j].Id,
			})
			if err != nil {
				return err
			}
		}
		// 考虑存在多层子目录，需级联查询并删除
		respDir, err := rpc.ClientSysAppDocDirService.Lists(ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		if len(respDir.Data) != 0 {
			err = DelAppEntryByRecurse(respDir.Data)
			if err != nil {
				return err
			}
		}
		// 删除目录
		_, err = rpc.ClientSysAppDocDirService.Delete(ctx, &protosService.SysAppDocDir{
			Id: data[i].Id,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 递归复制app帮助中心下目录和文档
func CopyAppEntryByRecurse(data []*protosService.SysAppDocDir, newHelpId, parentId int64) error {
	ctx := context.Background()
	for i := range data {
		// 先查询app目录设置id关联表
		respEntrySeting, err := rpc.ClientSysAppEntrySetingService.Lists(ctx, &protosService.SysAppEntrySetingListRequest{
			Query: &protosService.SysAppEntrySeting{
				DirId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		// 遍历关联表列表，通过settingId来复制文档
		dirId := iotutil.GetNextSeqInt64()
		// entrySets := []*protosService.SysAppEntrySeting{}
		for j := range respEntrySeting.Data {
			// 先通过setintId来获取多语言文档列表
			respEntry, err := rpc.ClientSysAppEntryService.Lists(ctx, &protosService.SysAppEntryListRequest{
				Query: &protosService.SysAppEntry{
					SetingId: respEntrySeting.Data[j].Id,
				},
			})
			if err != nil {
				return err
			}
			// 重置文档id和setingId, 并新建文档记录
			setingId := iotutil.GetNextSeqInt64()
			for k := range respEntry.Data {
				respEntry.Data[k].Id = iotutil.GetNextSeqInt64()
				respEntry.Data[k].SetingId = setingId
			}
			// 批量创建文档
			if len(respEntry.Data) != 0 {
				_, err = rpc.ClientSysAppEntryService.CreateBatch(ctx, &protosService.SysAppEntryBatchRequest{
					SysAppEntrys: respEntry.Data,
				})
				if err != nil {
					return err
				}
			}
			respEntrySeting.Data[j].Id = setingId
			respEntrySeting.Data[j].DirId = dirId
		}
		// 批量创建app文档setingId关联记录
		if len(respEntrySeting.Data) != 0 {
			_, err = rpc.ClientSysAppEntrySetingService.CreateBatch(ctx, &protosService.SysAppEntrySetingBatchRequest{
				SysAppEntrySetings: respEntrySeting.Data,
			})
			if err != nil {
				return err
			}
		}

		// 考虑存在多层子目录，需级联查询并复制
		respDir, err := rpc.ClientSysAppDocDirService.Lists(ctx, &protosService.SysAppDocDirListRequest{
			Query: &protosService.SysAppDocDir{
				ParentId: data[i].Id,
			},
		})
		if err != nil {
			return err
		}
		if len(respDir.Data) != 0 {
			err = CopyAppEntryByRecurse(respDir.Data, newHelpId, dirId)
			if err != nil {
				return err
			}
		}

		data[i].Id = dirId
		data[i].ParentId = parentId
		data[i].HelpId = newHelpId
	}

	// 批量创建目录
	if len(data) != 0 {
		_, err := rpc.ClientSysAppDocDirService.CreateBatch(ctx, &protosService.SysAppDocDirBatchRequest{
			SysAppDocDirs: data,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
