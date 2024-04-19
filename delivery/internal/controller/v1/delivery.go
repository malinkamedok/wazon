package v1

import (
	"delivery/internal/entity"
	"delivery/internal/usecase"
	"delivery/pkg/web"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type deliveryRoutes struct {
	s usecase.DeliveryContract
	m usecase.DeliveryMQContract
}

type allOrdersResponse struct {
	Orders  []entity.OrderList `json:"orders"`
	Service string             `json:"service"`
}

type orderResponse struct {
	Order   entity.Order `json:"order"`
	Service string       `json:"service"`
}

type orderChangeRequest struct {
	Id        string `json:"uuid"`
	NewStatus string `json:"status"`
}

func NewUserRoutes(r chi.Router, s usecase.DeliveryContract, m usecase.DeliveryMQContract) {
	dr := &deliveryRoutes{s: s, m: m}

	r.Get("/", dr.GetAllOrders)
	r.Get("/{uuid}", dr.GetOrderByUUID)
	r.Post("/create", dr.CreateOrder)
	r.Post("/update", dr.EditOrder)
}

func (dr *deliveryRoutes) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dr.s.GetAllOrders(r.Context())
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := allOrdersResponse{Orders: orders, Service: "delivery"}
	render.JSON(w, r, response)
}

func (dr *deliveryRoutes) GetOrderByUUID(w http.ResponseWriter, r *http.Request) {
	order, err := dr.s.GetOrderByUUID(r.Context(), chi.URLParam(r, "uuid"))
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		log.Println("No order obtained")
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := orderResponse{Order: order, Service: "delivery"}
	render.JSON(w, r, response)
}

func (dr *deliveryRoutes) EditOrder(w http.ResponseWriter, r *http.Request) {
	var orderIn orderChangeRequest
	err := json.NewDecoder(r.Body).Decode(&orderIn)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("JSON parse error")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}

	order, err := dr.s.UpdateOrderByUUID(r.Context(), orderIn.Id, orderIn.NewStatus)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("No order edited")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}

	err = dr.m.SendOrderUpdateMessage(r.Context(), orderIn.Id, orderIn.NewStatus)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("No order status update send to queue")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}

	response := orderResponse{Order: order, Service: "delivery"}
	render.JSON(w, r, response)

}

func (dr *deliveryRoutes) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderIn orderChangeRequest
	err := json.NewDecoder(r.Body).Decode(&orderIn)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("JSON parse error")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}

	order, err := dr.s.CreateOrder(r.Context(), orderIn.Id)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("No order created")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := orderResponse{Order: order, Service: "delivery"}
	render.JSON(w, r, response)
}
