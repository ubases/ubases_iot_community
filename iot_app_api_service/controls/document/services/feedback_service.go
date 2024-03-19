package services

import (
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"sort"
	"time"
)

type UcFeedBackService struct {
	Ctx context.Context
}

func (s UcFeedBackService) SetContext(ctx context.Context) UcFeedBackService {
	s.Ctx = ctx
	return s
}

// 用户反馈详细
func (s UcFeedBackService) GetUcFeedBackDetail(lang, appKey, tenantId, id string) (entitys.UcUserFeedbackEntitys, error) {
	if id == "" {
		return entitys.UcUserFeedbackEntitys{}, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.UcFeedbackService.FeedbackLists(s.Ctx, &protosService.UcUserFeedbackListRequest{Query: &protosService.UcUserFeedback{
		Id: rid,
	}})
	if err != nil {
		return entitys.UcUserFeedbackEntitys{}, err
	}
	if req.Code != 200 {
		return entitys.UcUserFeedbackEntitys{}, errors.New(req.Message)
	}
	feedTypes, err := s.GetFeedBackQuestionType(lang, appKey, tenantId, req.Data[0].QuestionTypeId)
	feedTypesMap := make(map[int64]string)
	for _, v := range feedTypes {
		feedTypesMap[v.Id] = v.Name
	}
	res := entitys.UcUserFeedback_pb2e(req.Data[0])
	res.QuestionTypeName = feedTypesMap[req.Data[0].QuestionTypeId]
	return res, err
}

// QueryUcFeedBackList 用户反馈列表
func (s UcFeedBackService) QueryUcFeedBackList(lang, appKey, tenantId string, filter entitys.UcUserFeedbackQuery) ([]map[string]interface{}, int64, error) {
	var resultList = []entitys.UcUserFeedbackEntitys{}
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.UcFeedbackService.FeedbackLists(s.Ctx, &protosService.UcUserFeedbackListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query: &protosService.UcUserFeedback{
			UserId: filter.UserId,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	feedTypes, err := s.GetFeedBackQuestionType(lang, appKey, tenantId, 0)
	feedTypesMap := make(map[int64]string)
	for _, v := range feedTypes {
		feedTypesMap[v.Id] = v.Name
	}

	for _, item := range rep.Data {
		row := entitys.UcUserFeedback_pb2e(item)
		if v, ok := feedTypesMap[item.QuestionTypeId]; ok {
			row.QuestionTypeName = v
		}
		resultList = append(resultList, row)
	}

	dateList := map[string][]entitys.UcUserFeedbackEntitys{}
	for _, row := range resultList {
		createAt := time.Unix(row.CreatedAt, 0)

		date := iotutil.TimeFormat(createAt)
		list, ok := dateList[date]
		if !ok {
			list = []entitys.UcUserFeedbackEntitys{}
		}
		list = append(list, row)
		dateList[date] = list
	}

	resultArr := []map[string]interface{}{}
	for key, value := range dateList {
		//根据sort进行排序
		sort.Slice(value, func(i, j int) bool {
			return value[i].CreatedAt > value[j].CreatedAt
		})
		resultArr = append(resultArr, map[string]interface{}{
			"date": key,
			"list": value,
		})
	}
	//根据sort进行排序
	sort.Slice(resultArr, func(i, j int) bool {
		return resultArr[i]["date"].(string) > resultArr[j]["date"].(string)
	})

	return resultArr, rep.Total, err
}

// AddUcFeedBack 新增用户反馈
func (s UcFeedBackService) AddUcFeedBack(req entitys.UcUserFeedbackEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.UcUserFeedback_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.UcFeedbackService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改用户反馈
func (s UcFeedBackService) UpdateUcFeedBack(req entitys.UcUserFeedbackEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.UcFeedbackService.Update(s.Ctx, entitys.UcUserFeedback_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除用户反馈
func (s UcFeedBackService) DeleteUcFeedBack(req entitys.UcUserFeedbackFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.UcFeedbackService.Delete(s.Ctx, &protosService.UcUserFeedback{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// 反馈问题类型列表查询
func (s UcFeedBackService) GetFeedBackQuestionType(lang string, appKey, tenantId string, id int64) ([]*entitys.FeedBackQuestionType, error) {
	if appKey == "" {
		return nil, errors.New("appKey not found")
	}
	req, err := rpc.ClientOemFeedbackTypeService.Lists(s.Ctx, &protosService.OemFeedbackTypeListRequest{
		Query: &protosService.OemFeedbackType{
			AppKey: appKey,
			Id:     id,
		}})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_OEM_FEEDBACK_TYPE)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	res := make([]*entitys.FeedBackQuestionType, 0)
	for _, r := range req.Data {
		r.Name = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%v_name", lang, r.Id)], r.Name)
		res = append(res, entitys.FeedBackQuestionType_pb2e(lang, r))
	}
	return res, nil
}
