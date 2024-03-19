package services

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/app/entitys"
	services "cloud_platform/iot_cloud_api_service/controls/global"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strings"
)

type FeedbackService struct {
	Ctx context.Context
}

func (s FeedbackService) SetContext(ctx context.Context) FeedbackService {
	s.Ctx = ctx
	return s
}

// GetUcFeedbackDetail 用户反馈详细
func (s FeedbackService) GetUcFeedbackDetail(tenantId, id string) (*entitys.UcUserFeedbackDetails, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.UcFeedbackService.FeedbackDetails(s.Ctx, &protosService.UcUserFeedbackFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	//获取操作人昵称
	nickName := ""
	if req.ReplyList != nil && len(req.ReplyList) > 0 {
		nickName = s.getOperatorNickName(req.ReplyList[0].Operator)
	}
	appMap := map[string]string{req.AppInfo.AppName: req.AppInfo.AppName}
	req.AppInfo.AppName = s.getAppMaps(appMap)[req.AppInfo.AppName]

	typeMap := s.getQuestionTypeMap(tenantId, req.Data.AppKey)
	req.Data.QuestionTypeName = typeMap[req.Data.QuestionTypeId]

	//匹配IOS手机型号
	if strings.Index(req.AppInfo.MobileModel, "iPhone") != -1 {
		iosModuleModes, _ := new(services.DictTempData).GetDictByCode(iotconst.Dict_type_ios_mobile_mode)
		if v := iosModuleModes.ValueStr(req.AppInfo.MobileModel); v != "" {
			req.AppInfo.MobileModel = v
		}
	}
	return entitys.UcUserFeedbackDetails_pb2e(req, nickName), err
}

func (s *FeedbackService) getOperatorNickName(operator int64) string {
	//查询回复用户名称
	res, err := rpc.ClientOpenUserService.FindById(s.Ctx, &protosService.OpenUserFilter{
		Id: operator,
	})
	if err != nil {
		return ""
	}
	if res.Code != 200 && len(res.Data) > 0 {
		return ""
	}
	return res.Data[0].UserNickname
}

// QueryUcFeedbackList 用户反馈列表
func (s FeedbackService) QueryUcFeedbackList(filter entitys.UcUserFeedbackQuery, tenantId, lang string) ([]*entitys.UcUserFeedbackEntitys, int64, error) {
	var resultList = []*entitys.UcUserFeedbackEntitys{}
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.UcFeedbackService.Lists(s.Ctx, &protosService.UcUserFeedbackListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query: &protosService.UcUserFeedback{
			ProductKey:     filter.Query.ProductKey,
			QuestionTypeId: filter.Query.TypeId, //产品类型
			AppKey:         filter.Query.AppKey,
			TenantId:       tenantId,
			Status:         filter.Query.Status,
			TimeQueryMode:  filter.Query.TimeQueryMode, //字典：time_query_mode 1= 全部 =2 最近一周 =3 最近一个月 =4 最近三个月
		},
		OrderKey:  "created_at",
		OrderDesc: "desc",
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	var appKeysMap map[string]string = make(map[string]string)
	for _, item := range rep.Data {
		if _, ok := appKeysMap[item.AppKey]; !ok {
			appKeysMap[item.AppKey] = item.AppKey
		}
	}

	appMap := s.getAppMaps(appKeysMap)
	proMap := s.getProducts()
	//proCached := services.ProductCachedData{}

	typeMap := s.getQuestionTypeMap(tenantId, "")

	for _, item := range rep.Data {
		rowItem := entitys.UcUserFeedback_pb2e(item)
		//读取应用名称
		if val, ok := appMap[item.AppKey]; ok {
			rowItem.AppName = val
		}
		if val, ok := proMap[item.ProductKey]; ok {
			rowItem.ProductName = val
		}
		if val, ok := typeMap[item.QuestionTypeId]; ok {
			rowItem.QuestionTypeName = val
		}
		resultList = append(resultList, rowItem)
	}
	return resultList, rep.Total, err
}

func (s FeedbackService) getAppMaps(appKeysMap map[string]string) map[string]string {
	var appKeys []string
	for key, _ := range appKeysMap {
		appKeys = append(appKeys, key)
	}
	//查询开发者的APP信息
	apps, err := rpc.ClientOemAppService.Lists(s.Ctx, &protosService.OemAppListRequest{
		Query: &protosService.OemApp{
			AppKeys: appKeys,
		},
	})
	if err != nil {
		return nil
	}
	if apps.Code != 200 {
		return nil
	}
	var appMap map[string]string = make(map[string]string)
	for _, app := range apps.Data {
		appMap[app.AppKey] = app.Name
	}
	return appMap
}

func (s FeedbackService) getProducts() map[string]string {
	resMap := make(map[string]string)
	rep, err := rpc.ClientOpmProductService.Lists(s.Ctx, &protosService.OpmProductListRequest{
		IsPlatform: false,
	})
	if err != nil || rep.Code != 200 {
		return resMap
	}

	for _, app := range rep.Data {
		resMap[app.ProductKey] = app.Name
	}
	return resMap
}

func (s FeedbackService) getQuestionTypeMap(tenantId, appKey string) map[int64]string {
	resMap := make(map[int64]string)
	//查询开发者的APP信息
	rep, err := rpc.ClientFeedbackTypeService.Lists(s.Ctx, &protosService.OemFeedbackTypeListRequest{
		Query: &protosService.OemFeedbackType{
			AppKey:   appKey,
			TenantId: tenantId,
		},
	})
	if err != nil {
		return resMap
	}
	if rep.Code != 200 {
		return resMap
	}
	for _, app := range rep.Data {
		resMap[app.Id] = app.Name
	}
	return resMap
}

// FeedbackReplySubmit 提交反馈
func (s FeedbackService) FeedbackReplySubmit(req entitys.FeedbackReplySubmit) error {
	if req.FeedbackID == 0 {
		return errors.New("feedbackId not found")
	}
	pictures := ""
	videos := ""
	if req.Pictures != nil {
		pictures = strings.Join(req.Pictures, ",")
	}
	if req.Videos != nil {
		videos = strings.Join(req.Videos, ",")
	}
	feedback, err := rpc.UcFeedbackService.FindById(s.Ctx, &protosService.UcUserFeedbackFilter{Id: req.FeedbackID})
	if err != nil {
		return err
	}
	if feedback.Code != 200 {
		return errors.New(feedback.Message)
	}
	if len(feedback.Data) == 0 {
		return errors.New("未找到反馈信息")
	}
	tenantId := feedback.Data[0].TenantId
	appKey := feedback.Data[0].AppKey
	userId := feedback.Data[0].UserId
	res, err := rpc.UcFeedbackResultService.Create(s.Ctx, &protosService.UcUserFeedbackResult{
		FeedbackId:    req.FeedbackID,
		HandleStatus:  3, //反馈即已处理
		NotifyFlag:    1, //默认通知
		NotifyContent: req.Content,
		Pictures:      pictures,
		Videos:        videos,
		CreatedBy:     req.Operator,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	//推送反馈消息通知
	controls.SendFeedbackReplyMessage(controls.MessageAppInfo{AppKey: appKey, TenantId: tenantId}, userId, iotutil.ToString(req.FeedbackID))
	return nil
}
