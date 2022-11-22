package backend

import (
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/securecookie"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var BookDownloader = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")

	http.Redirect(w, r, "/book/{id}", http.StatusSeeOther)
	return
})

type userInputRegister struct {
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Region           string    `json:"region,omitempty"`
	Country          string    `json:"country,omitempty"`
	RegisterDateTime time.Time `json:"datetime,omitempty"`
}

type userInputLogin struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	LoginDateTime time.Time `json:"datetime,omitempty"`
}

var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")

	v := &userInputRegister{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
	}
	bodyEscaped, err := strconv.Unquote(string(body))
	if err != nil {
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
	}
	err = json.Unmarshal([]byte(bodyEscaped), v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
	}

	err, _ = RegisterUser(DB, *v)
	if err != nil {

	}
})

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")

	v := &userInputLogin{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}
	var klantId int
	err, klantId = LoginUser(DB, *v)
	if err != nil {
		return
	}

	SetAccountCookie(w, strconv.Itoa(klantId))
})

type Response struct {
	AccountId string
}

var CheckLoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")

	w.WriteHeader(http.StatusOK)
	response := Response{AccountId: ReadAccountCookie(r)}
	res, err := json.Marshal(response)
	if err != nil {
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
		return
	}
	_, err = fmt.Fprint(w, res)
	if err != nil {
		return
	}
})
