package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
	"storefront/internal/entity"
	"storefront/internal/usecase"
	"storefront/pkg/logger"
	"storefront/pkg/web"
)

type storefrontRoutes struct {
	s usecase.StorefrontContract
}

type allProductsResponse struct {
	Products []entity.ProductList `json:"products"`
	Service  string               `json:"service"`
}

type productResponse struct {
	Product entity.Product `json:"product"`
	Service string         `json:"service"`
}

func NewUserRoutes(r chi.Router, s usecase.StorefrontContract) {
	sr := &storefrontRoutes{s: s}

	r.Get("/", sr.GetAllProducts)
	r.Get("/{uuid}", sr.GetProductByUUID)
}

func (sr *storefrontRoutes) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	logger.Info("got getAllProducts request")
	products, err := sr.s.GetAllProducts(r.Context())
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			logger.Error("render error", zap.Error(err))
			return
		}
		return
	}
	response := allProductsResponse{Products: products, Service: "storefront"}
	render.JSON(w, r, response)
}

func (sr *storefrontRoutes) GetProductByUUID(w http.ResponseWriter, r *http.Request) {
	logger.Info("got GetProductByUUID request")
	productID := chi.URLParam(r, "uuid")
	product, err := sr.s.GetProductByUUID(r.Context(), productID)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			logger.Error("render error", zap.Error(err))
			return
		}
		return
	}
	response := productResponse{Product: product, Service: "storefront"}
	render.JSON(w, r, response)
}
