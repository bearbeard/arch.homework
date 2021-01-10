package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const version = "1.0"

type response struct {
	Status  string
	Message string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthHandler).
		Methods("GET")
	router.HandleFunc("/version", versionHandler).
		Methods("GET")
	router.HandleFunc("/", defaultHandler).
		Methods("GET")

	//pgConfig := pgx.ConnConfig{
	//	Database: os.Getenv("DATABASE_URL"),
	//	User: os.Getenv("DATABASE_USER"),
	//	Password: os.Getenv("DATABASE_PASSWORD"),
	//}
	//conn, err := pgx.Connect(pgConfig)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close()
	//
	//userService := user.NewUserService(conn)
	//
	//router.HandleFunc("/user/{id}", userService.GetUser).
	//	Methods("GET")
	//router.HandleFunc("/user", userService.CreateUser).
	//	Methods("PUT")
	//router.HandleFunc("/user", userService.UpdateUser).
	//	Methods("POST")
	//router.HandleFunc("/user/{id}", userService.DeleteUser).
	//	Methods("DELETE")

	http.Handle("/", router)
	server := http.Server{
		Handler:      router,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	response := response{}
	rw.WriteHeader(http.StatusOK)
	response.Status = http.StatusText(http.StatusOK)
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshalling response")
	}
	_, _ = rw.Write(bytes)
}

func versionHandler(rw http.ResponseWriter, r *http.Request) {
	response := response{}
	rw.WriteHeader(http.StatusOK)
	response.Message = version
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshalling response")
	}
	_, _ = rw.Write(bytes)
}

func defaultHandler(rw http.ResponseWriter, r *http.Request) {
	response := response{}
	rw.WriteHeader(http.StatusOK)
	response.Status = "Trololo"
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshalling response")
	}
	_, _ = rw.Write(bytes)
}
