package repository

import (
	"errors"
	"gorm.io/gorm"
	"qlnv/internal/dto/filter"
	"qlnv/internal/model"
)

type UserRepository interface {
	GetListUser(userFilter filter.UserFilter) ([]model.User, int64, error)
	GetUserById(id string) (*model.User, error)
	GetUserByUserName(username string) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetListUser(filter filter.UserFilter) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	query := u.db.Model(&model.User{})
	if filter.Name != "" {
		query = u.db.Where("username LIKE ?", "%"+filter.Name+"%")
	}
	query.Count(&total)
	offset := (filter.Page - 1) * filter.PageSize
	result := query.Limit(filter.PageSize).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, total, nil
}

func (u *userRepository) GetUserById(id string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUserByUserName(username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) CreateUser(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	return user, u.db.Updates(user).Error
}

func (u *userRepository) DeleteUser(user *model.User) error {
	return u.db.Delete(user).Error
}
