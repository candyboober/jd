package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"jd/core"
	"jd/models"
	"net/http"
	"strconv"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return core.Secret, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func ListVacancy(w http.ResponseWriter, r *http.Request) {
	var vacansies []models.Vacancy
	core.Database.Connect.Find(&vacansies)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vacansies); err != nil {
		panic(err)
	}
}

func RetrieveVacancy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	vacancy := &models.Vacancy{}
	core.Database.Connect.First(vacancy, i)

	if vacancy.ID == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vacancy); err != nil {
		panic(err)
	}
}

func CreateVacancy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, core.MaxBodySize))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	vacancy := &models.Vacancy{}
	if err := json.Unmarshal(body, vacancy); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	core.Database.Connect.Create(vacancy)
	if core.Database.Connect.NewRecord(vacancy) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vacancy); err != nil {
		panic(err)
	}
}

func UpdateVacancy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	vacancy := &models.Vacancy{ID: uint(idInt)}

	if err := json.Unmarshal(body, &vacancy); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	core.Database.Connect.Save(vacancy)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vacancy); err != nil {
		panic(err)
	}
}

func DestroyVacancy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	vacancy := &models.Vacancy{ID: uint(idInt)}
	core.Database.Connect.Delete(vacancy)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vacancy); err != nil {
		panic(err)
	}
}
