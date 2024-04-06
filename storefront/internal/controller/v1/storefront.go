package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"storefront/internal/entity"
	"storefront/internal/usecase"
	"storefront/pkg/web"
)

type storefrontRoutes struct {
	s usecase.StorefrontContract
}

type allProductsResponse struct {
	Products []entity.ProductList `json:"products"`
	Service  string               `json:"service"`
}

func NewUserRoutes(r chi.Router, s usecase.StorefrontContract) {
	sr := &storefrontRoutes{s: s}

	r.Get("/", sr.GetAllProducts)
	r.Get("/{}", sr.GetProductByUUID)
}

func (sr *storefrontRoutes) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := sr.s.GetAllProducts(r.Context())
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Println("Render error")
			return
		}
		return
	}
	response := allProductsResponse{Products: products, Service: "storefront"}
	render.JSON(w, r, response)
}

func (sr *storefrontRoutes) GetProductByUUID(w http.ResponseWriter, r *http.Request) {}
