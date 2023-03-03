package server

import (
	"net/http"

	"example.com/m/handler"
	"example.com/m/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	historyUsecase     usecase.HistoryUsecase
	calendarUsecase    usecase.CalendarUsecase
	transactionUsecase usecase.TransactionUsecase
}

func NewRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()

	handlerHistory := handler.NewHistoryHandler(handler.HandlerHistoryConfig{HistoryUsecase: c.historyUsecase})
	handlerCalendar := handler.NewCalendarHandler(handler.HandlerCalendarConfig{CalendarUsecase: c.calendarUsecase})
	handlerTransaction := handler.NewTransactionHandler(handler.HandlerTransactionConfig{TransactionUsecase: c.transactionUsecase})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	currency := r.Group("/currency")
	currency.GET("/history", handlerHistory.GetCurrencyHistory)
	currency.DELETE("/delete", handlerHistory.DeleteCurrencyHistory)
	currency.POST("/add", handlerHistory.AddHistory)

	calendar := r.Group("/calendar")
	calendar.POST("/add", handlerCalendar.AddNotes)
	calendar.POST("/notes", handlerCalendar.FetchNotes)
	calendar.POST("/day", handlerCalendar.FetchNotesDay)
	calendar.POST("/update", handlerCalendar.UpdateNotes)
	calendar.POST("/delete", handlerCalendar.DeleteNotes)

	machine := r.Group("/machine")
	machine.GET("/transaction", handlerTransaction.FetchMachineUtil)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"code":        "PAGE_NOT_FOUND",
			"message":     "Page not found",
		})
	})

	return r
}
