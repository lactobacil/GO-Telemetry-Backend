package usecase

import (
	"fmt"

	"example.com/m/dto"
	"example.com/m/entity"
	"example.com/m/repository"
)

type HistoryUsecase interface {
	GetHistory() (*[]entity.History, error)
	DeleteHistory(historyId int) error
	AddHistory(newHistory dto.InputHistory) error
}

type historyUsecaseImpl struct {
	historyRepo repository.HistoryRepo
}

type HistoryUsecaseConfig struct {
	HistoryRepo repository.HistoryRepo
}

func NewHistoryUsecase(h HistoryUsecaseConfig) HistoryUsecase {
	return &historyUsecaseImpl{
		historyRepo: h.HistoryRepo,
	}
}

func (h *historyUsecaseImpl) GetHistory() (*[]entity.History, error) {
	history, err := h.historyRepo.FetchHistory()

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (h *historyUsecaseImpl) DeleteHistory(historyId int) error {

	fmt.Println(historyId)
	fmt.Println("Usecase")

	isHistoryId, err := h.historyRepo.FindHistory(historyId)

	if !isHistoryId {
		return err
	}

	err = h.historyRepo.DeleteHistory(historyId)

	if err != nil {
		return err
	}

	return nil
}

func (h *historyUsecaseImpl) AddHistory(data dto.InputHistory) error {

	historyInput := &entity.History{
		Country: data.Country,
		Value:   int(data.Value),
	}

	err := h.historyRepo.AddHistory(historyInput)

	if err != nil {
		return err
	}

	return nil
}
