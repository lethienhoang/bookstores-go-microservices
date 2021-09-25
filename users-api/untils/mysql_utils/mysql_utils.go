package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/bookstores/users-api/untils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRow       = "no rows in result set"
	errorEmailUnique = "email_UNIQUE"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError("no record matching given id")
		}
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicate key")
	}

	return errors.NewInternalError(fmt.Sprintf("number: %d, message: %s", sqlErr.Number, sqlErr.Message))
}
