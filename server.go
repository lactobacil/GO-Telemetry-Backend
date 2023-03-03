package server

import (
	"database/sql"
	"fmt"

	"example.com/m/config"
	"example.com/m/repository"
	"example.com/m/usecase"
	"github.com/gin-gonic/gin"
)

func initRouter(db *sql.DB) *gin.Engine {

	historyRepoConfig := repository.HistoryRepositoryConfig{DB: db}
	historyRepo := repository.NewHistoryRepository(historyRepoConfig)

	historyUsecaseConfig := usecase.HistoryUsecaseConfig{HistoryRepo: historyRepo}
	historyUsecase := usecase.NewHistoryUsecase(historyUsecaseConfig)

	calendarRepoConfig := repository.CalendarRepositoryConfig{DB: db}
	calendarRepo := repository.NewCalendarRepository(calendarRepoConfig)

	calendarUsecaseConfig := usecase.CalendarUsecaseConfig{CalendarRepo: calendarRepo}
	calendarUsecase := usecase.NewCalendarUsecase(calendarUsecaseConfig)

	transactionRepoConfig := repository.TransactionRepositoryConfig{DB: db}
	transactionRepo := repository.NewTransactionRepository(transactionRepoConfig)

	transactionUsecaseConfig := usecase.TransactionUsecaseConfig{TransactionRepo: transactionRepo}
	transactionUsecase := usecase.NewTransactionUsecase(transactionUsecaseConfig)

	r := NewRouter(RouterConfig{
		historyUsecase:     historyUsecase,
		calendarUsecase:    calendarUsecase,
		transactionUsecase: transactionUsecase,
	})

	return r
}

func Init() {

	fmt.Println("Text")

	db := config.ConnectDB()
	r := initRouter(db)

	err := r.Run()
	if err != nil {
		fmt.Println("Error while running server", err)
		return
	}
}
