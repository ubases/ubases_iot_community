package iotconst

var (
	APP_MESSAGE_DEVICE_ACTIVE           = "DeviceActiveMessage"            //设备激活
	APP_MESSAGE_DEVICE_OFFLINE          = "DeviceOffline"                  //设备离线
	APP_MESSAGE_DEVICE_ONLINE           = "DeviceOnline"                   //设备上线
	APP_MESSAGE_UPDATE_PASSWORD         = "UpdatePasswordMessage"          //修改密码，强制退出
	APP_MESSAGE_REMOVE_DEVICE           = "RemoveDeviceMessage"            //移除设备消息
	APP_MESSAGE_REMOVE_HOME             = "RemoveHomeMessage"              //家庭被移除
	APP_MESSAGE_JOIN_HOME               = "JoinHomeMessage"                //新成员加入
	APP_MESSAGE_QUIT_HOME               = "QuitHomeMessage"                //成员退出家庭
	APP_MESSAGE_REMOVE_MEMBERS          = "RemoveMembersMessage"           //被移除家庭
	APP_MESSAGE_FEEDBACK_REPLY          = "FeedbackReplyNotice"            //反馈回复通知
	APP_MESSAGE_DEVICE_UPGRADE_FAIL     = "DeviceUpgradeFail"              //OTA升级结果-失败
	APP_MESSAGE_DEVICE_UPGRADE_SUCCESS  = "DeviceUpgradeSuccess"           //OTA升级结果-成功
	APP_MESSAGE_ADD_SHARED              = "SendAddSharedMessage"           //添加共享消息
	APP_MESSAGE_RECEIVE_SHARED          = "SendReceiveSharedMessage"       //接受共享消息
	APP_MESSAGE_CANCEL_SHARED           = "SendCancelSharedMessage"        //取消共享消息
	APP_MESSAGE_REFUSE_RECEIVE_SHARED   = "SendRefuseReceiveSharedMessage" //拒绝接受共享消息
	APP_MESSAGE_DISBAND_GROUP_HOME      = "SendDisbandGroupMessage"        //解除群组
	APP_MESSAGE_AUTO_DISBAND_GROUP_HOME = "SendAutoDisbandGroupMessage"    //解除群组(自动解除群组)
)
