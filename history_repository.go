package repository

import (
	"database/sql"
	"fmt"

	"example.com/m/entity"
)

type HistoryRepo interface {
	FetchHistory() (*[]entity.History, error)
	DeleteHistory(HistoryId int) error
	AddHistory(history *entity.History) error
	FindHistory(historyId int) (bool, error)
}

type historyRepoImpl struct {
	db *sql.DB
}

type HistoryRepositoryConfig struct {
	DB *sql.DB
}

func NewHistoryRepository(c HistoryRepositoryConfig) HistoryRepo {
	return &historyRepoImpl{
		db: c.DB,
	}
}

func (r *historyRepoImpl) FindHistory(historyId int) (bool, error) {

	var country string
	var value int64

	findHistory := `SELECT country, value FROM "histories" where history_id=$1`
	row := r.db.QueryRow(findHistory, historyId)
	err := row.Scan(&country, &value)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}

func (r *historyRepoImpl) FetchHistory() (*[]entity.History, error) {
	var histories []entity.History

	rows, err := r.db.Query(`SELECT "history_id", "country", "value" FROM "histories"`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		history := new(entity.History)
		err = rows.Scan(&history.HistoryId, &history.Country, &history.Value)
		histories = append(histories, *history)

		if err != nil {
			return nil, err
		}
	}

	return &histories, nil
}

func (r *historyRepoImpl) DeleteHistory(HistoryId int) error {

	deleteHistory := `delete from "histories" where history_id=$1`
	_, err := r.db.Exec(deleteHistory, HistoryId)

	if err != nil {
		return err
	}

	return nil
}

func (r *historyRepoImpl) AddHistory(history *entity.History) error {

	insertHistory := `insert into "histories" ("country", "value") values($1, $2)`
	_, err := r.db.Exec(insertHistory, history.Country, history.Value)

	if err != nil {
		return err
	}

	return nil
}
