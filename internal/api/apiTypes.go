package api

import (
	gmux "github.com/gorilla/mux"
	logging "github.com/kthucydi/bs_go_logrus"
	"messaggio_test/internal/config"
	"messaggio_test/internal/pgstorage"
	"net/http"
)

var (
	PG  = &pgstorage.PGDB
	cfg = config.Cfg
	Log = &logging.Log
)

type JSONMessage struct {
	Message   string `json:"message"`
	ErrorBool bool   `json:"errorbool"`
}

type ErrMessage struct {
	Message string `json:"errmessage"`
}

type Resp struct {
	Message     string
	ContentType string
	StructOut   interface{}
}

type SetVar struct {
	Auth     string
	StructIn interface{}
	Handler  http.HandlerFunc
}

type Route struct {
	Auth        string
	Tags        []string
	Description string
	Name        string
	Sets        SetVar
	Handler     http.HandlerFunc
	Middlewares []gmux.MiddlewareFunc
	Responces   map[int]Resp
}

type ApiRoutes struct {
	Routes           map[string]Methods
	AuthMiddleware   map[string]gmux.MiddlewareFunc
	CommonMiddleware map[string]gmux.MiddlewareFunc
}

type Methods map[string]Route
