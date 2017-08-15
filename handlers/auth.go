package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"jd/core"
	"jd/models"
	"net/http"
	"time"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	// find user in db
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, core.MaxBodySize))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// if you read from body, error is not irrelevant
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	userData := &models.User{}
	if err := json.Unmarshal(body, userData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	userData.HashingPassword()
	user := &models.User{}
	core.Database.Connect.Find(user, user)
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// write user info
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(core.Secret)
	if err != nil {
		panic(err)
	}

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	// find user in db
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, core.MaxBodySize))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	user := &models.User{}
	if err := json.Unmarshal(body, user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	user.HashingPassword()
	core.Database.Connect.Create(user)
	if core.Database.Connect.NewRecord(user) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}
