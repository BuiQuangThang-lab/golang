package dto

type UserInfo struct {
	Id          string `json:"id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	FullName    string `json:"full_name"`
	Role        string `json:"role"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	IsDelete    int    `json:"is_delete"`
	CreatedDate string `json:"create_at"`
}
