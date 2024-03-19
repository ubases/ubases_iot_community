package entitys

//联系我们通知请求
type ContractUsNoticeEntity struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
	Content     string `json:"content"`
}
