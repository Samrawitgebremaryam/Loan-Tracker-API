package router

import (
	"loan_tracker_api/delivery/controller/user_controller"
	"loan_tracker_api/infrastructure/auth"
	"loan_tracker_api/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, env *bootstrap.Env) {
	// User routes
	router.POST("/signup", uc.SignUp)
	router.POST("/login", uc.Login)
	router.POST("/refresh", uc.RefreshTokens)
	router.POST("/forgot-password", uc.ForgotPassword)
	router.POST("/reset-password", uc.ResetPassword)

	authMiddleware := auth.JwtAuthMiddleware(env.AccessTokenSecret)
	userRoutes := router.Group("/users")

	userRoutes.Use(authMiddleware) // Apply JWT middleware to all /users route
	{
		userRoutes.GET("/profile", uc.GetProfile)
	}
	// Admin
	adminRoutes := userRoutes.Group("/admin")
	adminRoutes.Use(auth.AdminMiddleware())
	{
		adminRoutes.Use(authMiddleware) // Apply JWT middleware to all /admin routes
		adminRoutes.GET("/users", uc.GetUsers)
		adminRoutes.DELETE("/users/:id", uc.DeleteUser)
	}
}
