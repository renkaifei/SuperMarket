package repo

import (
	"database/sql"
	"errors"
)

type Goods struct {
	GoodsId    int
	GoodsCode  string
	GoodsName  string
	CategoryId int
}

func NewGoods(goodsId int, goodsCode string, goodsName string, categoryId int) *Goods {
	return &Goods{GoodsId: goodsId, GoodsCode: goodsCode, GoodsName: goodsName, CategoryId: categoryId}
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
	result, err := tx.Exec(" insert into goods(goodsCode,goodsName,goodsCategoryId)values(?,?,?) ", a.GoodsCode, a.GoodsName, a.CategoryId)
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
	_, err = tx.Exec(" update Goods set GoodsCode = ?,GoodsName = ? where GoodsId = ?  ", a.GoodsCode, a.GoodsName, a.GoodsId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
