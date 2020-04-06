package repo

import (
	"database/sql"
	"errors"
)

type MerchantGoods struct {
	MerchantGoodsId int
	MerchantId      int
	GoodsId         int
	GoodsName       string
	Price           float64
	Discount        float64
}

func (a *MerchantGoods) Create() error {
	var (
		merchantGoodsId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow("select MerchantGoodsId from MerchantGoods where merchantId = ? and GoodsId = ? ", a.MerchantId, a.GoodsId)
	err = row.Scan(&merchantGoodsId)
	if err != sql.ErrNoRows {
		return errors.New("商品已经存在")
	}
	result, err := tx.Exec("insert into MerchantGoods(merchantId,goodsId,price,discount)values(?,?,?,?)", a.MerchantId, a.GoodsId, a.Price, a.Discount)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	merchantGoodsId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	a.MerchantGoodsId = int(merchantGoodsId)
	return nil
}

func (a *MerchantGoods) Update() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("update MerchantGoods set Price = ?,Discount = ? where MerchantGoodsId = ? ", a.Price, a.Discount, a.MerchantGoodsId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *MerchantGoods) Delete() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("delete from MerchantGoods where MerchantGoodsId = ？", a.MerchantGoodsId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *MerchantGoods) SelectById() error {
	row := mySqlDB.QueryRow("select MerchantId,GoodsId,Price,Discount from MerchantGoods where MerchantGoodsId = ? ", a.MerchantGoodsId)
	err := row.Scan(&a.MerchantId, &a.GoodsId, &a.Price, &a.Discount)
	if err != nil {
		return err
	}
	return nil
}

func (a *MerchantGoods) SelectByMerchantIdAndGoodsId() error {
	row := mySqlDB.QueryRow("select MerchantGoodsId,MerchantId,GoodsId,Price,Discount from MerchantGoods where MerchantId = ? and GoodsId = ? ", a.MerchantId, a.GoodsId)
	err := row.Scan(&a.MerchantGoodsId, &a.MerchantId, &a.GoodsId, &a.Price, &a.Discount)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

type MerchantGoodses struct {
	Values     []*MerchantGoods
	PageIndex  int
	PageSize   int
	TotalCount int
}

func (a *MerchantGoodses) SelectPageByMerchantId(content string, merchantId int, pageIndex int, pageSize int) error {
	a.PageIndex = pageIndex
	a.PageSize = pageSize
	row := mySqlDB.QueryRow(" select count(*) totalCount from MerchantGoods join goods on merchantGoods.GoodsId = goods.GoodsId where goods.GoodsName like ? and MerchantId = ? ", "%"+content+"%", merchantId)
	err := row.Scan(&a.TotalCount)
	if err != nil {
		return nil
	}
	rows, err := mySqlDB.Query("select MerchantGoodsId,MerchantId,Goods.GoodsId,Goods.GoodsName, Price,Discount from MerchantGoods join Goods on MerchantGoods.GoodsId = Goods.goodsId where Goods.GoodsName like ? and  merchantId = ? limit ?,? ", "%"+content+"%", merchantId, (pageIndex-1)*pageSize, pageSize)
	if err != nil {
		return nil
	}
	defer rows.Close()
	a.Values = make([]*MerchantGoods, 0)
	for rows.Next() {
		item := &MerchantGoods{}
		rows.Scan(&item.MerchantGoodsId, &item.MerchantId, &item.GoodsId, &item.GoodsName, &item.Price, &item.Discount)
		a.Values = append(a.Values, item)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
