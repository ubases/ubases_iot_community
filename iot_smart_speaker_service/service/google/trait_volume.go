package google

import (
	"errors"

	"cloud_platform/iot_smart_speaker_service/service/google/proto"
)

//音量

type VolumeTrait struct {
	//最大音量（假设基线为 0（静音）。Google 助理会相应地调整谓词命令（例如“让电视稍微大声点”）。
	VolumeMaxLevel int
	//指示设备是否可以将音量静音和取消静音。
	//“静音”是一个单独的选项，因为“静音”行为将音量设为 0，同时记住上一个音量，以便“取消静音”功能能够恢复该音量。
	//这反映在音量状态下。 如果音量为 5，并且用户为静音，则音量保持为 5 并且 isMuted 为 true。
	VolumeCanMuteAndUnmute bool
	//用户或制造商定义的默认音量的音量（以百分比为单位）。范围必须为 0-100。
	VolumeDefaultPercentage int
	//相关音量查询（例如“<设备名称>的音量调高”功能）的默认步进大小。
	LevelStepSize int
	//指示设备是以单向 (true) 还是双向 (false) 通信方式运行。
	//例如，如果控制器可以在发送请求后确认新设备状态，则此字段为 false。
	//如果无法确认请求是否成功执行或获取设备的状态（例如，如果设备是传统红外线遥控器），请将此字段设置为 true。
	CommandOnlyVolume             bool
	OnExecuteChangeMute           MuteCommand
	OnExecuteChangeSetVolume      SetVolumeCommand
	OnExecuteChangeVolumeRelative VolumeRelativeCommand

	OnStateHandler func(Context) (int, proto.ErrorCode)
}

func (t VolumeTrait) ValidateTrait() error {
	if t.OnExecuteChangeMute == nil {
		return errors.New("OnExecuteChangeMute cannot be nil")
	}
	if t.OnExecuteChangeSetVolume == nil {
		return errors.New("OnExecuteChangeSetVolume cannot be nil")
	}
	if t.OnExecuteChangeVolumeRelative == nil {
		return errors.New("OnExecuteChangeVolumeRelative cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}

	return nil
}
func (t VolumeTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_VOLUME
}

func (t VolumeTrait) TraitStates(ctx Context) []State {
	var state State
	state.Name = "currentVolume"
	state.Value, state.Error = t.OnStateHandler(ctx)
	return []State{state}
}

func (t VolumeTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChangeMute, t.OnExecuteChangeSetVolume, t.OnExecuteChangeVolumeRelative}
}

func (t VolumeTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "volumeMaxLevel",
			Value: t.VolumeMaxLevel,
		},
		{
			Name:  "volumeCanMuteAndUnmute",
			Value: t.VolumeCanMuteAndUnmute,
		},
		{
			Name:  "volumeDefaultPercentage",
			Value: t.VolumeDefaultPercentage,
		},
		{
			Name:  "levelStepSize",
			Value: t.LevelStepSize,
		},
		{
			Name:  "commandOnlyVolume",
			Value: t.CommandOnlyVolume,
		},
	}
}
