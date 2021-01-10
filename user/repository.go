package user

import (
	"github.com/jackc/pgx"
)

type repository interface {
	FindUser(id int64) (User, error)
	InsertUser(user User) (int64, error)
	UpdateUser(user User) error
	DeleteUser(id int64) error
}

type UserRepository struct {
	db *pgx.Conn
}

func (r *UserRepository) FindUser(id int64) (User, error) {
	user := User{}
	err := r.db.QueryRow("select id, firstname, lastname, surname, age, gender from USER where ID = $1", id).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Surname,
		&user.Age,
		&user.Gender,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) InsertUser(user User) (int64, error) {
	var id int64
	err := r.db.QueryRow("insert into USER (firstname, lastname, surname, age, gender) values ($1, $2, $3, $4, $5) returning id",
		user.Firstname,
		user.Lastname,
		user.Surname,
		user.Age,
		user.Gender,
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *UserRepository) UpdateUser(user User) error {
	repoUser, err := r.FindUser(user.ID)
	if err != nil {
		return err
	}
	if user.Firstname == "" {
		user.Firstname = repoUser.Firstname
	}
	if user.Lastname == "" {
		user.Lastname = repoUser.Lastname
	}
	if user.Surname == "" {
		user.Surname = repoUser.Surname
	}
	if user.Age == 0 {
		user.Age = repoUser.Age
	}
	if user.Gender == "" {
		user.Gender = repoUser.Gender
	}

	_, err = r.db.Exec("update USER set firstname = $2, lastname = $3, surname = $4, age = $5, gender = $6 where ID = $1",
		user.ID,
		user.Firstname,
		user.Lastname,
		user.Surname,
		user.Age,
		user.Gender,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int64) error {
	_, err := r.db.Exec("delete from USER where ID = $1", id)
	if err != nil {
		return err
	}
	return nil
}
