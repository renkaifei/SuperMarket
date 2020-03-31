package repo

import (
	"database/sql"
	"errors"
)

type Merchanter struct {
	MerchanterId     int
	MerchanterOpenId string
	MerchanterName   string
	MerchanterPwd    string
	MobilePhone      string
	Pwd              string
	IDNumber         string
	MerchantId       string
}

func (a *Merchanter) SelectByOpenId() error {
	row := mySqlDB.QueryRow(" select merchanterId,merchanterOpenId,merchanterName,mobilePhone,pwd,merchantId from Merchanter where merchanterOpenId = ? ", a.MerchanterOpenId)
	err := row.Scan(&a.MerchanterId, &a.MerchanterOpenId, &a.MerchanterName, &a.MobilePhone, &a.Pwd, &a.MerchantId)
	if err == sql.ErrNoRows {
		return errors.New("登陆密码错误")
	}
	return err
}
