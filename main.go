package main

import (
	"demo-backend/route"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	user_handler "demo-backend/service/user/handler"
	user_repository "demo-backend/service/user/repository"
	user_usecase "demo-backend/service/user/usecase"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=pheet1234 dbname=pheet_db_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	/* Init Repository */
	userRepo := user_repository.NewUserRepository(db)

	/* Init Usecase */
	userUs := user_usecase.NewUserUsecase(userRepo)

	/* Init Usecase */
	userHandler := user_handler.NewUserHandler(userUs)

	/* web server gofiber, echo, gin*/
	app := gin.Default()
	app.GET("/v1/users", userHandler.FetchUsers)

	/* Route */
	route := route.NewRoute(app)
	route.RegisterUser(userHandler)

	/* Start Web Server */
	app.Run(":8080")
}
