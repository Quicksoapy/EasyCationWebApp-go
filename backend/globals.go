package backend

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var DB *sql.DB

func InitializeDB() (err error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5555, "root", "root", "EasyCation")

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected!")

	DB = db
	return
}

var s = securecookie.New(securecookie.GenerateRandomKey(50), securecookie.GenerateRandomKey(32))

func SetAccountCookie(w http.ResponseWriter, inputValue string) {
	//make securecookie instance, first value how to encrypt name, second to encrypt value

	value := map[string]string{
		"AccountId": inputValue,
	}
	if encoded, err := s.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:     "cookie",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}

func ReadAccountCookie(r *http.Request) string {
	if cookie, err := r.Cookie("cookie"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("cookie", cookie.Value, &value); err == nil {
			return value["AccountId"]
		}
	}
	return ""
}

func hash(input string) ([]byte, error) {
	output, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		return output, err
	}
	return output, err
}
