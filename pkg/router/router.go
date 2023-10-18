package router

import (
	"application/pkg/router/routers"
	"github.com/gorilla/mux"
)

func GetRouter() (r *mux.Router) {
	r = mux.NewRouter()
	routers.ConfigureRoute(r)
	return
}
