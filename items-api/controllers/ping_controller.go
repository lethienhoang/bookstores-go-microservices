package controllers

import (
	"net/http"
)

const (
	pong = "pong"
)

var (
	PingsController IPingControllerInterface = &PingController{}
)

type IPingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type PingController struct{}

func (c *PingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}