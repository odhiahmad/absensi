package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/absensi/config"
	"github.com/odhiahmad/absensi/controller"
	"github.com/odhiahmad/absensi/repository"
	"github.com/odhiahmad/absensi/service"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	absenRepository  repository.AbsenRepository  = repository.NewAbsenRepository(db)


	jwtService   service.JWTService   = service.NewJwtService()
	authService  service.AuthService  = service.NewAuthService(userRepository)
	userService  service.UserService  = service.NewUserService(userRepository)
	absenService  service.AbsenService  = service.NewAbsenService(absenRepository)


	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	absenController  controller.AbsenController  = controller.NewAbsenController(absenService, jwtService)

)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Static("/image", "./fileupload")

	r.Use(CORSMiddleware())

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
	// middleware.AuthorizeJWT(jwtService)

	userRoutes := r.Group("api/user")
	{
		userRoutes.POST("/create", userController.CreateUser)
		userRoutes.PUT("/update", userController.UpdateUser)
	}

	absenRoutes := r.Group("api/absen")
	{
		absenRoutes.POST("/create", userController.CreateUser)
		absenRoutes.PUT("/update", userController.UpdateUser)
	}

	r.Run()
}
