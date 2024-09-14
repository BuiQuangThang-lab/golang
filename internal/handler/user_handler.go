package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qlnv/internal/dto"
	filter "qlnv/internal/dto/filter"
	"qlnv/internal/model"
	"qlnv/internal/service"
)

type UserHandler struct {
	service      service.UserService
	basicHandler *BasicHandler
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) GetListUsers(c *gin.Context) {
	var search filter.UserFilter
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "des": "Invalid query parameters", "data": nil})
		return
	}
	if search.Page < 1 {
		search.Page = 1
	}
	if search.PageSize < 1 {
		search.PageSize = 10
	}
	users, total, err := u.service.GetListUser(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Status: http.StatusInternalServerError,
			Des:    "Failed to get user",
		})
		return
	}
	c.JSON(http.StatusOK, model.APIResponse{
		Status: http.StatusOK,
		Des:    "Get list successfully",
		Data:   users,
		Page:   search.Page,
		Size:   search.PageSize,
		Total:  total,
	})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Status: http.StatusBadRequest,
			Des:    "Invalid request payload",
		})
		return
	}

	err := u.service.CreateUser(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Status: http.StatusInternalServerError,
			Des:    "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Status: http.StatusOK,
		Des:    "User created successfully",
		Data:   nil,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userDto dto.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Status: 400,
			Des:    "Invalid request payload",
		})
		return
	}

	// Gọi service để update user
	_, err := h.service.UpdateUser(id, userDto)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, model.APIResponse{
				Status: 404,
				Des:    err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, model.APIResponse{
				Status: 500,
				Des:    "Failed to update user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Status: 200,
		Des:    "User updated successfully",
		Data:   nil,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, model.APIResponse{
				Status: 400,
				Des:    "Invalid request payload",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, model.APIResponse{
				Status: 500,
				Des:    "Failed to delete user",
			})
		}
	}
	c.JSON(http.StatusOK, model.APIResponse{
		Status: 200,
		Des:    "User deleted successfully",
		Data:   nil,
	})
}
