package handlers

import (
	"WildberriesL0/server/internal/cache"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type orderHandler struct {
	orders *cache.OrderCache
}

func RouterHandler(orders *cache.OrderCache) *orderHandler {
	return &orderHandler{orders: orders}
}

func (handler *orderHandler) Register(router *httprouter.Router) {
	router.GET("/:uuid", handler.getOrder())
}

func (handler *orderHandler) getOrder() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		orderUUID := ps.ByName("uuid")
		order := handler.orders.GetOrderByUUID(orderUUID)

		if order != nil {
			w.WriteHeader(200)
			orderJson, err := json.MarshalIndent(&order, "", "    ")
			if err != nil {

			}
			w.Write(orderJson)
		}
	}
}
