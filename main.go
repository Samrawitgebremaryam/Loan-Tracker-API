package main

import (
	"fmt"
	"loan_tracker_api/delivery/controller/loan_controller"
	"loan_tracker_api/delivery/controller/user_controller"
	"loan_tracker_api/delivery/router"
	"loan_tracker_api/infrastructure/auth"
	"loan_tracker_api/infrastructure/bootstrap"
	"loan_tracker_api/infrastructure/email"
	"loan_tracker_api/repository/loan_repository"
	"loan_tracker_api/repository/refresh_token_repository"
	"loan_tracker_api/repository/reset_token_repository"
	"loan_tracker_api/repository/user_repository"
	"loan_tracker_api/usecase/loan_usecase"
	"loan_tracker_api/usecase/user_usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	// Initialize Repositories
	userCollection := db.Collection("users")
	refreshTokenCollection := db.Collection("refresh-tokens")
	resetTokenCollection := db.Collection("reset-tokens")
	loanCollection := db.Collection("loans")

	userRepo := user_repository.NewUserRepository(userCollection)
	refreshTokenRepo := refresh_token_repository.NewRefreshTokenRepository(refreshTokenCollection)
	resetTokenRepo := reset_token_repository.NewResetTokenRepository(resetTokenCollection)
	loanRepo := loan_repository.NewLoanRepository(loanCollection)

	// Initialize Services
	authService := auth.NewAuthService(refreshTokenRepo, resetTokenRepo, env.AccessTokenSecret, env.RefreshTokenSecret, env.ResetTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour, env.ResetTokenExpiryHour)
	emailService := email.NewEmailService(env.SMTPServer, env.SMTPPort, env.SMTPUser, env.SMTPPassword, env.FromAddress)

	// Initialize Usecases
	userUsecase := user_usecase.NewUserUsecase(userRepo, authService, emailService, time.Duration(env.ContextTimeout))
	loanUsecase := loan_usecase.NewLoanUsecase(loanRepo, time.Duration(env.ContextTimeout))

	// Initialize Controllers
	userController := user_controller.NewUserController(userUsecase, authService, env)
	loanController := loan_controller.NewLoanController(loanUsecase)

	// Setup Router
	r := gin.Default()
	router.SetRouter(r, userController, loanController, env)
	r.Run(env.ServerAddress)

	fmt.Println("Server running on", env.ServerAddress)
}
