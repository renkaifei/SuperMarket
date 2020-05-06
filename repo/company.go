package repo

type Company struct {
	CompanyId      int
	CompanyName    string
	CompanyAddress string
	IndustryId     string
}

func (a *Company) SelectById() error {
	row := mySqlDB.QueryRow(" select companyId,companyName,companyAddress from company where companyId = ? ", a.CompanyId)
	err := row.Scan(&a.CompanyId, &a.CompanyName, &a.CompanyAddress)
	return err
}

type Companys struct {
	Values []*Company
}

func (a *Companys) SelectByIndustry(industryId int) error {
	rows, err := mySqlDB.Query("select companyId,companyName,companyAddress from company where industryId = ? ", industryId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		company := &Company{}
		if err = rows.Scan(&company.CompanyId, &company.CompanyName, &company.CompanyAddress); err != nil {
			return err
		}
		a.Values = append(a.Values, company)
	}
	return nil
}
