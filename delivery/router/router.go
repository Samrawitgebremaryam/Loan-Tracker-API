package router

import (
	"loan_tracker_api/delivery/controller/loan_controller"
	"loan_tracker_api/delivery/controller/user_controller"
	"loan_tracker_api/infrastructure/auth"
	"loan_tracker_api/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, lc *loan_controller.LoanController, env *bootstrap.Env) {
	authMiddleware := auth.JwtAuthMiddleware(env.AccessTokenSecret)

	// User routes (protected by JWT middleware)
	userRoutes := router.Group("/")
	userRoutes.Use(authMiddleware)
	{
		userRoutes.POST("/signup", uc.SignUp)
		userRoutes.POST("/login", uc.Login)
		userRoutes.POST("/refresh", uc.RefreshTokens)
		userRoutes.POST("/forgot-password", uc.ForgotPassword)
		userRoutes.POST("/reset-password", uc.ResetPassword)
		userRoutes.GET("/users/profile", uc.GetProfile)
	}

	// Loan routes (protected by JWT middleware)
	loanRoutes := router.Group("/loans")
	loanRoutes.Use(authMiddleware)
	{
		loanRoutes.POST("/", lc.ApplyForLoan)
		loanRoutes.GET("/:id", lc.ViewLoanStatus)
	}

	// Admin routes (protected by JWT and Admin middleware)
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(auth.AdminMiddleware())
	{
		adminRoutes.GET("/loans", lc.ViewAllLoans)
		adminRoutes.PATCH("/loans/:id/status", lc.UpdateLoanStatus)
		adminRoutes.DELETE("/loans/:id", lc.DeleteLoan)
		adminRoutes.GET("/users", uc.GetUsers)
		adminRoutes.DELETE("/users/:id", uc.DeleteUser)
	}
}
