package repo

import (
	"database/sql"
	"errors"
)

type Goods struct {
	GoodsId      int
	GoodsCode    string
	GoodsName    string
	CategoryId   int
	GoodsBarCode string
}

func NewGoods(goodsId int, goodsCode string, goodsName string, categoryId int, goodsBarCode string) *Goods {
	return &Goods{GoodsId: goodsId, GoodsCode: goodsCode, GoodsName: goodsName, CategoryId: categoryId, GoodsBarCode: goodsBarCode}
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
	row := tx.QueryRow(" select goodsId from goods where goodsCode = ? ", a.GoodsCode)
	err = row.Scan(&goodsId)
	if err != sql.ErrNoRows {
		return errors.New("商品编码[" + a.GoodsCode + "]已经存在")
	}
	result, err := tx.Exec(" insert into goods(goodsCode,goodsName,goodsCategoryId,goodsBarCode)values(?,?,?,?) ", a.GoodsCode, a.GoodsName, a.CategoryId, a.GoodsBarCode)
	if err != nil {
		return err
	}
	goodsId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
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
	row := tx.QueryRow("select goodsId from Goods where goodsCode = ?", a.GoodsCode)
	err = row.Scan(&goodsId)
	if err != sql.ErrNoRows {
		return errors.New("商品编码[" + a.GoodsCode + "]已经存在")
	}
	_, err = tx.Exec(" update Goods set GoodsCode = ?,GoodsName = ?,GoodsBarCode = ? where GoodsId = ?  ", a.GoodsCode, a.GoodsName, a.GoodsBarCode, a.GoodsId)
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
	row := mySqlDB.QueryRow(" select GoodsId,GoodsCode,GoodsName,GoodsCategoryId,GoodsBarCode from Goods where GoodsId = ? ", a.GoodsId)
	err := row.Scan(&a.GoodsId, &a.GoodsCode, &a.GoodsName, &a.CategoryId, &a.GoodsBarCode)
	if err != nil {
		return err
	}
	return nil
}
