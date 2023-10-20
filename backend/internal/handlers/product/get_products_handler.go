package product

import (
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/product"
	"online-store/internal/resources"

	"github.com/go-openapi/runtime/middleware"
)

func GetProducts(rt *service.Runtime, r *product.GetV1ProductsParams, p *models.Principal) middleware.Responder {
	params := resources.GetListParams{
		Page:    r.Page,
		Limit:   r.Limit,
		Keyword: r.Keyword,
	}

	data, pagination, err := resources.GetProducts(rt, &params)
	if err != nil {
		rt.Error().Println(err)
		return product.NewGetV1ProductsDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return product.NewGetV1ProductsOK().WithPayload(&product.GetV1ProductsOKBody{
		Data:     data,
		Metadata: pagination,
	})
}
