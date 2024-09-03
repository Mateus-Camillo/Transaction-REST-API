package main

import (
	"github.com/Mateus-Camillo/Transaction-REST-API/controller"
	"github.com/Mateus-Camillo/Transaction-REST-API/db"
	"github.com/Mateus-Camillo/Transaction-REST-API/repository"
	"github.com/Mateus-Camillo/Transaction-REST-API/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	AccountRepository := repository.NewAccountRepository(dbConnection)
	AccountUseCase := usecase.NewAccountUseCase(AccountRepository)
	AccountController := controller.NewAccountController(AccountUseCase)

	TransferRepository := repository.NewTransferRepository(dbConnection)
	TransferUseCase := usecase.NewTransferUseCase(TransferRepository)
	TransferController := controller.NewTransferController(TransferUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/account", AccountController.CreateAccount)

	server.POST("/transfer", TransferController.TransferAmount)

	server.Run(":8000")
}
