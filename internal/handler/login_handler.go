package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"qlnv/internal/dto"
	"qlnv/internal/model"
	"qlnv/internal/service"
	"time"
)

type LoginHandler struct {
	secretKey string
	service   service.UserService
}

func NewLoginHandler(secretKey string, service *service.UserService) *LoginHandler {
	return &LoginHandler{secretKey: secretKey, service: *service}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "des": "Invalid request payload"})
		return
	}

	user, err := h.service.Authenticate(loginDTO.Username, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "des": "Invalid credentials"})
		return
	}

	// Sinh JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token hết hạn sau 1 giờ
	})

	tokenString, err := token.SignedString([]byte(h.secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Status: 500,
			Des:    "Could not create token",
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Status: 200,
		Des:    "Login successful",
		Data:   map[string]string{"token": tokenString},
	})
}
