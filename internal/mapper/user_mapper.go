package mapper

import (
	"golang.org/x/crypto/bcrypt"
	"qlnv/internal/dto"
	"qlnv/internal/model"
)

func MapUserDtoToUser(dto dto.UserDTO, user model.User) model.User {
	user.UserName = dto.UserName
	user.Email = dto.Email
	user.Phone = dto.Phone
	user.FullName = dto.FullName
	user.Role = dto.Role
	user.Gender = dto.Gender
	user.Birthday = dto.Birthday
	if dto.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
	}
	return user
}
