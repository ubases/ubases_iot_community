package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type SysAppEntryService struct {
	Ctx context.Context
}

func (s SysAppEntryService) SetContext(ctx context.Context) SysAppEntryService {
	s.Ctx = ctx
	return s
}

func (s SysAppEntryService) EntryList(req entitys.SysAppEntryListReq) ([]*entitys.SysAppEntryListRes, int64, error) {

	res, err := rpc.ClientSysAppEntryService.ListDiy(s.Ctx, &protosService.SysAppEntryListDiyRequqest{
		DirId: iotutil.ToInt64(req.DirId),
		//DocId: iotutil.ToInt64(req.DocId) ,
		Lang:     req.Lang,
		Title:    req.Title,
		IsNormal: req.IsNormal,
		IsEnable: req.IsEnable,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, 0, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, 0, errors.New(res.Message)
	}
	var rs = make([]*entitys.SysAppEntryListRes, 0)
	if res.Data == nil || len(res.Data) <= 0 {
		return rs, 0, nil
	}
	for _, v := range res.Data {
		rs = append(rs, &entitys.SysAppEntryListRes{
			SetingId: iotutil.ToString(v.SetingId),
			//DocId: iotutil.ToString(v.DocId),
			DirId:     iotutil.ToString(v.DirId),
			Lang:      v.Lang,
			Title:     v.Title,
			IsNormal:  v.IsNormal,
			IsEnable:  v.IsEnable,
			DirName:   v.DirName,
			Sort:      v.Sort,
			UpdatedAt: v.UpdatedAt.AsTime().Unix(),
		})
	}
	return rs, res.Total, nil

}

// 词条保存
func (s SysAppEntryService) EntrySave(req entitys.SysAppEntrySaveReq) (string, error) {
	//修改时候的设置id
	setingIdReq := iotutil.ToInt64(req.SetingId)
	//新增时候的设置id
	var setingIdAdd int64
	//新增词条
	if setingIdReq == 0 {
		if req.DirId == "" {
			return "", errors.New("参数错误, 目录ID不可为空.")
		}
		setingIdAdd = iotutil.GetNextSeqInt64()
		res, err := rpc.ClientSysAppEntrySetingService.Create(s.Ctx, &protosService.SysAppEntrySeting{
			Id:       setingIdAdd,
			DirId:    iotutil.ToInt64(req.DirId),
			Sort:     0,
			IsEnable: 2,
			IsNormal: 2,
		})
		if err != nil {
			return "", err
		}
		if res.Code != 200 {
			return "", errors.New(res.Message)
		}
	}

	//新增
	if setingIdReq == 0 {
		_, errEntry := s.EntryCreate(&protosService.SysAppEntry{
			Lang:      req.Lang,
			Title:     req.Title,
			Content:   req.Content,
			SetingId:  setingIdAdd,
			UpdatedAt: timestamppb.New(time.Now()),
		})
		if errEntry != nil {
			return "", errEntry
		}

	} else {
		//修改
		set, errSet := s.GetEntryByLangSetingId(req.Lang, setingIdReq)
		if errSet != nil {
			return "", errSet
		}
		if set != nil {
			//修改
			rpc.ClientSysAppEntryService.UpdateFields(s.Ctx, &protosService.SysAppEntryUpdateFieldsRequest{
				Fields: []string{"title", "content", "updated_at"},
				Data: &protosService.SysAppEntry{
					Id:        set.Id,
					Title:     req.Title,
					Content:   req.Content,
					UpdatedAt: timestamppb.New(time.Now()),
				},
			})
		} else {
			//新增
			_, errEntry := s.EntryCreate(&protosService.SysAppEntry{
				Lang:      req.Lang,
				Title:     req.Title,
				Content:   req.Content,
				SetingId:  setingIdReq,
				UpdatedAt: timestamppb.New(time.Now()),
			})
			if errEntry != nil {
				return "", errEntry
			}

		}

	}

	return "success", nil
}

func (s SysAppEntryService) EntryCreate(req *protosService.SysAppEntry) (int64, error) {
	entryId := iotutil.GetNextSeqInt64()
	resEntry, errEntry := rpc.ClientSysAppEntryService.Create(s.Ctx, &protosService.SysAppEntry{
		Id:        entryId,
		Lang:      req.Lang,
		Title:     req.Title,
		Content:   req.Content,
		SetingId:  req.SetingId,
		UpdatedAt: req.UpdatedAt,
	})

	if errEntry != nil {
		return 0, errEntry
	}
	if resEntry.Code != 200 {
		return 0, errors.New(resEntry.Message)
	}
	return entryId, nil
}

// 根据语种和设置id查到唯一的词条
func (s SysAppEntryService) GetEntryByLangSetingId(lang string, setingId int64) (*protosService.SysAppEntry, error) {

	res, err := rpc.ClientSysAppEntryService.Find(s.Ctx, &protosService.SysAppEntryFilter{
		Lang:     lang,
		SetingId: setingId,
	})

	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, nil
	}
	return res.Data[0], nil
}

func (s SysAppEntryService) EntryDetail(setingId string, lang string) (*entitys.SysAppEntrySaveReq, error) {
	res, err := s.GetEntryByLangSetingId(lang, iotutil.ToInt64(setingId))
	if err != nil {
		return nil, err
	}
	var rs = entitys.SysAppEntrySaveReq{}
	if res != nil {
		rs.Content = res.Content
		rs.Lang = res.Lang
		rs.SetingId = iotutil.ToString(res.SetingId)
		rs.Title = res.Title
	} else {
		rs.SetingId = setingId
		rs.Lang = lang
	}

	return &rs, nil
}

// 词条设置保存
func (s SysAppEntryService) EntrySetingSave(req entitys.SysAppEntrySetingSaveReq) (string, error) {
	res, err := rpc.ClientSysAppEntrySetingService.UpdateFields(s.Ctx, &protosService.SysAppEntrySetingUpdateFieldsRequest{
		Fields: []string{"is_normal", "is_enable", "sort", "dir_id"},
		Data: &protosService.SysAppEntrySeting{
			Id:       iotutil.ToInt64(req.SetingId),
			DirId:    iotutil.ToInt64(req.DirId),
			Sort:     int32(req.Sort),
			IsEnable: int32(req.IsEnable),
			IsNormal: int32(req.IsNormal),
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

// 词条设置详细
func (s SysAppEntryService) EntrySetingDetail(setingId string) (*entitys.SysAppEntrySetingSaveReq, error) {
	res, err := rpc.ClientSysAppEntrySetingService.FindById(s.Ctx, &protosService.SysAppEntrySetingFilter{
		Id: iotutil.ToInt64(setingId),
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var rs = entitys.SysAppEntrySetingSaveReq{}
	if res.Data != nil && len(res.Data) > 0 {
		rs.DirId = iotutil.ToString(res.Data[0].DirId)
		rs.IsEnable = int64(res.Data[0].IsEnable)
		rs.IsNormal = int64(res.Data[0].IsNormal)
		rs.SetingId = setingId
		rs.Sort = int64(res.Data[0].Sort)
	}

	return &rs, nil
}

func (s SysAppEntryService) EntryDelete(setingId string) (string, error) {

	res, err := rpc.ClientSysAppEntryService.Delete(s.Ctx, &protosService.SysAppEntry{
		SetingId: iotutil.ToInt64(setingId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	resSet, errSet := rpc.ClientSysAppEntrySetingService.DeleteById(s.Ctx, &protosService.SysAppEntrySeting{
		Id: iotutil.ToInt64(setingId),
	})

	if errSet != nil {
		return "", errSet
	}
	if resSet.Code != 200 {
		return "", errors.New(resSet.Message)
	}
	return "success", nil
}
