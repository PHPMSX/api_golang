package main

import (
	"golang_api/config"
	"golang_api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnect()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnect(db)
	r := gin.Default()
	authRouters := r.Group("api/auth")
	{
		authRouters.POST("/login", authController.Login)
		authRouters.POST("/register", authController.Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}