package product

import (
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/product"
	"online-store/internal/resources"

	"github.com/go-openapi/runtime/middleware"
)

func CreateProducts(rt *service.Runtime, r product.PostV1ProductsParams, p *models.Principal) middleware.Responder {
	err := resources.CreateProduct(rt, r.Data)
	if err != nil {
		rt.Error().Println(err)
		return product.NewPostV1ProductsDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return product.NewPostV1ProductsCreated().WithPayload(&models.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success create new product",
	})
}
