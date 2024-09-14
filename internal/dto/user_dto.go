package dto

type UserDTO struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Role     int    `json:"role"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}
