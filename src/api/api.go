package api

import (
	"net/http"
	"fmt"
	"models"
	"encoding/json"
)


func GetVacancy (w http.ResponseWriter, r *http.Request) {
	var vacansies []models.Vacancy
	models.Database.Connect.Find(&vacansies)
	data, err := json.Marshal(vacansies)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	fmt.Fprintf(w, "%s", string(data))
}
