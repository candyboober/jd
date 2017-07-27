package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"jd/models"
)

func GetVacancyList(w http.ResponseWriter, r *http.Request) {
	var vacansies []models.Vacancy
	models.Database.Connect.Find(&vacansies)
	err := json.NewEncoder(w).Encode(vacansies)
	if err != nil {
		panic(err)
	}
}

func GetVacancy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	v := models.Vacancy{}
	v.Get(id)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}
