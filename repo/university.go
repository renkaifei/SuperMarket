package repo

import (
	"database/sql"
	"errors"
)

type University struct {
	UniversityId      int    `json:"universityId"`
	UniversityName    string `json:"universityName"`
	UniversityCode    string `json:"universityCode"`
	UniversityAddress string `json:"universityAddress"`
	UniversityPicture string `json:"universityPicture"`
}

func (a *University) Create() error {
	var (
		universityId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow("select UniversityId from University where UniversityId <> ? and UniversityName = ? ", a.UniversityId, a.UniversityName)
	err = row.Scan(&universityId)
	if err != sql.ErrNoRows {
		return errors.New("高校[" + a.UniversityName + "]已经存。")
	}
	result, err := tx.Exec(" insert into University(UniversityName,UniversityCode,UniversityAddress,UniversityPicture)values(?,?,?,?) ", a.UniversityName, a.UniversityCode, a.UniversityAddress, a.UniversityPicture)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	universityId, err = result.LastInsertId()
	if err != nil {
		return nil
	}
	a.UniversityId = int(universityId)
	return nil
}
