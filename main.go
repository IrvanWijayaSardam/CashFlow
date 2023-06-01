package main

import (
	"github.com/IrvanWijayaSardam/CashFlow/config"
	"github.com/IrvanWijayaSardam/CashFlow/controller"
	"github.com/IrvanWijayaSardam/CashFlow/middleware"
	"github.com/IrvanWijayaSardam/CashFlow/repository"
	"github.com/IrvanWijayaSardam/CashFlow/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                         = config.SetupDatabaseConnection()
	userRepository repository.UserRepository        = repository.NewUserRepository(db)
	trxRepository  repository.TransactionRepository = repository.NewTransactionRepository(db)
	jwtService     service.JWTService               = service.NewJWTService()
	userService    service.UserService              = service.NewUserService(userRepository)
	authService    service.AuthService              = service.NewAuthService(userRepository)
	authController controller.AuthController        = controller.NewAuthController(authService, jwtService)
	userController controller.UserController        = controller.NewUserController(userService, jwtService)
	trxSertvice    service.TransactionService       = service.NewTransactionService(trxRepository)
	trxController  controller.TransactionContoller  = controller.NewTransactionController(trxSertvice, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	trxRoutes := r.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		trxRoutes.GET("/", trxController.All)
		trxRoutes.POST("/", trxController.Insert)
		trxRoutes.PUT("/", trxController.Update)
		trxRoutes.DELETE("/:id", trxController.Delete)
	}

	reportRoutes := r.Group("api/report", middleware.AuthorizeJWT(jwtService))
	{
		reportRoutes.GET("/", trxController.SumGroupId)
		reportRoutes.GET("/summary", trxController.TransactionReport)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.POST("/picture", userController.SaveFile)
	}

	cdnRoutes := r.Group("api/cdn")
	{
		cdnRoutes.GET("/picture/:file_name", userController.GetFile)
	}
	// userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	// {
	// 	userRoutes.GET("/profile", userController.Profile)
	// 	userRoutes.PUT("/profile", userController.Update)
	// }

	// noteRoutes := r.Group("api/notes", middleware.AuthorizeJWT(jwtService))
	// {
	// 	noteRoutes.GET("/", noteController.All)
	// 	noteRoutes.POST("/", noteController.Insert)
	// 	noteRoutes.GET("/:id", noteController.FindById)
	// 	noteRoutes.PUT("/:id", noteController.Update)
	// 	noteRoutes.DELETE("/:id", noteController.Delete)
	// }

	// pagerRoutes := r.Group("pager", middleware.AuthorizeJWT(jwtService))
	// {
	// 	pagerRoutes.GET("/", pagerController.All)
	// 	pagerRoutes.POST("/", pagerController.Insert)
	// 	pagerRoutes.GET("/:id", pagerController.FindById)
	// 	pagerRoutes.PUT("/:id", pagerController.Update)
	// 	pagerRoutes.DELETE("/:id", pagerController.Delete)
	// }

	// pagerStatusRoutes := r.Group("/api/status")
	// {
	// 	pagerStatusRoutes.GET("/:id", pagerController.FindStatusById)
	// }

	r.Run(":8001")
}
