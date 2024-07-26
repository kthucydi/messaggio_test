package api

import (
	"fmt"

	gmux "github.com/gorilla/mux"
	server "github.com/kthucydi/bs_go_server"
)

var API *ApiRoutes

func init() {
	Log.Debugf("FROM API GetNow:%s ", fmt.Sprint(PG.GetNow()))
	API = &ApiRoutes{}

	API.CommonMiddleware = map[string]gmux.MiddlewareFunc{}
	API.AuthMiddleware = map[string]gmux.MiddlewareFunc{}
	API.Routes = make(map[string]Methods)

	API.Routes["/"] = Methods{
		"Post": Route{Name: "basicListener", Auth: "", Handler: messageHandler},
	}

	API.Routes["/status"] = Methods{
		"Post": Route{Name: "basicListener", Auth: "", Handler: serverStatus},
	}

	API.Routes["/getstat"] = Methods{
		"GET": Route{Name: "requestCodeToEmail", Auth: "", Handler: statHandler},
		// 		Middlewares: []gmux.MiddlewareFunc{mw.Jwt.Handler}},

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
