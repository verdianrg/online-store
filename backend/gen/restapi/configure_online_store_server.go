// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"online-store/gen/models"
	"online-store/gen/restapi/operations"
	"online-store/gen/restapi/operations/cart"
	"online-store/gen/restapi/operations/history"
	"online-store/gen/restapi/operations/login"
	"online-store/gen/restapi/operations/product"
)

//go:generate swagger generate server --target ../../gen --name OnlineStoreServer --spec ../../api/swagger.yml --principal models.Principal --exclude-main

func configureFlags(api *operations.OnlineStoreServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OnlineStoreServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-online-store-key" header is set
	if api.KeyAuth == nil {
		api.KeyAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (key) x-online-store-key from header param [x-online-store-key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.CartGetV1CartsHandler == nil {
		api.CartGetV1CartsHandler = cart.GetV1CartsHandlerFunc(func(params cart.GetV1CartsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation cart.GetV1Carts has not yet been implemented")
		})
	}
	if api.ProductGetV1ProductsHandler == nil {
		api.ProductGetV1ProductsHandler = product.GetV1ProductsHandlerFunc(func(params product.GetV1ProductsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation product.GetV1Products has not yet been implemented")
		})
	}
	if api.CartPostV1CartsHandler == nil {
		api.CartPostV1CartsHandler = cart.PostV1CartsHandlerFunc(func(params cart.PostV1CartsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation cart.PostV1Carts has not yet been implemented")
		})
	}
	if api.HistoryPostV1HistoriesHandler == nil {
		api.HistoryPostV1HistoriesHandler = history.PostV1HistoriesHandlerFunc(func(params history.PostV1HistoriesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation history.PostV1Histories has not yet been implemented")
		})
	}
	if api.LoginPostV1LoginHandler == nil {
		api.LoginPostV1LoginHandler = login.PostV1LoginHandlerFunc(func(params login.PostV1LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation login.PostV1Login has not yet been implemented")
		})
	}
	if api.ProductPostV1ProductsHandler == nil {
		api.ProductPostV1ProductsHandler = product.PostV1ProductsHandlerFunc(func(params product.PostV1ProductsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation product.PostV1Products has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
