package entitys

// 1.0.3用户分析，响应数据
type OpenAppUserEntitys struct {
	AppUserTodayActive int32 `json:"appUserTodayActive"`
	AppUserToday       int32 `json:"appUserToday"`
	AppUserAll         int32 `json:"appUserAll"`
	AppUser            Data  `json:"appUser"`
	ActiveUser         Data  `json:"activeUser"`
}
