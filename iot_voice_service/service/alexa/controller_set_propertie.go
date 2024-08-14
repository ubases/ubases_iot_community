package alexa

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_voice_service/entitys"
	"errors"
	"time"
)

// 发现设备
func SetProperties(res *entitys.DirectiveRequet, data []byte, userId string, token string) error {
	svc := AlexaService{}
	devResp, err := svc.SetDeviceProperty(data, res.Directive.Header, userId, false, "alexa")
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set device properties user token error: ", err)
	}
	if err != nil {
		return err
	}
	return setPropertiesResponse(res, devResp)
}

// 发现设备响应
func setPropertiesResponse(res *entitys.DirectiveRequet, deviceResponse []entitys.DeviceResponse) error {
	properties := make([]entitys.AlexaControlProperties, 0)
	var err error
	if len(deviceResponse) > 0 {
		p := deviceResponse[0]
		if p.ErrorCode != "SUCCESS" {
			err = errors.New(p.Message)
		}
		for k, v := range p.Data {
			properties = append(properties, entitys.AlexaControlProperties{
				Namespace:                 res.Directive.Header.Namespace,
				Name:                      k,
				Value:                     iotutil.ToString(v),
				TimeOfSample:              time.Now().String(),
				UncertaintyInMilliseconds: 500,
			})
		}
	}
	res.Directive.Payload = entitys.AlexaControlPropertiesContext{Properties: properties}
	res.Directive.Header.Name = "Response"
	res.Directive.Header.Namespace = "Alexa"
	return err
}
