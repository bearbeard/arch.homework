package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"log"
	"net/http"
	"strconv"
)

type service interface {
	GetUser(rw http.ResponseWriter, r *http.Request)
	CreateUser(rw http.ResponseWriter, r *http.Request)
	UpdateUser(rw http.ResponseWriter, r *http.Request)
	DeleteUser(rw http.ResponseWriter, r *http.Request)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(db *pgx.Conn) UserService {
	s := UserService{
		repo: UserRepository{db: db},
	}
	return s
}

func (s *UserService) GetUser(rw http.ResponseWriter, r *http.Request) {
	response := response{}
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	user, err := s.repo.FindUser(id)
	if err != nil {
		log.Println("ERRROR " + err.Error())
		response.Message = err.Error()
	} else {
		response.User = user
	}
	bytes, _ := json.Marshal(response)
	_, _ = rw.Write(bytes)
}

func (s *UserService) CreateUser(rw http.ResponseWriter, r *http.Request) {
	request := request{}
	response := response{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("ERRROR " + err.Error())
	}
	user := User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Surname:   request.Surname,
		Age:       request.Age,
		Gender:    request.Gender,
	}
	id, err := s.repo.InsertUser(user)
	if err != nil {
		log.Println("ERRROR " + err.Error())
		response.Message = err.Error()
	} else {
		user.ID = id
		response.User = user
	}
	bytes, _ := json.Marshal(response)
	_, _ = rw.Write(bytes)
}

func (s *UserService) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	request := request{}
	response := response{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("ERRROR " + err.Error())
	}
	user := User{
		ID:        request.ID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Surname:   request.Surname,
		Age:       request.Age,
		Gender:    request.Gender,
	}
	_, err = s.repo.InsertUser(user)
	if err != nil {
		log.Println("ERRROR " + err.Error())
		response.Message = err.Error()
	} else {
		response.User = user
	}
	bytes, _ := json.Marshal(response)
	_, _ = rw.Write(bytes)
}

func (s *UserService) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	request := request{}
	response := response{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("ERRROR " + err.Error())
	}

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	err = s.repo.DeleteUser(id)
	if err != nil {
		log.Println("ERRROR " + err.Error())
		response.Message = err.Error()
	} else {
		response.User.ID = id
	}
	bytes, _ := json.Marshal(response)
	_, _ = rw.Write(bytes)
}
