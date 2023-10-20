package cart

import (
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/cart"
	"online-store/internal/resources"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"gorm.io/gorm"
)

func CreateCartsHandler(rt *service.Runtime, r cart.PostV1CartsParams, p *models.Principal) middleware.Responder {
	err := ValidateParameter(rt, r.Data)
	if err != nil {
		return handleError(rt, err)
	}

	data, err := ValidateCarts(rt, r.Data, p.UserID)
	if err != nil {
		return handleError(rt, err)
	}

	err = resources.CreateCarts(rt, data, p.UserID)
	if err != nil {
		return handleError(rt, err)
	}

	return cart.NewPostV1CartsCreated().WithPayload(&models.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success create cart",
	})
}

func ValidateParameter(rt *service.Runtime, request []*models.CartRequest) error {
	var productIDs []uint64
	for _, data := range request {
		productIDs = append(productIDs, *data.ProductID)
	}

	products, err := resources.GetProductByIDs(rt, productIDs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return rt.SetError(http.StatusNotFound, "record not found!")
		}

		return rt.SetError(http.StatusInternalServerError, err.Error())
	}

	if len(productIDs) > len(products) {
		return rt.SetError(http.StatusBadRequest, "bad request!")
	}

	return nil
}

func ValidateCarts(rt *service.Runtime, data []*models.CartRequest, userID int64) ([]*models.Cart, error) {
	carts, _, err := resources.GetCarts(rt, &resources.GetListParams{Page: swag.Int32(1), Limit: swag.Int32(10)}, userID)
	if err != nil {
		return nil, rt.SetError(http.StatusInternalServerError, err.Error())
	}

	if len(carts) == 0 {
		return createNewCarts(data, userID), nil
	}

	updateExistingCarts(carts, data, userID)

	return carts, nil

}

func createNewCarts(data []*models.CartRequest, userID int64) []*models.Cart {
	var carts []*models.Cart
	for _, cart := range data {
		carts = append(carts, &models.Cart{
			ProductID: *cart.ProductID,
			UserID:    uint64(userID),
			CartData: models.CartData{
				Quantity: *cart.Quantity,
			},
		})
	}
	return carts
}

func updateExistingCarts(carts []*models.Cart, data []*models.CartRequest, userID int64) {
	cartMap := make(map[uint64]*models.Cart)
	for _, cart := range carts {
		productID := cart.ProductID
		cartMap[productID] = cart
	}

	for _, item := range data {
		productID := *item.ProductID
		if cart, ok := cartMap[productID]; ok {
			cart.Quantity = *item.Quantity
		} else {
			carts = append(carts, &models.Cart{
				ProductID: productID,
				UserID:    uint64(userID),
				CartData: models.CartData{
					Quantity: *item.Quantity,
				},
			})
		}
	}
}

func handleError(rt *service.Runtime, err error) middleware.Responder {
	rt.Error().Println(err)
	errData := rt.GetError(err)
	return cart.NewPostV1CartsDefault(int(errData.Code())).WithPayload(&models.BaseResponse{
		Code:    errData.Code(),
		Message: errData.Error(),
	})
}
