package handler

import (
	"fmt"
	"net/http"

	"example.com/m/dto"
	"example.com/m/errors"
	"example.com/m/usecase"
	"github.com/gin-gonic/gin"
)

type HandlerHistoryImpl struct {
	historyUsecase usecase.HistoryUsecase
}

type HandlerHistoryConfig struct {
	HistoryUsecase usecase.HistoryUsecase
}

func NewHistoryHandler(c HandlerHistoryConfig) *HandlerHistoryImpl {
	return &HandlerHistoryImpl{
		historyUsecase: c.HistoryUsecase,
	}
}

func (h *HandlerHistoryImpl) GetCurrencyHistory(c *gin.Context) {

	history, err := h.historyUsecase.GetHistory()

	if err != nil {
		serverErr := errors.ErrInternalServer
		c.AbortWithStatusJSON(500, serverErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": history,
	})
}

func (h *HandlerHistoryImpl) DeleteCurrencyHistory(c *gin.Context) {

	var reqDelete dto.DeleteRequest

	if err := c.ShouldBindJSON(&reqDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	err := h.historyUsecase.DeleteHistory(int(reqDelete.NumberId))

	fmt.Println("Handler")
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "delete success",
	})
}

func (h *HandlerHistoryImpl) AddHistory(c *gin.Context) {
	var reqInput dto.InputHistory

	if err := c.ShouldBindJSON(&reqInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
	}

	err := h.historyUsecase.AddHistory(reqInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success added entity",
	})
}
