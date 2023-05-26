package db

import (
	"database/sql"
	"time"
)

type Logs struct {
	Db *sql.DB
}

func (log *Logs) CreateLog(DNI int, score int, wallet_id int) error {
	createDate := time.Now()
	dateString := createDate.Format("2006-01-02")
	var id int
	err := log.Db.QueryRow("INSERT INTO Logs(DNI, petition_date, status, wallet_id) values ($1, $2, $3, $4) RETURNING id", DNI, dateString, score, wallet_id).Scan(&id)

	if err != nil {
		return err
	}
	return err
}
