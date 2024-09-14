package presentation

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"qlnv/internal/handler"
	"qlnv/internal/model"
	"qlnv/internal/repository"
	"qlnv/internal/service"
)

func UserRouter(db *gorm.DB) *gin.Engine {
	model.Migrate(db)
	r := gin.Default()
	repo := repository.NewUserRepository(db)
	serv := service.NewUserService(repo)

	hdl := handler.NewUserHandler(serv)

	secretKey := "123456"

	loginHandler := handler.NewLoginHandler(secretKey, &serv)
	r.POST("/login", loginHandler.Login)

	basicHandler := handler.NewBasicHandler(secretKey)

	authorized := r.Group("/user")
	authorized.Use(basicHandler.AuthMiddleware())
	authorized.GET("", hdl.GetListUsers)
	authorized.POST("", hdl.CreateUser)
	authorized.PUT("/:id", hdl.UpdateUser)
	authorized.DELETE("/:id", hdl.DeleteUser)

	return r

}

func StartHTTPServer(db *gorm.DB) {
	r := UserRouter(db)
	r.Run(":8080")
}
