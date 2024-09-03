package controller

import (
	"net/http"

	"errors"

	"github.com/Mateus-Camillo/Transaction-REST-API/model"
	"github.com/Mateus-Camillo/Transaction-REST-API/usecase"
	"github.com/gin-gonic/gin"
)

type accountController struct {
	accountUseCase usecase.AccountUseCase
}

func NewAccountController(usecase usecase.AccountUseCase) accountController {
	return accountController{
		accountUseCase: usecase,
	}
}

func (a *accountController) CreateAccount(ctx *gin.Context) {
	var account model.Account
	err := ctx.BindJSON(&account)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdAccount, err := a.accountUseCase.CreateAccount(account)
	if err != nil {
		if errors.Is(err, model.ErrInvalidPassword) {
			response := model.Response{
				Message: "Password must be, at least, 8 characters long",
			}

			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdAccount)
}
