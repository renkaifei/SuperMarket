package menu

import (
	"io/ioutil"
	"net/http"
	"strings"
	"superMarket/repo"
)

func CreateMenu() (string, error) {
	content, err := ioutil.ReadFile("menu.json")
	if err != nil {
		return "", err
	}
	menu := string(content)
	accessToken, err := repo.GetKey("wx_access_token")
	if err != nil {
		return "", err
	}
	rd := strings.NewReader(menu)
	req, err := http.NewRequest("post", "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+accessToken, rd)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func DeleteMenu() (string, error) {
	accessToken, err := repo.GetKey("wx_access_token")
	if err != nil {
		return "", err
	}
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" + accessToken)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	return string(result), err
}
