package controllers

import (
	"net/http"
	"strconv"

	"github.com/bookstores/users-api/requests"
	"github.com/bookstores/users-api/services"
	"github.com/bookstores/users-api/untils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, psrErr := getUserId(c.Param("user_id"))
	if psrErr != nil {
		c.JSON(psrErr.Code, psrErr)
		return
	}

	result, getErr := services.UserService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CreateUser(c *gin.Context) {
	var request requests.CreateOrUpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := services.UserService.Create(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	var request requests.CreateOrUpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	userId, psrErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if psrErr != nil {
		restErr := errors.NewBadRequestError("invalid params")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := services.UserService.Update(&request, userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, psrErr := getUserId(c.Param("user_id"))
	if psrErr != nil {
		c.JSON(psrErr.Code, psrErr)
		return
	}

	err := services.UserService.Delete(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func getUserId(userIdPram string) (int64, *errors.RestError) {
	userId, psrErr := strconv.ParseInt(userIdPram, 10, 64)
	if psrErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		newErr := errors.NewBadRequestError("status is required")
		c.JSON(newErr.Code, newErr)
		return
	}

	users, err := services.UserService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
