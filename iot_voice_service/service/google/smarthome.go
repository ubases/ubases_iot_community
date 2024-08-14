package google

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_voice_service/service/common"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"google.golang.org/api/homegraph/v1"

	"cloud_platform/iot_voice_service/service/google/proto"
)

var _SmartHomeOnce sync.Once
var _SmartHomeSingle *SmartHome

func GetSmartHome() *SmartHome {
	_SmartHomeOnce.Do(func() {
		_SmartHomeSingle = &SmartHome{
			agents:      make(map[string]agentContext),
			credentials: make(map[string]string),
		}
	})
	return _SmartHomeSingle
}

type agentContext struct {
	AgentUserId string
	Devices     map[string]Device
}

type SmartHome struct {
	lock   sync.RWMutex
	agents map[string]agentContext
	//service map[string]*homegraph.Service
	credentials map[string]string //userid->证书
	Aspects     []IntentAspect
}

func (s *SmartHome) AddAspects(i IntentAspect) {
	for _, o := range s.Aspects {
		if i.Intent == o.Intent {
			return
		}
	}
	s.Aspects = append(s.Aspects, i)
}

func (s *SmartHome) PreLoad(agentUserId string, Intent string) {
	for _, a := range s.Aspects {
		if a.Intent == Intent {
			a.Func(agentUserId)
		}
	}
}

func (s *SmartHome) TestHandle() {
	//var cmdstr = `{"inputs":[{"context":{"locale_country":"US","locale_language":"en"},"intent":"action.devices.EXECUTE","payload":{"commands":[{"devices":[{"customData":{"productKey":"PK7j2CkD"},"id":"1tPyiHmdIrhaWM"}],"execution":[{"command":"action.devices.commands.SetModes","params":{"updateModeSettings":{"countdown":"five"}}}]}]}}],"requestId":"5551902160073714756"}`
	//var cmdstr = `{"inputs":[{"intent":"action.devices.QUERY","payload":{"devices":[{"customData":{"productKey":"PK7j2CkD"},"id":"1tPyiHmdIrhaWM"},{"customData":{"productKey":"PKjBgN85"},"id":"A1sbCZVL2wmKJ5"}]}}],"requestId":"8588180855975679163"}`
	//var cmdstr = `{"inputs":[{"intent":"action.devices.DISCONNECT"}],"requestId":"6719617639275888665"}`
	var cmdstr = `{"inputs":[{"intent":"action.devices.QUERY","payload":{"devices":[{"customData":{"productKey":"PKjBgN85"},"id":"A1sbCZVL2wmKJ5"},{"customData":{"productKey":"PKjBgN85"},"id":"zOHuwLJYAUipp7"}]}}],"requestId":"4680044611410762883"}`
	req := proto.IntentMessageRequest{}
	err := json.Unmarshal([]byte(cmdstr), &req)
	if err != nil {
		return
	}
	res := proto.IntentMessageResponse{
		RequestId: req.RequestId,
	}
	agentUserId := "134076070113476608"

	for _, i := range req.Inputs {
		switch i.Intent {
		case IntentSync:
			s.PreLoad(agentUserId, i.Intent)
			res.Payload = s.handleSyncIntent(agentUserId)
		case IntentExecute:
			requestBody := proto.ExecRequest{}
			if err := json.Unmarshal(i.Payload, &requestBody); err == nil {
				res.Payload = s.handleExecuteIntent(agentUserId, requestBody)
			} else {
				res.Payload = proto.ErrorResponse{
					Status:    proto.CommandStatusError,
					ErrorCode: proto.ErrorCodeProtocolError.Error(),
				}
			}
		case IntentQuery:
			requestBody := proto.QueryRequest{}
			if err := json.Unmarshal(i.Payload, &requestBody); err == nil {
				res.Payload = s.handleQueryIntent(requestBody)
			} else {
				res.Payload = proto.ErrorResponse{
					Status:    proto.CommandStatusError,
					ErrorCode: proto.ErrorCodeProtocolError.Error(),
				}
			}
		case IntentDISCONNECT:
			s.handleDisconnectIntent(context.Background(), req.RequestId, agentUserId)
			iotlogger.LogHelper.Infof("userId %s DISCONNECT ", agentUserId)
		}
	}
}

func (s *SmartHome) Handle(c *gin.Context) {
	agentUserId := GetAgentUserIdFromHeader(c)
	if agentUserId == "" {
		iotgin.ResJSON(c, http.StatusBadRequest, errors.New("agent_user_id missing"))
		return
	}
	req := proto.IntentMessageRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	res := proto.IntentMessageResponse{
		RequestId: req.RequestId,
	}

	for _, i := range req.Inputs {
		switch i.Intent {
		case IntentSync:
			s.PreLoad(agentUserId, i.Intent)
			res.Payload = s.handleSyncIntent(agentUserId)
			iotgin.ResJSON(c, http.StatusOK, res)
		case IntentExecute:
			requestBody := proto.ExecRequest{}
			if err := json.Unmarshal(i.Payload, &requestBody); err == nil {
				res.Payload = s.handleExecuteIntent(agentUserId, requestBody)
			} else {
				res.Payload = proto.ErrorResponse{
					Status:    proto.CommandStatusError,
					ErrorCode: proto.ErrorCodeProtocolError.Error(),
				}
			}
			iotgin.ResJSON(c, http.StatusOK, res)
		case IntentQuery:
			requestBody := proto.QueryRequest{}
			if err := json.Unmarshal(i.Payload, &requestBody); err == nil {
				res.Payload = s.handleQueryIntent(requestBody)
			} else {
				res.Payload = proto.ErrorResponse{
					Status:    proto.CommandStatusError,
					ErrorCode: proto.ErrorCodeProtocolError.Error(),
				}
			}
			iotgin.ResJSON(c, http.StatusOK, res)
		case IntentDISCONNECT:
			s.handleDisconnectIntent(context.Background(), req.RequestId, agentUserId)
			iotlogger.LogHelper.Infof("userId %s DISCONNECT ", agentUserId)
			iotgin.ResJSON(c, http.StatusOK, nil)
			return
		}
	}
}

func (s *SmartHome) executeCommandForResponse(dev Device, ex proto.CommandRequest) proto.CommandResponse {
	// 遍历所有的设备特诊
	for _, trait := range dev.DeviceTraits() {
		// 遍历所有的命令
		for _, cmd := range trait.TraitCommands() {
			// 匹配命令
			if ex.Command == cmd.Name() {
				// 匹配到了命令就执行
				ctx := Context{Target: dev}
				res := cmd.Execute(ctx, ex.Params)
				proto.SetIds(&res, dev.DeviceId())
				online := common.GetDeviceOnline(dev.DeviceId())
				res.States.Online = online
				if !online {
					res.Status = string(proto.CommandStatusOFFLINE)
				} else {
					if res.ErrorCode == nil {
						res.Status = string(proto.CommandStatusSuccess)
					} else {
						res.Status = string(proto.CommandStatusError)
					}
				}
				return res
			}
		}
	}
	return proto.CommandResponse{
		ErrorCode: proto.ErrorCodeNotSupported,
	}
}

func (s *SmartHome) handleExecuteIntent(agentUserId string, req proto.ExecRequest) proto.ExecResponse {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var resCount int
	var responseBody proto.ExecResponse
	responses := make(chan proto.CommandResponse)
	// 根据用户代理ID找到对应的代理
	if agent, ok := s.agents[agentUserId]; ok {
		// 遍历请求的命令表
		for _, c := range req.Commands {
			// 遍历请求的设备列表
			for _, d := range c.Devices {
				//根据设备ID进行匹配
				if devCtx, ok := agent.Devices[d.ID]; ok {
					// 如果匹配到，则执行命令
					for _, e := range c.Execution {
						resCount++
						go func(s *SmartHome, ch chan proto.CommandResponse, dev Device, ex proto.CommandRequest) {
							ch <- s.executeCommandForResponse(dev, ex)
						}(s, responses, devCtx, e)
					}
				} else {
					responseBody.Commands = append(responseBody.Commands, proto.CommandResponse{
						ErrorCode: proto.ErrorCodeDeviceNotFound,
					})
				}
			}
		}
	}

	for len(responseBody.Commands) < resCount {
		responseBody.Commands = append(responseBody.Commands, <-responses)
	}

	//oresp, _ := json.Marshal(responseBody)
	//iotlogger.LogHelper.Infof("[%s] EXEC RESPONSE: %s\n", agentUserId, string(oresp))

	return responseBody
}

func (s *SmartHome) encodeDeviceForSyncResponse(dev Device) proto.Device {
	devTraits := dev.DeviceTraits()
	traits := make([]string, 0, len(devTraits))
	attributes := make(map[string]interface{})

	for _, t := range dev.DeviceTraits() {
		traits = append(traits, t.TraitName())
		for _, a := range t.TraitAttributes() {
			attributes[a.Name] = a.Value
		}
	}

	var info proto.DeviceInfo
	if p, ok := dev.(DeviceInfoProvider); ok {
		info = p.DeviceInfo()
	}

	var roomHint string
	if p, ok := dev.(DeviceRoomHintProvider); ok {
		roomHint = p.DeviceRoomHint()
	}

	d := proto.Device{
		Id:              dev.DeviceId(),
		Type:            dev.DeviceType(),
		Traits:          traits,
		Name:            dev.DeviceName(),
		DeviceInfo:      info,
		RoomHint:        roomHint,
		Attributes:      attributes,
		CustomData:      make(map[string]interface{}),
		WillReportState: false,
		OtherDeviceIds:  dev.GetOtherDeviceIds(),
	}
	for k, v := range dev.DeviceCustomData() {
		d.CustomData[k] = v
	}

	return d
}

func (s *SmartHome) handleSyncIntent(agentUserId string) proto.SyncResponse {
	response := proto.SyncResponse{
		AgentUserId: agentUserId,
		Devices:     make([]proto.Device, 0),
	}

	s.lock.RLock()
	defer s.lock.RUnlock()

	if agent, ok := s.agents[agentUserId]; ok {
		for _, d := range agent.Devices {
			response.Devices = append(response.Devices, s.encodeDeviceForSyncResponse(d))
		}
	}

	//oresp, _ := json.Marshal(response)
	//iotlogger.LogHelper.Infof("[%s] SYNC RESPONSE: %s\n", agentUserId, string(oresp))

	return response
}

func (s *SmartHome) handleQueryIntentOld(request proto.QueryRequest) proto.QueryResponse {
	res := proto.QueryResponse{
		Devices: make(map[string]map[string]interface{}),
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, r := range request.Devices {
		for _, a := range s.agents {
			if d, ok := a.Devices[r.ID]; ok {
				if _, ok := res.Devices[r.ID]; !ok {
					res.Devices[r.ID] = make(map[string]interface{})
				}
				ctx := Context{Target: d}
				for _, t := range d.DeviceTraits() {
					for _, s := range t.TraitStates(ctx) {
						res.Devices[d.DeviceId()][s.Name] = s.Value
						res.Devices[d.DeviceId()]["online"] = s.Error == nil
						if _, ok := res.Devices[d.DeviceId()]["status"]; ok &&
							res.Devices[d.DeviceId()]["status"] == proto.CommandStatusError {
							if s.Error != nil {
								res.Devices[d.DeviceId()]["status"] = proto.CommandStatusError
								res.Devices[d.DeviceId()]["errorCode"] = s.Error
							}
						}
					}
				}
			}
		}
	}

	//ores, _ := json.Marshal(res)
	//iotlogger.LogHelper.Infof("QUERY RESPONSE %s\n", string(ores))
	return res
}
func (s *SmartHome) handleQueryIntent(request proto.QueryRequest) proto.QueryResponse {
	res := proto.QueryResponse{
		Devices: make(map[string]map[string]interface{}),
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, r := range request.Devices {
		for _, a := range s.agents {
			if d, ok := a.Devices[r.ID]; ok {
				if _, ok := res.Devices[r.ID]; !ok {
					res.Devices[r.ID] = make(map[string]interface{})
				}
				ctx := Context{Target: d}
				online := d.GetDeviceOnlineStatus()
				res.Devices[d.DeviceId()]["online"] = online
				if !online {
					res.Devices[d.DeviceId()]["status"] = proto.CommandStatusOFFLINE
				} else {
					res.Devices[d.DeviceId()]["status"] = proto.CommandStatusSuccess
				}
				for _, t := range d.DeviceTraits() {
					for _, s := range t.TraitStates(ctx) {
						if s.Error == nil {
							res.Devices[d.DeviceId()][s.Name] = s.Value
						}
						//else {
						//	res.Devices[d.DeviceId()]["status"] = proto.CommandStatusError
						//	res.Devices[d.DeviceId()]["errorCode"] = s.Error
						//}
					}
				}
			}
		}
	}

	//ores, _ := json.Marshal(res)
	//iotlogger.LogHelper.Infof("QUERY RESPONSE %s\n", string(ores))
	return res
}

func (s *SmartHome) GetDeviceState(devId string) (map[string]map[string]interface{}, string) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var ret map[string]map[string]interface{}
	var userAgentId string
	for _, a := range s.agents {
		if d, ok := a.Devices[devId]; ok {
			userAgentId = a.AgentUserId
			ret = make(map[string]map[string]interface{})
			ret[devId] = make(map[string]interface{})
			ctx := Context{Target: d}
			for _, t := range d.DeviceTraits() {
				for _, s := range t.TraitStates(ctx) {
					if s.Error == nil {
						ret[devId][s.Name] = s.Value
					}
				}
			}
			break
		}
	}
	return ret, userAgentId
}

func (s *SmartHome) handleDisconnectIntent(ctx context.Context, requestId, agentUserID string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.agents, agentUserID)
	delete(s.credentials, agentUserID)
	return nil
}

// 允许重复调用
func (s *SmartHome) RegisterOrUpdateDevice(agentUserId string, dev Device) error {
	//参数验证
	reducedTraits := make(map[string]Trait)
	for _, t := range dev.DeviceTraits() {
		if _, ok := reducedTraits[t.TraitName()]; ok {
			return errors.New("duplicate trait found: " + t.TraitName())
		}
		if err := t.ValidateTrait(); err != nil {
			return err
		}
		reducedTraits[t.TraitName()] = t
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	if s.agents == nil {
		s.agents = make(map[string]agentContext)
	}

	if _, ok := s.agents[agentUserId]; !ok {
		s.agents[agentUserId] = agentContext{
			AgentUserId: agentUserId,
			Devices:     make(map[string]Device),
		}
	}
	s.agents[agentUserId].Devices[dev.DeviceId()] = dev

	return nil
}

func (s *SmartHome) RequestSync(ctx context.Context, agentUserID string) error {
	svr, err := s.GetService(agentUserID)
	if err != nil {
		return err
	}
	call := svr.Devices.RequestSync(&homegraph.RequestSyncDevicesRequest{
		AgentUserId: agentUserID,
		Async:       false,
	})
	call.Context(ctx)
	resp, err := call.Do()
	if err != nil {
		iotlogger.LogHelper.Error("error requesting sync", zap.String("agent_user_id", agentUserID), zap.Error(err))
		return err
	}
	if resp.ServerResponse.HTTPStatusCode != http.StatusOK {
		iotlogger.LogHelper.Error("failed request sync", zap.String("agent_user_id", agentUserID), zap.Int("status_code", resp.ServerResponse.HTTPStatusCode))
		return proto.ErrSyncFailed
	}
	return nil
}

// 报告设备状态
func (s *SmartHome) ReportState(ctx context.Context, devId string) error {
	mapData, agentUserID := s.GetDeviceState(devId)
	if agentUserID == "" || mapData == nil || mapData[devId] == nil {
		return nil
	}
	svr, err := s.GetService(agentUserID)
	if err != nil {
		return err
	}
	jsonState, err := json.Marshal(mapData)
	if err != nil {
		iotlogger.LogHelper.Error("error serializing device states to json", zap.String("agent_user_id", agentUserID), zap.Error(err))
	}
	req := homegraph.ReportStateAndNotificationRequest{
		AgentUserId: agentUserID,
		RequestId:   uuid.New().String(),
		Payload: &homegraph.StateAndNotificationPayload{
			Devices: &homegraph.ReportStateAndNotificationDevice{
				States: jsonState,
			},
		},
	}
	//jsonStr, err := json.Marshal(req)
	//fmt.Println(jsonStr)
	call := svr.Devices.ReportStateAndNotification(&req)
	call.Context(ctx)
	resp, err := call.Do()
	if err != nil {
		iotlogger.LogHelper.Error("error reporting state", zap.String("agent_user_id", agentUserID), zap.Error(err))
		return err
	}
	if resp.ServerResponse.HTTPStatusCode != http.StatusOK {
		iotlogger.LogHelper.Error("failed report state", zap.String("agent_user_id", agentUserID), zap.Int("status_code", resp.ServerResponse.HTTPStatusCode))
		return proto.ErrSyncFailed
	}
	return nil
}

func (s *SmartHome) RegisterCredentialsFile(agentUserID string, file string) error {
	//检查本地文件
	content, err := GetFileContent(file)
	if err != nil {
		//首先下载密钥文件
		content, err = DownloadFile(file)
		if err != nil {
			iotlogger.LogHelper.Error("RegisterCredentialsFile error.", zap.String("file", file), zap.Error(err))
			return err
		}
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.credentials[agentUserID] = string(content)
	return nil
}

func (s *SmartHome) GetService(agentUserID string) (*homegraph.Service, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	content, ok := s.credentials[agentUserID]
	if !ok {
		return nil, fmt.Errorf("userID: %s credentials is null", agentUserID)
	}
	ctx := context.Background()
	hgService, err := homegraph.NewService(ctx, option.WithCredentialsJSON([]byte(content)))
	if err != nil {
		iotlogger.LogHelper.Error("RegisterCredentialsFile error:", zap.String("agentUserID", agentUserID), zap.Error(err))
		return nil, fmt.Errorf("initializing homegraph error:%s", err.Error())
	}
	return hgService, nil
}

// 下载到文件
func DownloadFile(fileOssUrl string) ([]byte, error) {
	var resp *resty.Response
	var err error
	for i := 0; i < 3; i++ {
		client := resty.New()
		resp, err = client.R().Get(fileOssUrl)
		if err != nil {
			time.Sleep(6 * time.Second)
		} else if resp.StatusCode() == http.StatusOK {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	fileName := GetUrlFileName(fileOssUrl)
	if err := WriteFileContent(fileName, resp.Body()); err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func WriteFileContent(fileName string, data []byte) error {
	filePath := filepath.Join(iotconst.GetWorkTempDir(), fileName)
	localPath := filepath.Dir(filePath)
	if err := os.MkdirAll(localPath, os.ModePerm); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filePath, data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func GetUrlFileName(url string) string {
	fileName := ""
	urlSeg := strings.Split(url, "/")
	if len(urlSeg) > 0 {
		fileName = urlSeg[len(urlSeg)-1]
	}
	return fileName
}

func GetFileContent(fileOssUrl string) ([]byte, error) {
	fileName := GetUrlFileName(fileOssUrl)
	fileName = filepath.Join(iotconst.GetWorkTempDir(), fileName)
	buf, err := ioutil.ReadFile(fileName)
	return buf, err
}
