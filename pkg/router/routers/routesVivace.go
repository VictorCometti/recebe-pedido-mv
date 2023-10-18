package routers

import (
	"application/pkg/handler"
	"net/http"
)

var routes = []Route{
	{
		URI:    "/recebe-pedido",
		Method: http.MethodPost,
		Func:   handler.RecebePedidoVivace,
	},
}
