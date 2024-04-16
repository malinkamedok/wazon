package v1

import (
	"accountservice/internal/entity"
	"accountservice/internal/usecase"
	"accountservice/pkg/web"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type accountServiceRoutes struct {
	service usecase.AccountServiceContract
	rest    usecase.IntegrationContract
}

type userResponse struct {
	User    entity.User `json:"user"`
	Service string      `json:"service"`
}

type createProductResponse struct {
	Product entity.Product `json:"product"`
	Service string         `json:"service"`
}

type productsResponse struct {
	Products []entity.Product `json:"product"`
	Service  string           `json:"service"`
}

func NewUserRoutes(router chi.Router, contract usecase.AccountServiceContract, rest usecase.IntegrationContract) {
	route := &accountServiceRoutes{service: contract, rest: rest}

	router.Get("/user/{id}", route.GetUserById)
	router.Get("/healthCheck", route.HealthCheck)
	router.Get("/getAllProducts", route.GetAllProducts)
	router.Get("/cart/user/{id}", route.GetAllProductsFromCart)
	router.Post("/product", route.CreateProduct)
}

func (routes *accountServiceRoutes) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("convert to uuid error")
		return
	}
	user, err := routes.service.GetUserById(r.Context(), userId)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := userResponse{User: user, Service: "accountservice"}
	render.JSON(w, r, response)
}

func (routes *accountServiceRoutes) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := "accountservice alive!"
	render.JSON(w, r, response)
}

func (routes *accountServiceRoutes) CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос на создание/обновление продукта")
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("JSON parse error")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}
	err = routes.service.InsertOrUpdateProduct(r.Context(), product)
	if err != nil {
		errRender := render.Render(w, r, web.ErrRender(err))
		log.Println("No order edited")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := createProductResponse{Product: product, Service: "accountservice"}
	render.JSON(w, r, response)
}

func (routes *accountServiceRoutes) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	product := routes.rest.GetAllProducts()
	product.Service = "accountservice"
	render.JSON(w, r, product)
}

func (routes *accountServiceRoutes) GetAllProductsFromCart(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("convert to uuid error")
		return
	}
	product, err := routes.service.GetAllProductsFromCart(r.Context(), id)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := productsResponse{Products: product, Service: "accountservice"}
	render.JSON(w, r, response)
}
