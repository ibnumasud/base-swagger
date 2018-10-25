// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	cors "github.com/rs/cors"

	"github.com/ibnumasud/base-swagger/models"
	"github.com/ibnumasud/base-swagger/restapi/operations"
	"github.com/ibnumasud/base-swagger/restapi/operations/default_operations"
	"github.com/ibnumasud/base-swagger/restapi/operations/upload_s3"
)

//go:generate swagger generate server --target .. --name upload-s3 --spec ../swagger.yml

func configureFlags(api *operations.UploadS3API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UploadS3API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string) (*models.Principal, error) {
		if token == "abcdefuvwxyz" {
			prin := models.Principal(token)
			return &prin, nil
		}
		return nil, errors.New(401, "incorrect api key auth")

	}

	api.DefaultOperationsDefaultHandler = default_operations.DefaultHandlerFunc(func(params default_operations.DefaultParams, principal *models.Principal) middleware.Responder {
		return default_operations.NewDefaultOK()
	})
	api.UploadS3UploadS3Handler = upload_s3.UploadS3HandlerFunc(upload_s3.Upload)

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}
