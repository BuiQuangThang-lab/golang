package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"qlnv/internal/model"
	"strings"
)

type BasicHandler struct {
	secretKey string
}

func NewBasicHandler(secretKey string) *BasicHandler {
	return &BasicHandler{secretKey: secretKey}
}

// AuthMiddleware kiểm tra token để bảo vệ API
func (h *BasicHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusBadRequest, model.APIResponse{
				Status: http.StatusUnauthorized,
				Des:    "Missing or invalid token",
			})
			c.Abort()
			return
		}

		// Xử lý xác thực token
		if !h.ValidateToken(tokenStr) {
			c.JSON(http.StatusBadRequest, model.APIResponse{
				Status: http.StatusUnauthorized,
				Des:    "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (h *BasicHandler) ValidateToken(tokenStr string) bool {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(h.secretKey), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}
