package users

import (
	"fmt"

	"github.com/bookstores/users-api/db"
	"github.com/bookstores/users-api/untils/crypto"
	"github.com/bookstores/users-api/untils/date_utils"
	"github.com/bookstores/users-api/untils/errors"
	"github.com/bookstores/users-api/untils/loggers"
	"github.com/bookstores/users-api/untils/mysql_utils"
)

type User struct {
	Id          int64 `copier:"-"`
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
	Status      string
	Password    string
}

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created, Status, Password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT Id, first_name, last_name, email, date_created FROM users WHERE Id = ?;"
	queryUpdateUser = "UPDATE users SET first_name = ?, last_name = ?, email = ?, date_created = ? WHERE Id = ?;"
	queryDeleteUser = "DELETE FROM users WHERE Id = ?"
	queryFindUser   = "SELECT Id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

const (
	StatusActive   = "ACTIVE"
	StatusInActive = "INACTIVE"
)

func (user *User) Get() *errors.RestError {
	stmt, err := db.DB.Prepare(queryGetUser)
	if err != nil {
		loggers.Error("error when trying to prepare statement", err)
		return errors.NewInternalError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		loggers.Error("error when trying to get data from db", err)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Create() *errors.RestError {
	stmt, err := db.DB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	user.Status = StatusActive
	user.Password = crypto.GetMd5Hash(user.Password)
	resultExec, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := resultExec.LastInsertId()
	if err != nil {
		return errors.NewInternalError(fmt.Sprintf("error when trying to get user id: %s", err.Error()))
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := db.DB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	_, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Id)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := db.DB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	defer stmt.Close()
	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := db.DB.Prepare(queryFindUser)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}

	return results, nil
}
