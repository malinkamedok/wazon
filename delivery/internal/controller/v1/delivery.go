package v1

import (
	"delivery/internal/entity"
	"delivery/internal/usecase"
	"delivery/pkg/web"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type deliveryRoutes struct {
	s usecase.DeliveryContract
}

type allOrdersResponse struct {
	Orders  []entity.OrderList `json:"orders"`
	Service string             `json:"service"`
}

type orderResponse struct {
	Order   entity.Order `json:"order"`
	Service string       `json:"service"`
}

func NewUserRoutes(r chi.Router, s usecase.DeliveryContract) {
	dr := &deliveryRoutes{s: s}

	r.Get("/", dr.GetAllOrders)
	r.Get("/{uuid}", dr.GetOrderByUUID)
	r.Post("/{uuid}/edit", dr.EditOrder)
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
	uuid, err := uuid.Parse(chi.URLParam(r, "uuid"))
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		log.Println("Incorrect uuid")
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	order, err := dr.s.GetOrderByUUID(r.Context(), uuid)
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

}
