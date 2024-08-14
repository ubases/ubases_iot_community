/**
 * @Author: hogan
 * @Date: 2022/3/30 12:44
 */
package ioterrs

type KVCode struct {
	Code int
	Msg  string
}

var (
	ERROR_FAIL                     = KVCode{Code: -1, Msg: "fail"}
	SUCCESS                        = KVCode{Code: 200, Msg: "ok"}
	ERROR                          = KVCode{Code: 500, Msg: "fail"}
	INVALID_PARAMS                 = KVCode{Code: 400, Msg: "请求参数错误"}
	ERROR_AUTH_CHECK_TOKEN_EXPIRE  = KVCode{Code: 401, Msg: "Token已过期"}
	ERROR_EXIST_ACOUNT             = KVCode{Code: 10021, Msg: "手机号或邮箱已经注册"}
	ERROR_EXIST_ACOUNT_NOT_BIND    = KVCode{Code: 100001, Msg: "该账号被注册了,无法进行绑定"}
	ERROR_AUTH_CHECK_TOKEN_FAIL    = KVCode{Code: 200001, Msg: "Token鉴权失败"}
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = KVCode{Code: 200002, Msg: "Token已超时"}
	ERROR_AUTH_TOKEN               = KVCode{Code: 200003, Msg: "Token生成失败"}
	ERROR_AUTH                     = KVCode{Code: 200004, Msg: "登录认证失败"}
	ERROR_ACCOUNTPASSWORD          = KVCode{Code: 200005, Msg: "账号或密码错误"}
	ERROR_NODELETIONPERMISSION     = KVCode{Code: 200006, Msg: "当前用户不是家庭所有者，不具备删除家庭的权限"}
	ERROR_ONLYFAMILY               = KVCode{Code: 200007, Msg: "仅有一个家庭时无法删除"}
	ERROR_HOMEDEVICEBIND           = KVCode{Code: 200008, Msg: "家庭有设备绑定不允许删除"}
	ERROR_INVALIDINVITATIONCODE    = KVCode{Code: 200009, Msg: "邀请码无效，请联系邀请者重新生成"}
	ERROR_ISMEMBER                 = KVCode{Code: 200010, Msg: "你已是该家庭成员，请勿重复加入"}
	ERROR_USER_IS_NOT_EXIST        = KVCode{Code: 200011, Msg: "用户信息不存在"}
	ERROR_DEVICE_ALREADY_SHARED    = KVCode{Code: 200012, Msg: "已共享给该用户，无须重复共享"}
	ERROR_INVALID_INVITATION_CODE  = KVCode{Code: 200013, Msg: "邀请码无效"}
	ERROR_NOT_BELONG_TO_USER       = KVCode{Code: 200014, Msg: "此账号不属于当前登录用户"}
	ERROR_NOT_SHARE_YOURSELF       = KVCode{Code: 200015, Msg: "亲不能共享给自己哦"}
	ERROR_INVITATION_SENT          = KVCode{Code: 200016, Msg: "已发送邀请，请等待对方处理"}
	ERROR_HOME_ID                  = KVCode{Code: 200017, Msg: "当前家庭id不属于当前用户"}
	ERROR_BIND_BY_OTHER_USER       = KVCode{Code: 200018, Msg: "当前第三方账号已被其他用户账号绑定"}
)
