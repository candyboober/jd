package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"jd/models"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secret = []byte("secret")

var SignIn = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// find user in db
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxBodySize))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	models.Database.Connect.Find(user, user)
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// write user info
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"username": user.Username,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})

var SignUp = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// find user in db
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxBodySize))
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
	models.Database.Connect.Create(user)
	if models.Database.Connect.NewRecord(user) {
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
})

//func authMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		secret := []byte("{YOUR-API-CLIENT-SECRET}")
//		secretProvider := auth0.NewKeyProvider(secret)
//		audience := []string{"{YOUR-AUTH0-API-AUDIENCE}"}
//
//		configuration := auth0.NewConfiguration(
//			secretProvider,
//			audience,
//			"https://{YOUR-AUTH0-DOMAIN}.auth0.com/",
//			jose.HS256)
//		validator := auth0.NewValidator(configuration)
//
//		token, err := validator.ValidateRequest(r)
//
//		if err != nil {
//			fmt.Println(err)
//			fmt.Println("Token is not valid:", token)
//			w.WriteHeader(http.StatusUnauthorized)
//			w.Write([]byte("Unauthorized"))
//		} else {
//			next.ServeHTTP(w, r)
//		}
//	})
//}

func ListVacancy(w http.ResponseWriter, r *http.Request) {
	var vacansies []models.Vacancy
	models.Database.Connect.Find(&vacansies)

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
	models.Database.Connect.First(vacancy, i)

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
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxBodySize))
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

	models.Database.Connect.Create(vacancy)
	if models.Database.Connect.NewRecord(vacancy) {
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

	models.Database.Connect.Save(vacancy)
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
	models.Database.Connect.Delete(vacancy)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vacancy); err != nil {
		panic(err)
	}
}
