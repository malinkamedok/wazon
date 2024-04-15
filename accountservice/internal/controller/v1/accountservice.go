package v1

import (
	"accountservice/internal/entity"
	"accountservice/internal/usecase"
	"accountservice/internal/usecase/integration"
	"accountservice/pkg/web"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strconv"
)

type accountServiceRoutes struct {
	accountServiceStruct usecase.AccountServiceContract
}

type userResponse struct {
	User    entity.User `json:"user"`
	Service string      `json:"service"`
}

type productResponse struct {
	Product entity.Product `json:"product"`
	Service string         `json:"service"`
}

func NewUserRoutes(router chi.Router, contract usecase.AccountServiceContract) {
	route := &accountServiceRoutes{accountServiceStruct: contract}

	router.Get("/user/{id}", route.GetUserById)
	router.Get("/healthCheck", route.HealthCheck)
	router.Get("/getAllProducts", route.GetAllProducts)
	router.Post("/product", route.CreateProduct)
}

func (routes *accountServiceRoutes) GetUserById(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	fmt.Printf("Получен id = %s", idStr)
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("convert to int error")
		return
	}
	fmt.Printf("Получен id = %d", idInt)
	user, err := routes.accountServiceStruct.GetUserById(request.Context(), idInt)
	if err != nil {
		err := render.Render(writer, request, web.ErrRender(err))
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := userResponse{User: user, Service: "accountservice"}
	render.JSON(writer, request, response)
}

func (routes *accountServiceRoutes) HealthCheck(writer http.ResponseWriter, request *http.Request) {
	response := "accountservice alive!"
	render.JSON(writer, request, response)
}

func (routes *accountServiceRoutes) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	log.Println("Получен запрос на создание/обновление продукта")
	var product entity.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		errRender := render.Render(writer, request, web.ErrRender(err))
		log.Println("JSON parse error")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}
	err = routes.accountServiceStruct.InsertOrUpdateProduct(request.Context(), product)
	if err != nil {
		errRender := render.Render(writer, request, web.ErrRender(err))
		log.Println("No order edited")
		if errRender != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := productResponse{Product: product, Service: "accountservice"}
	render.JSON(writer, request, response)
}

func (routes *accountServiceRoutes) GetAllProducts(writer http.ResponseWriter, request *http.Request) {
	product := integration.GetAllProducts()
	product.Service = "accountservice"
	render.JSON(writer, request, product)
}
