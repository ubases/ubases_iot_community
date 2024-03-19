package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OemAppEntryService struct {
	Ctx context.Context
}

func (s OemAppEntryService) SetContext(ctx context.Context) OemAppEntryService {
	s.Ctx = ctx
	return s
}

func (s OemAppEntryService) EntryList(req entitys.OemAppEntryListReq) ([]*entitys.OemAppEntryListRes, int64, error) {

	res, err := rpc.ClientOemAppEntryService.ListDiy(s.Ctx, &protosService.OemAppEntryListDiyRequqest{
		DirId: iotutil.ToInt64(req.DirId),
		// DocId: iotutil.ToInt64(req.DocId) ,
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
	var rs = make([]*entitys.OemAppEntryListRes, 0)
	if res.Data == nil || len(res.Data) <= 0 {
		return rs, 0, nil
	}
	for _, v := range res.Data {
		rs = append(rs, &entitys.OemAppEntryListRes{
			SetingId:  iotutil.ToString(v.SetingId),
			DocId:     iotutil.ToString(v.DocId),
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

// 创建词条设置.
func (s OemAppEntryService) EntrySetingCreate(req entitys.OemAppEntrySetingSaveReq) error {
	var setingIdAdd int64
	if req.SetingId != "" && req.SetingId != "0" {
		setingIdAdd = iotutil.ToInt64(req.SetingId)
	} else {
		setingIdAdd = iotutil.GetNextSeqInt64()
	}
	res, err := rpc.ClientOemAppEntrySetingService.Create(s.Ctx, &protosService.OemAppEntrySeting{
		Id:       setingIdAdd,
		DirId:    iotutil.ToInt64(req.DirId),
		Sort:     int32(req.Sort),
		IsEnable: int32(req.IsEnable),
		IsNormal: int32(req.IsNormal),
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// 词条保存
func (s OemAppEntryService) EntrySave(req entitys.OemAppEntrySaveReq) (string, error) {
	//修改时候的设置id
	setingIdReq := iotutil.ToInt64(req.SetingId)
	//新增时候的设置id
	var setingIdAdd int64
	//新增词条
	if setingIdReq == 0 {

		setingIdAdd = iotutil.GetNextSeqInt64()
		errSetCreate := s.EntrySetingCreate(entitys.OemAppEntrySetingSaveReq{
			SetingId: iotutil.ToString(setingIdAdd),
			DirId:    req.DirId,
			Sort:     0,
			IsEnable: 2,
			IsNormal: 2,
		})
		if errSetCreate != nil {
			return "", errSetCreate
		}
		//原逻辑暂时注释
		// setingIdAdd = iotutil.GetNextSeqInt64()
		// res,err := rpc.ClientOemAppEntrySetingService.Create(s.Ctx,&protosService.OemAppEntrySeting{
		// 	Id: setingIdAdd,
		// 	DirId: iotutil.ToInt64(req.DirId),
		// 	Sort: 0,
		// 	IsEnable: 2,
		// 	IsNormal: 2,
		// })
		// if err != nil {
		// 	return "", err
		// }
		// if res.Code != 200 {
		// 	return "", errors.New(res.Message)
		// }
	}

	//新增
	if setingIdReq == 0 {
		_, errEntry := s.EntryCreate(&protosService.OemAppEntry{
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
			rpc.ClientOemAppEntryService.UpdateFields(s.Ctx, &protosService.OemAppEntryUpdateFieldsRequest{
				Fields: []string{"title", "content", "updated_at"},
				Data: &protosService.OemAppEntry{
					Id:        set.Id,
					Title:     req.Title,
					Content:   req.Content,
					UpdatedAt: timestamppb.New(time.Now()),
				},
			})
		} else {
			//新增
			_, errEntry := s.EntryCreate(&protosService.OemAppEntry{
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

func (s OemAppEntryService) EntryCreate(req *protosService.OemAppEntry) (int64, error) {
	var entryId int64
	if req.Id != 0 {
		entryId = req.Id
	} else {
		entryId = iotutil.GetNextSeqInt64()
	}
	resEntry, errEntry := rpc.ClientOemAppEntryService.Create(s.Ctx, &protosService.OemAppEntry{
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
func (s OemAppEntryService) GetEntryByLangSetingId(lang string, setingId int64) (*protosService.OemAppEntry, error) {

	if setingId <= 0 {
		return nil, errors.New("参数错误")
	}

	res, err := rpc.ClientOemAppEntryService.Find(s.Ctx, &protosService.OemAppEntryFilter{
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

func (s OemAppEntryService) EntryDetail(setingId string, lang string) (*entitys.OemAppEntrySaveReq, error) {
	res, err := s.GetEntryByLangSetingId(lang, iotutil.ToInt64(setingId))
	if err != nil {
		return nil, err
	}
	var rs = entitys.OemAppEntrySaveReq{}
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
func (s OemAppEntryService) EntrySetingSave(req entitys.OemAppEntrySetingSaveReq) (string, error) {
	res, err := rpc.ClientOemAppEntrySetingService.UpdateFields(s.Ctx, &protosService.OemAppEntrySetingUpdateFieldsRequest{
		Fields: []string{"is_normal", "is_enable", "sort", "dir_id"},
		Data: &protosService.OemAppEntrySeting{
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
func (s OemAppEntryService) EntrySetingDetail(setingId string) (*entitys.OemAppEntrySetingSaveReq, error) {
	res, err := rpc.ClientOemAppEntrySetingService.FindById(s.Ctx, &protosService.OemAppEntrySetingFilter{
		Id: iotutil.ToInt64(setingId),
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	var rs = entitys.OemAppEntrySetingSaveReq{}
	if res.Data != nil && len(res.Data) > 0 {
		rs.DirId = iotutil.ToString(res.Data[0].DirId)
		rs.IsEnable = int64(res.Data[0].IsEnable)
		rs.IsNormal = int64(res.Data[0].IsNormal)
		rs.SetingId = setingId
		rs.Sort = int64(res.Data[0].Sort)
	}

	return &rs, nil
}

func (s OemAppEntryService) EntryDelete(setingId string) (string, error) {

	res, err := rpc.ClientOemAppEntryService.Delete(s.Ctx, &protosService.OemAppEntry{
		SetingId: iotutil.ToInt64(setingId),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	resSet, errSet := rpc.ClientOemAppEntrySetingService.DeleteById(s.Ctx, &protosService.OemAppEntrySeting{
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
