package rest

import (
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations"
	"online-store/gen/restapi/operations/cart"
	"online-store/gen/restapi/operations/history"
	"online-store/gen/restapi/operations/login"
	"online-store/gen/restapi/operations/product"
	"online-store/gen/restapi/operations/user"

	cart_domain "online-store/internal/handlers/cart"
	history_domain "online-store/internal/handlers/history"
	login_domain "online-store/internal/handlers/login"
	product_domain "online-store/internal/handlers/product"
	user_domain "online-store/internal/handlers/user"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *service.Runtime, api *operations.OnlineStoreServerAPI) {
	rt.Info().Println("Initialize Route")

	// LOGIN DOMAIN
	api.LoginPostV1LoginHandler = login.PostV1LoginHandlerFunc(func(pvlp login.PostV1LoginParams) middleware.Responder {
		return login_domain.Login(rt)(pvlp)
	})

	// USER DOMAIN
	api.UserPostV1UserHandler = user.PostV1UserHandlerFunc(func(pvup user.PostV1UserParams) middleware.Responder {
		return user_domain.CreateUserHandler(rt)(pvup)
	})

	api.UserGetV1UserIDHandler = user.GetV1UserIDHandlerFunc(func(gvui user.GetV1UserIDParams, p *models.Principal) middleware.Responder {
		return user_domain.GetUserHandler(rt)(gvui, p)
	})

	// PRODUCT DOMAIN
	api.ProductPostV1ProductsHandler = product.PostV1ProductsHandlerFunc(func(pvpp product.PostV1ProductsParams, p *models.Principal) middleware.Responder {
		return product_domain.CreateProducts(rt, pvpp, p)
	})

	api.ProductGetV1ProductsHandler = product.GetV1ProductsHandlerFunc(func(gvpp product.GetV1ProductsParams, p *models.Principal) middleware.Responder {
		return product_domain.GetProducts(rt, &gvpp, p)
	})

	// CART DOMAIN
	api.CartPostV1CartsHandler = cart.PostV1CartsHandlerFunc(func(pvcp cart.PostV1CartsParams, p *models.Principal) middleware.Responder {
		rt.Info().Println("kenapa ini priciple: ", p)
		return cart_domain.CreateCartsHandler(rt, pvcp, p)
	})

	api.CartGetV1CartsHandler = cart.GetV1CartsHandlerFunc(func(gvcp cart.GetV1CartsParams, p *models.Principal) middleware.Responder {
		return cart_domain.GetCartsHandler(rt, gvcp, p)
	})

	// HISTORY DOMAIN
	api.HistoryPostV1HistoriesHandler = history.PostV1HistoriesHandlerFunc(func(pvhp history.PostV1HistoriesParams, p *models.Principal) middleware.Responder {
		return history_domain.CreateHistoryHandler(rt, &pvhp, p)
	})

}
