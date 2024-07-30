package api

import (
	"fmt"

	gmux "github.com/gorilla/mux"
	server "github.com/kthucydi/bs_go_server"
	pgstorage "messaggio_test/internal/pgstorage"
)

var API *ApiRoutes

func init() {
	Log.Debugf("FROM API GetNow:%s ", fmt.Sprint(pgstorage.GetNow(pgstorage.PGDB)))
	API = &ApiRoutes{}

	API.CommonMiddleware = map[string]gmux.MiddlewareFunc{}
	API.AuthMiddleware = map[string]gmux.MiddlewareFunc{}
	API.Routes = make(map[string]Methods)

	API.Routes["/"] = Methods{
		"POST": Route{Name: "postMessage", Auth: "", Handler: messageHandler},
	}

	API.Routes["/getstat"] = Methods{
		"GET": Route{Name: "getStatistic", Auth: "", Handler: statHandler},
	}
}

func (a *ApiRoutes) GetRouteList() (routes map[string]server.Methods) {
	routes = make(map[string]server.Methods)

	for path, methods := range API.Routes {
		routes[path] = map[string]server.Route{}
		for method, apiRoute := range methods {
			routes[path][method] = server.Route{
				Auth:        apiRoute.Auth,
				Name:        apiRoute.Name,
				Handler:     apiRoute.Handler,
				Middlewares: apiRoute.Middlewares,
			}
		}
	}
	return routes
}

func (a *ApiRoutes) GetCommonMiddleware() map[string]gmux.MiddlewareFunc {
	return a.CommonMiddleware
}

func (a *ApiRoutes) GetAuthMiddleware() map[string]gmux.MiddlewareFunc {
	return a.AuthMiddleware
}
