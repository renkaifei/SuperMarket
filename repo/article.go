package repo

type Article struct {
	Id          int
	Title       string
	Description string
	PicUrl      string
	Url         string
}

func (a *Article) SelectByUrl(url string) error {
	row := mySqlDB.QueryRow("select articleId,articleTtile,articleDescription,articlePicUrl,articleUrl from Article where articleUrl = ? ", a.Url)
	err := row.Scan(&a.Id, &a.Title, &a.Description, &a.PicUrl, &a.Url)
	return err
}

type Articles struct {
	Values     []*Article
	PageSize   int
	PageIndex  int
	TotalCount int
}

func (a *Articles) SelectRecommend() error {
	row := mySqlDB.QueryRow("select Count(*) totalCount from article where isRecommend = 1 ")
	err := row.Scan(&a.TotalCount)
	if err != nil {
		return err
	}
	rows, err := mySqlDB.Query("select articleId,articleTitle,articleDescription,articlePicUrl,articleUrl from article where isRecommend = 1")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		value := &Article{}
		err = rows.Scan(&value.Id, &value.Title, &value.Description, &value.PicUrl, &value.Url)
		if err != nil {
			return err
		}
		a.Values = append(a.Values, value)
	}
	return nil
}
