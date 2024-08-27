package main

import (
	"fmt"
	"loan_tracker_api/infrastructure/bootstrap"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	fmt.Println("Connected to MongoDB!")
	fmt.Println(db)
	// userCollection := db.Collection("users")
	// userRepository := repository.NewUserRepository(userCollection)

}
