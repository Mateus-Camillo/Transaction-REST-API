package controller 

import (
	"net/http"

	"fmt"

	"github.com/Mateus-Camillo/Transaction-REST-API/model"
	"github.com/Mateus-Camillo/Transaction-REST-API/usecase"
)

type transferController struct {
	transferUseCase usecase.TransferUseCase
}

func NewTransferController(usecase usecase.TransferUseCase) transferController {
	return transferController{
		transferUseCase: usecase,
	}
}

func (t *transferController) TransferAmount(ctx *gin.Context) {
	var transfer model.Transfer
	err := ctx.BindJSON(&transfer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	} 

	err := t.transferUseCase.TransferBalance(transfer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println("Second Transfer Controller Error, develop error possiblities")
		return
	}

	response := model.Response{
		Message: "Money transfer done successfully"
	}

	ctx.JSON(http.StatusOK, response)
}