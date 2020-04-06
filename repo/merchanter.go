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
	IsAdmin          int
}

func (a *Merchanter) SelectByOpenId() error {
	row := mySqlDB.QueryRow(" select merchanterId,merchanterOpenId,merchanterName,mobilePhone,pwd,merchantId,IsAdmin from Merchanter where merchanterOpenId = ? ", a.MerchanterOpenId)
	err := row.Scan(&a.MerchanterId, &a.MerchanterOpenId, &a.MerchanterName, &a.MobilePhone, &a.Pwd, &a.MerchantId, &a.IsAdmin)
	if err == sql.ErrNoRows {
		return errors.New("登陆密码错误")
	}
	return err
}
