package handler

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	routers "cloud_platform/iot_smart_speaker_service/router"
	"context"
)

type ClientInfo struct{}

// 创建AppLogUser
func (ci *ClientInfo) CreateClientInfo(ctx context.Context, req *proto.ClientInfoReq, resp *proto.ClientInfoResp) error {
	var err error
	for i := range req.ClientInfo {
		routers.SetClientInfo(req.ClientInfo[i].ClientId, req.ClientInfo[i].ClientSecret, req.ClientInfo[i].Domain)
	}
	CommonResponse(resp, err)
	return nil
}

func CommonResponse(resp *proto.ClientInfoResp, err error) {
	if err != nil {
		resp.Code = 400
		resp.Msg = err.Error()
	} else {
		resp.Code = 200
		resp.Msg = "Success"
	}
}
