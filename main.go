package main

import (
	"fmt"
	"loan_tracker_api/delivery/controller/user_controller"
	"loan_tracker_api/delivery/router"
	"loan_tracker_api/infrastructure/auth"
	"loan_tracker_api/infrastructure/bootstrap"
	"loan_tracker_api/infrastructure/email"
	"loan_tracker_api/repository/refresh_token_repository"
	"loan_tracker_api/repository/reset_token_repository"
	"loan_tracker_api/repository/user_repository"
	"loan_tracker_api/usecase/user_usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	userCollection := db.Collection("users")
	refreshTokenCollection := db.Collection("refresh-tokens")
	resetTokenCollection := db.Collection("reset-tokens")
	userRepo := user_repository.NewUserRepository(userCollection)
	refreshTokenRepo := refresh_token_repository.NewRefreshTokenRepository(refreshTokenCollection)
	resetTokenRepo := reset_token_repository.NewResetTokenRepository(resetTokenCollection)

	authService := auth.NewAuthService(refreshTokenRepo, resetTokenRepo, env.AccessTokenSecret, env.RefreshTokenSecret, env.ResetTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour, env.ResetTokenExpiryHour)

	emailService := email.NewEmailService(env.SMTPServer, env.SMTPPort, env.SMTPUser, env.SMTPPassword, env.FromAddress)

	userUsecase := user_usecase.NewUserUsecase(userRepo, authService, emailService, time.Duration(env.ContextTimeout))
	userController := user_controller.NewUserController(userUsecase, authService, env)

	r := gin.Default()
	router.SetRouter(r, userController, env)
	r.Run(env.ServerAddress)

	fmt.Println("Work correctly", env.DBName)
}
