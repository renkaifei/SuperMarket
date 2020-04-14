package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"superMarket/repo"
)

type UniversityController struct {
}

func (a *UniversityController) Create(w http.ResponseWriter, r *http.Request) {
	universityName := r.PostFormValue("universityName")
	universityCode := r.PostFormValue("universityCode")
	universityAddress := r.PostFormValue("universityAddress")
	universityPicture := r.PostFormValue("universityPicture")

	university := &repo.University{UniversityName: universityName, UniversityCode: universityCode, UniversityAddress: universityAddress, UniversityPicture: universityPicture}
	err := university.Create()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(university)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
