package handler

import (
	"fmt"
	"net/http"

	"example.com/m/dto"
	"example.com/m/usecase"
	"github.com/gin-gonic/gin"
)

type HandlerTransactionImpl struct {
	transactionUsecase usecase.TransactionUsecase
}

type HandlerTransactionConfig struct {
	TransactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(c HandlerTransactionConfig) *HandlerTransactionImpl {
	return &HandlerTransactionImpl{
		transactionUsecase: c.TransactionUsecase,
	}
}

func (h *HandlerTransactionImpl) FetchMachineUtil(c *gin.Context) {
	var reqTransaction dto.MachineTransactionRequest

	if err := c.ShouldBindJSON(&reqTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	fmt.Println("Hello There")

	fmt.Println(reqTransaction)

	transactionResult, err := h.transactionUsecase.FetchMachineUtil(reqTransaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResult,
	})

}
