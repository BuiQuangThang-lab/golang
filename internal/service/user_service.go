package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"qlnv/internal/dto"
	"qlnv/internal/dto/filter"
	enum "qlnv/internal/enum"
	"qlnv/internal/mapper"
	"qlnv/internal/model"
	"qlnv/internal/repository"
	"qlnv/pkg/util"
	"time"
)

type UserService interface {
	GetListUser(userFilter filter.UserFilter) ([]dto.UserInfo, int64, error)
	CreateUser(user dto.UserDTO) error
	UpdateUser(id string, user dto.UserDTO) (*model.User, error)
	DeleteUser(id string) error
	GetById(id string) (*model.User, error)
	GetByUserName(username string) (*model.User, error)
	CheckPassword(hashedPassword, plainPassword string) bool
	Authenticate(username, password string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (u *userService) GetByUsername(username string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) GetListUser(filter filter.UserFilter) ([]dto.UserInfo, int64, error) {
	var userInfo []dto.UserInfo
	var users, total, err = u.repo.GetListUser(filter)
	if err != nil {
		return nil, 0, err
	}
	for _, u := range users {
		genderStr := enum.Gender(u.Gender).ConvertGender()
		roleStr := enum.Role(u.Role).ConvertRole()
		userInfo = append(userInfo, dto.UserInfo{
			Id:          u.Id,
			UserName:    u.UserName,
			Email:       u.Email,
			Phone:       u.Phone,
			FullName:    u.FullName,
			Role:        roleStr,
			Gender:      genderStr,
			Birthday:    u.Birthday,
			IsDelete:    u.IsDelete,
			CreatedDate: util.FormatDateTime(u.CreatedDate, "02/01/2006 15:04:00"),
		})
	}
	return userInfo, total, nil
}

func (u *userService) CreateUser(userDTO dto.UserDTO) error {
	var user model.User
	copier.Copy(&user, &userDTO)
	user.Id = uuid.New().String()
	user.IsDelete = 0
	user.CreatedDate = time.Now()
	if userDTO.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
	}
	var err = u.repo.CreateUser(&user)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"Name":  userDTO.UserName,
		"Email": userDTO.Email,
		"Role":  userDTO.Role,
	}).Info("create user input")
	return nil
}

func (u *userService) UpdateUser(id string, userDTO dto.UserDTO) (*model.User, error) {
	existingUser, err := u.repo.GetUserById(id)
	if err != nil || existingUser == nil {
		return nil, errors.New("user not found")
	}
	var user = mapper.MapUserDtoToUser(userDTO, *existingUser)
	userUpdate, errUpdate := u.repo.UpdateUser(&user)
	if errUpdate != nil {
		return nil, errUpdate
	}
	log.WithFields(log.Fields{
		"Id":    id,
		"Name":  userDTO.UserName,
		"Email": userDTO.Email,
		"Role":  userDTO.Role,
	}).Info("create user input")
	return userUpdate, nil
}

func (u *userService) DeleteUser(id string) error {
	existingUser, err := u.repo.GetUserById(id)
	if err != nil || existingUser == nil {
		return errors.New("user not found")
	}
	err = u.repo.DeleteUser(existingUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetById(id string) (*model.User, error) {
	return u.repo.GetUserById(id)
}

func (u *userService) GetByUserName(username string) (*model.User, error) {
	return u.repo.GetUserByUserName(username)
}

func (u *userService) CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (u *userService) Authenticate(username, password string) (*model.User, error) {
	// Lấy người dùng từ cơ sở dữ liệu dựa trên username
	user, err := u.repo.GetUserByUserName(username)
	if err != nil {
		return nil, err
	}

	// Kiểm tra xem password có khớp không
	if !ComparePasswords(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
