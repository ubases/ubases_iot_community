package google

import (
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
	"log"
)

func init1() {
	GetSmartHome().RegisterCredentialsFile("783946466434580480", "./conf/my-smart-home-64989-76b8028a81f7.json")
	if err := GetSmartHome().RegisterOrUpdateDevice("783946466434580480", BasicDevice{
		Id: "1234567890",
		Name: proto.DeviceName{
			Name: "Light 1",
		},
		Type: proto.DeviceTypeLight,
		Traits: []Trait{
			OnOffTrait{
				CommandOnlyOnOff: false,
				OnExecuteChange: func(ctx Context, state bool) proto.DeviceError {
					return nil
				},
				OnStateHandler: func(ctx Context) (bool, proto.ErrorCode) {
					return false, nil
				},
			}},
		Info: proto.DeviceInfo{
			HwVersion: "1.0",
		},
	}); err != nil {
		log.Fatal(err)
	}

	if err := GetSmartHome().RegisterOrUpdateDevice("783946466434580480", BasicDevice{
		Id: "1234567891",
		Name: proto.DeviceName{
			Name: "Blinds 1",
		},
		Type: proto.DeviceTypeBlinds,
		Traits: []Trait{
			MultiDirectionOpenCloseTrait{
				DiscreteOnlyOpenClose: false,
				OpenDirection:         []OpenCloseTraitDirection{OpenCloseTraitDirectionUp, OpenCloseTraitDirectionDown},
				QueryOnlyOpenClose:    false,
				OnExecuteChange: func(ctx Context, openPercent float64, openDirection OpenCloseTraitDirection) proto.DeviceError {
					return nil
				},
				OnStateHandler: func(ctx Context) ([]OpenState, proto.ErrorCode) {
					curOpenState := OpenState{OpenPercent: 100.0, OpenDirection: OpenCloseTraitDirectionUp}
					return []OpenState{curOpenState}, nil
				},
			}},
		Info: proto.DeviceInfo{
			HwVersion: "1.0",
		},
	}); err != nil {
		log.Fatal(err)
	}

	if err := GetSmartHome().RegisterOrUpdateDevice("783946466434580480", BasicDevice{
		Id: "1234567892",
		Name: proto.DeviceName{
			Name: "Blinds 2",
		},
		Type: proto.DeviceTypeBlinds,
		Traits: []Trait{
			OpenCloseTrait{
				DiscreteOnlyOpenClose: true,
				QueryOnlyOpenClose:    false,
				OnExecuteChange: func(ctx Context, openPercent float64) proto.DeviceError {
					log.Println("Percent of", ctx.Target.DeviceName(), "should be set to", openPercent)
					return nil
				},
				OnStateHandler: func(ctx Context) (float64, proto.ErrorCode) {
					log.Println("query state of", ctx.Target.DeviceName())
					return 100, nil
				},
			}},
		Info: proto.DeviceInfo{
			HwVersion: "1.0",
		},
	}); err != nil {
		log.Fatal(err)
	}
}
