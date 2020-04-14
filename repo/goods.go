package repo

import (
	"database/sql"
	"errors"
)

type Goods struct {
	GoodsId            int    `json:"goodsId"`
	GoodsBarCode       string `json:"goodsBarCode"`
	GoodsName          string `json:"goodsName"`
	GoodsSpecification string `json:"goodsSpecification"`
	GoodsDescription   string `json:"goodsDescription"`
	GoodsTradeMark     string `json:"goodsTradeMark"`
	Company            string `json:"company"`
	GoodsPicture       string `json:"goodsPicture"`
}

func (a *Goods) Create() error {
	var (
		goodsId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(" select goodsId from goods where goodsBarCode = ? ", a.GoodsBarCode)
	err = row.Scan(&goodsId)
	if err != sql.ErrNoRows {
		return errors.New("商品[" + a.GoodsName + "]已经存在")
	}
	result, err := tx.Exec(" insert into goods(goodsBarCode,goodsName,goodsSpecification,goodsDescription,goodsTradeMark,company,goodsPicture)values(?,?,?,?,?,?,?) ", a.GoodsBarCode, a.GoodsName, a.GoodsSpecification, a.GoodsDescription, a.GoodsTradeMark, a.Company, a.GoodsPicture)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	goodsId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	a.GoodsId = int(goodsId)
	return nil
}

func (a *Goods) Update() error {
	var (
		goodsId int
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow("select goodsId from Goods where goodsBarCode = ? and goodsId <> ? ", a.GoodsBarCode, a.GoodsId)
	err = row.Scan(&goodsId)
	if err != sql.ErrNoRows {
		return errors.New("商品编码[" + a.GoodsName + "]已经存在")
	}
	_, err = tx.Exec(" update Goods set GoodsBarCode = ?,GoodsName = ?,GoodsSpecification = ?,GoodsDescription = ?,GoodsTradeMark = ?,Company = ?,GoodsPicture = ? where GoodsId = ?  ", a.GoodsBarCode, a.GoodsName, a.GoodsSpecification, a.GoodsDescription, a.GoodsTradeMark, a.Company, a.GoodsPicture, a.GoodsId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *Goods) Delete() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("delete from Goods where GoodsId = ? ", a.GoodsId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *Goods) SelectById() error {
	row := mySqlDB.QueryRow(" select GoodsId,GoodsBarCode,GoodsName,GoodsSpecification,GoodsDescription,GoodsTradeMark,Company,GoodsPicture from Goods where GoodsId = ? ", a.GoodsId)
	err := row.Scan(&a.GoodsId, &a.GoodsBarCode, &a.GoodsName, &a.GoodsSpecification, &a.GoodsDescription, &a.GoodsTradeMark, &a.Company, &a.GoodsPicture)
	if err != nil {
		return err
	}
	return nil
}

func (a *Goods) SelectByBarCode() error {
	row := mySqlDB.QueryRow("select GoodsId,GoodsBarCode,GoodsName,GoodsSpecification,GoodsDescription,GoodsTradeMark,Company,GoodsPicture from Goods where GoodsBarCode = ? ", a.GoodsBarCode)
	err := row.Scan(&a.GoodsId, &a.GoodsBarCode, &a.GoodsName, &a.GoodsSpecification, &a.GoodsDescription, &a.GoodsTradeMark, &a.Company, &a.GoodsPicture)
	return err
}

type Goodses struct {
	Values     []*Goods
	PageIndex  int
	PageSize   int
	TotalCount int
}

func (a *Goodses) SelectOnePage(content string, pageIndex int, pageSize int) error {
	row := mySqlDB.QueryRow("select count(*) totalCount from goods where goodsName like ? ", "%"+content+"%")
	err := row.Scan(&a.TotalCount)
	if err != nil {
		return err
	}
	rows, err := mySqlDB.Query(" select goodsId,goodsBarCode,goodsName,goodsSpecification,goodsDescription,goodsTradeMark,company,GoodsPicture from goods where goodsName like ? limit ?,? ", "%"+content+"%", (pageIndex-1)*pageSize, pageSize)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		goods := &Goods{}
		if err := rows.Scan(&goods.GoodsId, &goods.GoodsBarCode, &goods.GoodsName, &goods.GoodsSpecification, &goods.GoodsDescription, &goods.GoodsTradeMark, &goods.Company, &goods.GoodsPicture); err != nil {
			return err
		}
		a.Values = append(a.Values, goods)
	}
	a.PageIndex = pageIndex
	a.PageSize = pageSize
	return nil
}
