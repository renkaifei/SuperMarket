package repo

type GoodsCategory struct {
	CategoryId   int
	CategoryName string
	ParentId     int
}

func (a *GoodsCategory) Create() error {
	var (
		categoryId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	result, err := tx.Exec(" insert into goodsCategory(categoryName,parentId)values(?,?) ", a.CategoryName, a.ParentId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	categoryId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	a.CategoryId = int(categoryId)
	return nil
}

func (a *GoodsCategory) Update() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("update goodsCategory set categoryName = ? , parentId = ? where categoryId = ? ", a.CategoryName, a.ParentId, a.CategoryId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *GoodsCategory) Delete() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	_, err = tx.Exec(" delete from goodsCategory where categoryId = ? ", a.CategoryId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *GoodsCategory) selectById() error {
	row := mySqlDB.QueryRow(" select categoryId,categoryName,parentId from goodsCategory where categoryId = ? ", a.CategoryId)
	err := row.Scan(&a.CategoryId, &a.CategoryName, &a.ParentId)
	return err
}

type GoodsCategorys struct {
}

func (a *GoodsCategorys) SelectByName(name string) (ret []*GoodsCategory, err error) {
	rows, err := mySqlDB.Query(" select categoryId,categoryName,parentId from goodsCategory where categoryName like '%" + name + "%'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categorys := make([]*GoodsCategory, 0)
	for rows.Next() {
		category := &GoodsCategory{}
		if err = rows.Scan(&category.CategoryId, &category.CategoryName, &category.ParentId); err != nil {
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
