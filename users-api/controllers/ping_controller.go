package controllers

import (
	"net/http"

	"github.com/bookstores/users-api/db"
	"github.com/gin-gonic/gin"
)

func PingSerivce(c *gin.Context) {
	c.JSON(http.StatusOK, "services is running")
}

func PingDatabases(c *gin.Context) {
	if err := db.DB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "database is connected successfully")
}
