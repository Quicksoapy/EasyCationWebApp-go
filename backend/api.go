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

	router.Handle("/book/{id}", BookDownloader).Methods("GET", "OPTIONS")

	router.Handle("/register", RegisterHandler).Methods("POST", "OPTIONS")

	router.Handle("/login", LoginHandler).Methods("GET", "OPTIONS")

	router.Handle("/CheckLogin", CheckLoginHandler).Methods("GET", "OPTIONS")

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	err = http.ListenAndServe(":"+strconv.Itoa(8000), handlers.CORS(origins, headers, methods)(router))
	return
}
