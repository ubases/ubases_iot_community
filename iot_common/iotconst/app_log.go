package iotconst

// APP日志类型
var (
	APP_OPERATE_LOG string = "操作日志"
	APP_ERROR_LOG   string = "错误日志"
)

var (
	LogEventMap = map[string]string{}

	APP_EVENT_REGISTER     string = "注册"
	APP_EVENT_LOGIN        string = "登录"
	APP_URI_CANCEL_ACCOUNT string = "注销账号"
	//APP_EVENT_REFRESH_TOKEN              string = "刷新token"
	//APP_EVENT_UPDATE_USER                string = "修改用户"
	APP_EVENT_CANCEL_ACCOUNT string = "注销账号"
	//APP_EVENT_FORGET_PASSWORD            string = "忘记密码"
	//APP_EVENT_SEND_SMS                   string = "发送短信验证码"
	//APP_EVENT_CHECK_CODE                 string = "校验验证码"
	//APP_EVENT_SEND_EMAIL                 string = "发送邮件"
	//APP_EVENT_CHECK_ACCOUNT              string = "校验账号是否注册"
	//APP_EVENT_CHANNEL_AUTH               string = "第三方登录"
	//APP_EVENT_CHANNEL_BIND               string = "第三方渠道注册"
	//APP_EVENT_ACCOUNT_BIND               string = "登录账号信息绑定"
	//APP_EVENT_ADD_CHANNEL_BIND           string = "增加第三方渠道账号绑定"
	//APP_EVENT_UNBIND_CHANNEL             string = "第三方渠道账号解绑"
	//APP_EVENT_HOME_ADD                   string = "添加家庭"
	//APP_EVENT_ROOM_ADD                   string = "添加家庭房间"
	//APP_EVENT_ROOM_SET_SORT              string = "家庭房间排序"
	//APP_EVENT_ROOM_DELETE                string = "家庭房间删除"
	//APP_EVENT_ROOM_UPDATE                string = "家庭房间设置"
	//APP_EVENT_HOME_SET_DEV_SORT          string = "家庭设备排序"
	//APP_EVENT_HOME_SEND_INVITATION_CODE  string = "发送邀请码"
	//APP_EVENT_HOME_JOIN_HOME             string = "成员加入家庭"
	//APP_EVENT_HOME_REMOVE_MEMBERS        string = "移除家庭成员"
	//APP_EVENT_HOME_TRANSFER_OWNERSHIP    string = "转移家庭所有权"
	//APP_EVENT_HOME_SET_ROLE              string = "家庭成员角色设置"
	//APP_EVENT_HOME_DELETE                string = "删除家庭"
	//APP_EVENT_HOME_QUIT                  string = "离开家庭"
	//APP_EVENT_HOME_UPDATE                string = "修改家庭"
	//APP_EVENT_DEV_TIMER_ADD              string = "新增定时器"
	//APP_EVENT_DEV_TIMER_UPDATE           string = "修改定时器"
	//APP_EVENT_DEV_TIMER_DELETE           string = "删除定时器"
	//APP_EVENT_DEV_TIMER_OPEN             string = "开启定时器"
	//APP_EVENT_DEV_TIMER_CLOSE            string = "关闭定时器"
	//APP_EVENT_DEV_COUNTDOWN_ADD          string = "新增倒计时"
	//APP_EVENT_DEV_COUNTDOWN_DELETE       string = "删除倒计时"
	//APP_EVENT_DEV_COUNTDOWN_OPEN         string = "开启倒计时"
	//APP_EVENT_DEV_COUNTDOWN_CLOSE        string = "关闭倒计时"
	//APP_EVENT_DEV_UPDATE                 string = "修改设备信息"
	//APP_EVENT_DEV_ADD                    string = "添加设备"
	//APP_EVENT_DEV_REMOVE                 string = "删除设备"
	//APP_EVENT_DEV_ROOM_REMOVE            string = "删除房间设备"
	//APP_EVENT_INTELLIGENCE_SAVE          string = "创建智能场景"
	//APP_EVENT_INTELLIGENCE_SET_SWITCH    string = "场景自动执行开关"
	//APP_EVENT_INTELLIGENCE_DEL           string = "场景删除"
	//APP_EVENT_INTELLIGENCE_LOG_DEL       string = "场景日志删除"
	//APP_EVENT_INTELLIGENCE_EXCUTE_SUBMIT string = "场景一键执行"
	//APP_EVENT_ALL_MESSAGE_DEL            string = "所有消息删除"
	//APP_EVENT_MESSAGE_DEL                string = "指定消息删除"
	//APP_EVENT_FEEDBACK_ADD               string = "反馈提交"
)

func init() {
	LogEventMap["/v1/platform/app/user/checkCode"] = "检查验证码"
	LogEventMap["/v1/platform/app/user/sendEmail"] = "发送邮件验证码"
	LogEventMap["/v1/platform/app/user/register"] = APP_EVENT_REGISTER //"用户注册"
	LogEventMap["/v1/platform/app/user/login"] = APP_EVENT_LOGIN       //"登录"
	LogEventMap["/v1/platform/app/user/forgetPassword"] = "忘记密码"
	LogEventMap["/v1/platform/app/user/cancelAccount"] = APP_EVENT_CANCEL_ACCOUNT //"注销账号"
	LogEventMap["/v1/platform/app/user/unbindChannel"] = "第三方渠道账号解绑"
	LogEventMap["/v1/platform/app/user/channelAuth"] = "第三方登录"
	LogEventMap["/v1/platform/app/user/addChannelBind"] = "第三方渠道账号绑定"
	LogEventMap["/v1/platform/app/user/updateUser"] = "修改用户信息"
	LogEventMap["/v1/platform/app/user/channelBind"] = "第三方渠道注册"
	LogEventMap["/v1/platform/app/user/refreshToken"] = "刷新token"
	LogEventMap["/v1/platform/app/user/accountBind"] = "登录账号信息绑定"
	LogEventMap["/v1/platform/app/room/add"] = "添加家庭房间"
	LogEventMap["/v1/platform/app/home/add"] = "添加家庭"
	LogEventMap["/v1/platform/app/room/delete"] = "家庭房间删除"
	LogEventMap["/v1/platform/app/home/transferOwnership"] = "转移家庭所有权"
	LogEventMap["/v1/platform/app/home/setDevSort"] = "家庭设备排序"
	LogEventMap["/v1/platform/app/home/update"] = "修改家庭"
	LogEventMap["/v1/platform/app/home/setRole"] = "家庭成员角色设置"
	LogEventMap["/v1/platform/app/home/removeMembers"] = "移除家庭成员"
	LogEventMap["/v1/platform/app/room/update"] = "家庭房间设置"
	LogEventMap["/v1/platform/app/home/joinHome"] = "成员加入家庭"
	LogEventMap["/v1/platform/app/home/quit"] = "离开家庭"
	LogEventMap["/v1/platform/app/home/sendInvitationCode"] = "发送邀请码"
	LogEventMap["/v1/platform/app/home/delete"] = "删除家庭"
	LogEventMap["/v1/platform/app/room/setSort"] = "家庭房间排序"
	LogEventMap["/v1/platform/app/dev/removeRoomDev"] = "房间移除设备"
	LogEventMap["/v1/platform/app/dev/update"] = "修改设备信息"
	LogEventMap["/v1/platform/app/dev/ota/setAutoUpgrade"] = "授权固件自动升级"
	LogEventMap["/v1/platform/app/dev/runRecord/clearDetail"] = "清空工作记录"
	LogEventMap["/v1/platform/app/dev/removeDev"] = "家庭删除设备"
	LogEventMap["/v1/platform/app/dev/timer/add"] = "新增定时器"
	LogEventMap["/v1/platform/app/dev/timer/delete"] = "删除定时器"
	LogEventMap["/v1/platform/app/dev/timer/close"] = "关闭定时器"
	LogEventMap["/v1/platform/app/dev/timer/open"] = "开启定时器"
	LogEventMap["/v1/platform/app/dev/timer/update"] = "修改定时器"
	LogEventMap["/v1/platform/app/dev/group/upsert"] = "创建修改群组"
	LogEventMap["/v1/platform/app/dev/group/remove"] = "移除群组"
	LogEventMap["/v1/platform/app/dev/group/execute"] = "群组控制"
	LogEventMap["/v1/platform/app/dev/group/executeSwitch"] = "群组开关控制"
	LogEventMap["/v1/platform/app/dev/addShared"] = "添加设备共享"
	LogEventMap["/v1/platform/app/dev/cancelReceiveShared"] = "取消接受共享"
	LogEventMap["/v1/platform/app/dev/cancelShare"] = "取消共享"
	LogEventMap["/v1/platform/app/dev/receiveShare"] = "接受共享"
	LogEventMap["/v1/platform/app/dev/countdown/open"] = "开启倒计时"
	LogEventMap["/v1/platform/app/dev/countdown/close"] = "关闭倒计时"
	LogEventMap["/v1/platform/app/dev/countdown/delete"] = "清除设备的倒计时任务"
	LogEventMap["/v1/platform/app/dev/countdown/add"] = "新增倒计时"
	LogEventMap["/v1/platform/app/feedback/add"] = "反馈提交"
	LogEventMap["/v1/platform/app/intelligence/logDel"] = "清空智能场景日志"
	LogEventMap["/v1/platform/app/intelligence/save"] = "创建/编辑智能场景"
	LogEventMap["/v1/platform/app/intelligence/del"] = "删除智能场景"
	LogEventMap["/v1/platform/app/intelligence/setSwitch"] = "自动场景开关"
	LogEventMap["/v1/platform/app/intelligence/excute/submit"] = "一键执行场景任务"
	LogEventMap["/v1/platform/app/common/uploadPic"] = "上传图片"

}

//
//func init() {
//	LogEventMap["注册"] = "/v1/platform/app/user/register"                  //注册
//	LogEventMap["登录"] = "/v1/platform/app/user/login"                     //登录
//	LogEventMap["刷新token"] = "/v1/platform/app/user/refreshToken"         //刷新token
//	LogEventMap["修改用户"] = "/v1/platform/app/user/updateUser"              //修改用户
//	LogEventMap["注销账号"] = "/v1/platform/app/user/cancelAccount"           //注销账号
//	LogEventMap["忘记密码"] = "/v1/platform/app/user/forgetPassword"          //忘记密码
//	LogEventMap["发送短信验证码"] = "/v1/platform/app/user/sendSms"              //发送短信验证码
//	LogEventMap["校验验证码"] = "/v1/platform/app/user/checkCode"              //校验验证码
//	LogEventMap["发送邮件"] = "/v1/platform/app/user/sendEmail"               //发送邮件
//	LogEventMap["校验账号是否注册"] = "/v1/platform/app/user/checkAccount"        //校验账号是否注册
//	LogEventMap["第三方登录"] = "/v1/platform/app/user/channelAuth"            //第三方登录
//	LogEventMap["第三方渠道注册"] = "/v1/platform/app/user/channelBind"          //第三方渠道注册
//	LogEventMap["登录账号信息绑定"] = "/v1/platform/app/user/accountBind"         //登录账号信息绑定
//	LogEventMap["增加第三方渠道账号绑定"] = "/v1/platform/app/user/addChannelBind"   //增加第三方渠道账号绑定
//	LogEventMap["第三方渠道账号解绑"] = "/v1/platform/app/user/unbindChannel"      //第三方渠道账号解绑
//	LogEventMap["添加家庭"] = "/v1/platform/app/home/add"                     //添加家庭
//	LogEventMap["添加家庭房间"] = "/v1/platform/app/room/add"                   //添加家庭房间
//	LogEventMap["家庭房间排序"] = "/v1/platform/app/room/setSort"               //家庭房间排序
//	LogEventMap["家庭房间删除"] = "/v1/platform/app/room/delete"                //家庭房间删除
//	LogEventMap["家庭房间设置"] = "/v1/platform/app/room/update"                //家庭房间设置
//	LogEventMap["家庭设备排序"] = "/v1/platform/app/home/setDevSort"            //家庭设备排序
//	LogEventMap["发送邀请码"] = "/v1/platform/app/home/sendInvitationCode"     //发送邀请码
//	LogEventMap["成员加入家庭"] = "/v1/platform/app/home/joinHome"              //成员加入家庭
//	LogEventMap["移除家庭成员"] = "/v1/platform/app/home/removeMembers"         //移除家庭成员
//	LogEventMap["转移家庭所有权"] = "/v1/platform/app/home/transferOwnership"    //转移家庭所有权
//	LogEventMap["家庭成员角色设置"] = "/v1/platform/app/home/setRole"             //家庭成员角色设置
//	LogEventMap["删除家庭"] = "/v1/platform/app/home/delete"                  //删除家庭
//	LogEventMap["离开家庭"] = "/v1/platform/app/home/quit"                    //离开家庭
//	LogEventMap["修改家庭"] = "/v1/platform/app/home/update"                  //修改家庭
//	LogEventMap["新增定时器"] = "/v1/platform/app/dev/timer/add"               //新增定时器
//	LogEventMap["修改定时器"] = "/v1/platform/app/dev/timer/update"            //修改定时器
//	LogEventMap["删除定时器"] = "/v1/platform/app/dev/timer/delete"            //删除定时器
//	LogEventMap["开启定时器"] = "/v1/platform/app/dev/timer/open"              //开启定时器
//	LogEventMap["关闭定时器"] = "/v1/platform/app/dev/timer/close"             //关闭定时器
//	LogEventMap["新增倒计时"] = "/v1/platform/app/dev/countdown/add"           //新增倒计时
//	LogEventMap["删除倒计时"] = "/v1/platform/app/dev/countdown/delete"        //删除倒计时
//	LogEventMap["开启倒计时"] = "/v1/platform/app/dev/countdown/open"          //开启倒计时
//	LogEventMap["关闭倒计时"] = "/v1/platform/app/dev/countdown/close"         //关闭倒计时
//	LogEventMap["修改设备信息"] = "/v1/platform/app/dev/update"                 //修改设备信息
//	LogEventMap["添加设备"] = "/v1/platform/app/dev/addDev"                   //添加设备
//	LogEventMap["删除设备"] = "/v1/platform/app/dev/removeDev"                //删除设备
//	LogEventMap["删除房间设备"] = "/v1/platform/app/dev/removeRoomDev"          //删除房间设备
//	LogEventMap["创建智能场景"] = "/v1/platform/app/intelligence/save"          //创建智能场景
//	LogEventMap["场景自动执行开关"] = "/v1/platform/app/intelligence/switch"      //场景自动执行开关
//	LogEventMap["场景删除"] = "/v1/platform/app/intelligence/del"             //场景删除
//	LogEventMap["场景日志删除"] = "/v1/platform/app/intelligence/log/del"       //场景日志删除
//	LogEventMap["场景一键执行"] = "/v1/platform/app/intelligence/excute/submit" //场景一键执行
//	LogEventMap["所有消息删除"] = "/v1/platform/app/message/allDel"             //所有消息删除
//}

//APP事件名
//var (
//	APP_EVENT_REGISTER     string = "注册"
//	APP_EVENT_LOGIN        string = "登录"
//	APP_URI_CANCEL_ACCOUNT string = "注销账号"
//	APP_EVENT_REFRESH_TOKEN              string = "刷新token"
//	APP_EVENT_UPDATE_USER                string = "修改用户"
//	APP_EVENT_CANCEL_ACCOUNT             string = "注销账号"
//	APP_EVENT_FORGET_PASSWORD            string = "忘记密码"
//	APP_EVENT_SEND_SMS                   string = "发送短信验证码"
//	APP_EVENT_CHECK_CODE                 string = "校验验证码"
//	APP_EVENT_SEND_EMAIL                 string = "发送邮件"
//	APP_EVENT_CHECK_ACCOUNT              string = "校验账号是否注册"
//	APP_EVENT_CHANNEL_AUTH               string = "第三方登录"
//	APP_EVENT_CHANNEL_BIND               string = "第三方渠道注册"
//	APP_EVENT_ACCOUNT_BIND               string = "登录账号信息绑定"
//	APP_EVENT_ADD_CHANNEL_BIND           string = "增加第三方渠道账号绑定"
//	APP_EVENT_UNBIND_CHANNEL             string = "第三方渠道账号解绑"
//	APP_EVENT_HOME_ADD                   string = "添加家庭"
//	APP_EVENT_ROOM_ADD                   string = "添加家庭房间"
//	APP_EVENT_ROOM_SET_SORT              string = "家庭房间排序"
//	APP_EVENT_ROOM_DELETE                string = "家庭房间删除"
//	APP_EVENT_ROOM_UPDATE                string = "家庭房间设置"
//	APP_EVENT_HOME_SET_DEV_SORT          string = "家庭设备排序"
//	APP_EVENT_HOME_SEND_INVITATION_CODE  string = "发送邀请码"
//	APP_EVENT_HOME_JOIN_HOME             string = "成员加入家庭"
//	APP_EVENT_HOME_REMOVE_MEMBERS        string = "移除家庭成员"
//	APP_EVENT_HOME_TRANSFER_OWNERSHIP    string = "转移家庭所有权"
//	APP_EVENT_HOME_SET_ROLE              string = "家庭成员角色设置"
//	APP_EVENT_HOME_DELETE                string = "删除家庭"
//	APP_EVENT_HOME_QUIT                  string = "离开家庭"
//	APP_EVENT_HOME_UPDATE                string = "修改家庭"
//	APP_EVENT_DEV_TIMER_ADD              string = "新增定时器"
//	APP_EVENT_DEV_TIMER_UPDATE           string = "修改定时器"
//	APP_EVENT_DEV_TIMER_DELETE           string = "删除定时器"
//	APP_EVENT_DEV_TIMER_OPEN             string = "开启定时器"
//	APP_EVENT_DEV_TIMER_CLOSE            string = "关闭定时器"
//	APP_EVENT_DEV_COUNTDOWN_ADD          string = "新增倒计时"
//	APP_EVENT_DEV_COUNTDOWN_DELETE       string = "删除倒计时"
//	APP_EVENT_DEV_COUNTDOWN_OPEN         string = "开启倒计时"
//	APP_EVENT_DEV_COUNTDOWN_CLOSE        string = "关闭倒计时"
//	APP_EVENT_DEV_UPDATE                 string = "修改设备信息"
//	APP_EVENT_DEV_ADD                    string = "添加设备"
//	APP_EVENT_DEV_REMOVE                 string = "删除设备"
//	APP_EVENT_DEV_ROOM_REMOVE            string = "删除房间设备"
//	APP_EVENT_INTELLIGENCE_SAVE          string = "创建智能场景"
//	APP_EVENT_INTELLIGENCE_SET_SWITCH    string = "场景自动执行开关"
//	APP_EVENT_INTELLIGENCE_DEL           string = "场景删除"
//	APP_EVENT_INTELLIGENCE_LOG_DEL       string = "场景日志删除"
//	APP_EVENT_INTELLIGENCE_EXCUTE_SUBMIT string = "场景一键执行"
//	APP_EVENT_ALL_MESSAGE_DEL            string = "所有消息删除"
//	APP_EVENT_MESSAGE_DEL                string = "指定消息删除"
//	APP_EVENT_FEEDBACK_ADD               string = "反馈提交"
//)

//// APP URI
//var (
//	APP_URI_REGISTER                   string = "/v1/platform/app/user/register"              //注册
//	APP_URI_LOGIN                      string = "/v1/platform/app/user/login"                 //登录
//	APP_URI_REFRESH_TOKEN              string = "/v1/platform/app/user/refreshToken"          //刷新token
//	APP_URI_UPDATE_USER                string = "/v1/platform/app/user/updateUser"            //修改用户
//	APP_URI_CANCEL_ACCOUNT             string = "/v1/platform/app/user/cancelAccount"         //注销账号
//	APP_URI_FORGET_PASSWORD            string = "/v1/platform/app/user/forgetPassword"        //忘记密码
//	APP_URI_SEND_SMS                   string = "/v1/platform/app/user/sendSms"               //发送短信验证码
//	APP_URI_SEND_CHECK_CODE            string = "/v1/platform/app/user/checkCode"             //校验验证码
//	APP_URI_SEND_SEND_EMAIL            string = "/v1/platform/app/user/sendEmail"             //发送邮件
//	APP_URI_SEND_CHECK_ACCOUNT         string = "/v1/platform/app/user/checkAccount"          //校验账号是否注册
//	APP_URI_SEND_CHANNEL_AUTH          string = "/v1/platform/app/user/channelAuth"           //第三方登录
//	APP_URI_SEND_CHANNEL_BIND          string = "/v1/platform/app/user/channelBind"           //第三方渠道注册
//	APP_URI_SEND_ACCOUNT_BIND          string = "/v1/platform/app/user/accountBind"           //登录账号信息绑定
//	APP_URI_ADD_CHANNEL_BIND           string = "/v1/platform/app/user/addChannelBind"        //增加第三方渠道账号绑定
//	APP_URI_UNBIND_CHANNEL             string = "/v1/platform/app/user/unbindChannel"         //第三方渠道账号解绑
//	APP_URI_HOME_ADD                   string = "/v1/platform/app/home/add"                   //添加家庭
//	APP_URI_ROOM_ADD                   string = "/v1/platform/app/room/add"                   //添加家庭房间
//	APP_URI_ROOM_SET_SORT              string = "/v1/platform/app/room/setSort"               //家庭房间排序
//	APP_URI_ROOM_SET_DELETE            string = "/v1/platform/app/room/delete"                //家庭房间删除
//	APP_URI_ROOM_SET_UPDATE            string = "/v1/platform/app/room/update"                //家庭房间设置
//	APP_URI_HOME_SET_DEV_SORT          string = "/v1/platform/app/home/setDevSort"            //家庭设备排序
//	APP_URI_HOME_SEND_INVITATION_CODE  string = "/v1/platform/app/home/sendInvitationCode"    //发送邀请码
//	APP_URI_HOME_SEND_JOIN_HOME        string = "/v1/platform/app/home/joinHome"              //成员加入家庭
//	APP_URI_HOME_REMOVE_MEMBERS        string = "/v1/platform/app/home/removeMembers"         //移除家庭成员
//	APP_URI_HOME_TRANSFER_OWNERSHIP    string = "/v1/platform/app/home/transferOwnership"     //转移家庭所有权
//	APP_URI_HOME_SET_ROLE              string = "/v1/platform/app/home/setRole"               //家庭成员角色设置
//	APP_URI_HOME_DELETE                string = "/v1/platform/app/home/delete"                //删除家庭
//	APP_URI_HOME_QUIT                  string = "/v1/platform/app/home/quit"                  //离开家庭
//	APP_URI_HOME_UPDATE                string = "/v1/platform/app/home/update"                //修改家庭
//	APP_URI_DEV_TIMER_ADD              string = "/v1/platform/app/dev/timer/add"              //新增定时器
//	APP_URI_DEV_TIMER_UPDATE           string = "/v1/platform/app/dev/timer/update"           //修改定时器
//	APP_URI_DEV_TIMER_DELETE           string = "/v1/platform/app/dev/timer/delete"           //删除定时器
//	APP_URI_DEV_TIMER_OPEN             string = "/v1/platform/app/dev/timer/open"             //开启定时器
//	APP_URI_DEV_TIMER_CLOSE            string = "/v1/platform/app/dev/timer/close"            //关闭定时器
//	APP_URI_DEV_COUNTDOWN_ADD          string = "/v1/platform/app/dev/countdown/add"          //新增倒计时
//	APP_URI_DEV_COUNTDOWN_DELETE       string = "/v1/platform/app/dev/countdown/delete"       //删除倒计时
//	APP_URI_DEV_COUNTDOWN_OPEN         string = "/v1/platform/app/dev/countdown/open"         //开启倒计时
//	APP_URI_DEV_COUNTDOWN_CLOSE        string = "/v1/platform/app/dev/countdown/close"        //关闭倒计时
//	APP_URI_DEV_UPDATE                 string = "/v1/platform/app/dev/update"                 //修改设备信息
//	APP_URI_DEV_ADD                    string = "/v1/platform/app/dev/addDev"                 //添加设备
//	APP_URI_DEV_REMOVE                 string = "/v1/platform/app/dev/removeDev"              //删除设备
//	APP_URI_DEV_ROOM_REMOVE            string = "/v1/platform/app/dev/removeRoomDev"          //删除房间设备
//	APP_URI_INTELLIGENCE_SAVE          string = "/v1/platform/app/intelligence/save"          //创建智能场景
//	APP_URI_INTELLIGENCE_SET_SWITCH    string = "/v1/platform/app/intelligence/switch"        //场景自动执行开关
//	APP_URI_INTELLIGENCE_DEL           string = "/v1/platform/app/intelligence/del"           //场景删除
//	APP_URI_INTELLIGENCE_LOG_DEL       string = "/v1/platform/app/intelligence/log/del"       //场景日志删除
//	APP_URI_INTELLIGENCE_EXCUTE_SUBMIT string = "/v1/platform/app/intelligence/excute/submit" //场景一键执行
//	APP_URI_ALL_MESSAGE_DEL            string = "/v1/platform/app/message/allDel"             //所有消息删除
//	APP_URI_MESSAGE_DEL                string = "/v1/platform/app/message/del"                //指定消息删除
//	APP_URI_FEEDBACK_ADD               string = "/v1/platform/app/feedback/add"               //反馈提交
//)
//
//var (
//	LogEventMap = map[string]string{
//		APP_URI_REGISTER:                  APP_EVENT_REGISTER,
//		APP_URI_LOGIN:                     APP_EVENT_LOGIN,
//		APP_URI_CANCEL_ACCOUNT:            APP_EVENT_CANCEL_ACCOUNT,
//		APP_URI_FORGET_PASSWORD:           APP_EVENT_FORGET_PASSWORD,
//		APP_URI_SEND_CHANNEL_AUTH:         APP_EVENT_CHANNEL_AUTH,
//		APP_URI_SEND_CHANNEL_BIND:         APP_EVENT_CHANNEL_BIND,
//		APP_URI_HOME_ADD:                  APP_EVENT_HOME_ADD,
//		APP_URI_HOME_SEND_INVITATION_CODE: APP_EVENT_HOME_SEND_INVITATION_CODE,
//		APP_URI_HOME_TRANSFER_OWNERSHIP:   APP_EVENT_HOME_TRANSFER_OWNERSHIP,
//		APP_URI_DEV_ADD:                   APP_EVENT_DEV_ADD,
//		APP_URI_DEV_REMOVE:                APP_EVENT_DEV_REMOVE,
//		APP_URI_INTELLIGENCE_SAVE:         APP_EVENT_INTELLIGENCE_SAVE,
//		APP_URI_INTELLIGENCE_DEL:          APP_EVENT_INTELLIGENCE_DEL,
//		APP_URI_FEEDBACK_ADD:              APP_EVENT_FEEDBACK_ADD,
//	}
//)
