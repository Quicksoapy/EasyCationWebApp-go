package backend

import (
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

func HandleRequests() (err error) {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/book/{id}", BookDownloader).Methods("GET")

	router.Handle("/register", RegisterHandler).Methods("POST")

	router.Handle("/login", LoginHandler).Methods("GET")

	router.Handle("/CheckLogin", CheckLoginHandler).Methods("GET")

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	err = http.ListenAndServe(":"+strconv.Itoa(8000), handlers.CORS(origins, headers, methods)(router))
	return
}
