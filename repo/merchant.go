package repo

import (
	"database/sql"
	"errors"
)

type Merchant struct {
	MerchantId          int
	MerchantName        string
	MerchantAlias       string
	MerchantAddress     string
	MerchantCEO         string
	MerchantMobilePhone string
}

func (a *Merchant) SelectById() error {
	row := mySqlDB.QueryRow(" select merchantId,merchantName,merchantAddress,ceo,mobilePhone,merchantAlias from merchant where merchantId = ? ", a.MerchantId)
	err := row.Scan(&a.MerchantId, &a.MerchantName, &a.MerchantAddress, &a.MerchantCEO, &a.MerchantMobilePhone, &a.MerchantAlias)
	if err != sql.ErrNoRows {
		errors.New("没有该商户的相关信息")
	}
	return nil
}
