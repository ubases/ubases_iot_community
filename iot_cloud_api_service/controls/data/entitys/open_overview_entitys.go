package entitys

//云管平台数据概览响应实体
type OpenOverviewEntitys struct {
	ActiveDevice Data `json:"activeDevice"`
	AppUser      Data `json:"appUser"`
	DeviceFault  Data `json:"deviceFault"`
}
