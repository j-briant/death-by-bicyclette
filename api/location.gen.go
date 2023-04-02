// Package location provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package location

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add a new location
	// (POST /location)
	AddLocation(c *gin.Context)
	// Update an existing location
	// (PUT /location)
	UpdateLocation(c *gin.Context)
	// Deletes a location
	// (DELETE /location/{locationId})
	DeleteLocation(c *gin.Context, locationId int64, params DeleteLocationParams)
	// Find location by ID
	// (GET /location/{locationId})
	GetLocationById(c *gin.Context, locationId int64)
	// Update a location with form data
	// (POST /location/{locationId})
	UpdateLocationWithForm(c *gin.Context, locationId int64)
	// Create user
	// (POST /user)
	CreateUser(c *gin.Context)
	// Creates list of users with given input array
	// (POST /user/createWithList)
	CreateUsersWithListInput(c *gin.Context)
	// Logs user into the system
	// (GET /user/login)
	LoginUser(c *gin.Context, params LoginUserParams)
	// Logs out current logged in user session
	// (GET /user/logout)
	LogoutUser(c *gin.Context)
	// Delete user
	// (DELETE /user/{username})
	DeleteUser(c *gin.Context, username string)
	// Get user by user name
	// (GET /user/{username})
	GetUserByName(c *gin.Context, username string)
	// Update user
	// (PUT /user/{username})
	UpdateUser(c *gin.Context, username string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AddLocation operation middleware
func (siw *ServerInterfaceWrapper) AddLocation(c *gin.Context) {

	c.Set(Location_authScopes, []string{"write:locations", "read:locations"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.AddLocation(c)
}

// UpdateLocation operation middleware
func (siw *ServerInterfaceWrapper) UpdateLocation(c *gin.Context) {

	c.Set(Location_authScopes, []string{"write:locations", "read:locations"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdateLocation(c)
}

// DeleteLocation operation middleware
func (siw *ServerInterfaceWrapper) DeleteLocation(c *gin.Context) {

	var err error

	// ------------- Path parameter "locationId" -------------
	var locationId int64

	err = runtime.BindStyledParameter("simple", false, "locationId", c.Param("locationId"), &locationId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter locationId: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(Location_authScopes, []string{"write:locations", "read:locations"})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteLocationParams

	headers := c.Request.Header

	// ------------- Optional header parameter "api_key" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("api_key")]; found {
		var ApiKey string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for api_key, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "api_key", runtime.ParamLocationHeader, valueList[0], &ApiKey)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter api_key: %s", err), http.StatusBadRequest)
			return
		}

		params.ApiKey = &ApiKey

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteLocation(c, locationId, params)
}

// GetLocationById operation middleware
func (siw *ServerInterfaceWrapper) GetLocationById(c *gin.Context) {

	var err error

	// ------------- Path parameter "locationId" -------------
	var locationId int64

	err = runtime.BindStyledParameter("simple", false, "locationId", c.Param("locationId"), &locationId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter locationId: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(Api_keyScopes, []string{""})

	c.Set(Location_authScopes, []string{"write:locations", "read:locations"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetLocationById(c, locationId)
}

// UpdateLocationWithForm operation middleware
func (siw *ServerInterfaceWrapper) UpdateLocationWithForm(c *gin.Context) {

	var err error

	// ------------- Path parameter "locationId" -------------
	var locationId int64

	err = runtime.BindStyledParameter("simple", false, "locationId", c.Param("locationId"), &locationId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter locationId: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(Location_authScopes, []string{"write:locations", "read:locations"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdateLocationWithForm(c, locationId)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateUser(c)
}

// CreateUsersWithListInput operation middleware
func (siw *ServerInterfaceWrapper) CreateUsersWithListInput(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateUsersWithListInput(c)
}

// LoginUser operation middleware
func (siw *ServerInterfaceWrapper) LoginUser(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginUserParams

	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", c.Request.URL.Query(), &params.Username)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "password" -------------

	err = runtime.BindQueryParameter("form", true, false, "password", c.Request.URL.Query(), &params.Password)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter password: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.LoginUser(c, params)
}

// LogoutUser operation middleware
func (siw *ServerInterfaceWrapper) LogoutUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.LogoutUser(c)
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameter("simple", false, "username", c.Param("username"), &username)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteUser(c, username)
}

// GetUserByName operation middleware
func (siw *ServerInterfaceWrapper) GetUserByName(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameter("simple", false, "username", c.Param("username"), &username)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetUserByName(c, username)
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameter("simple", false, "username", c.Param("username"), &username)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdateUser(c, username)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/location", wrapper.AddLocation)

	router.PUT(options.BaseURL+"/location", wrapper.UpdateLocation)

	router.DELETE(options.BaseURL+"/location/:locationId", wrapper.DeleteLocation)

	router.GET(options.BaseURL+"/location/:locationId", wrapper.GetLocationById)

	router.POST(options.BaseURL+"/location/:locationId", wrapper.UpdateLocationWithForm)

	router.POST(options.BaseURL+"/user", wrapper.CreateUser)

	router.POST(options.BaseURL+"/user/createWithList", wrapper.CreateUsersWithListInput)

	router.GET(options.BaseURL+"/user/login", wrapper.LoginUser)

	router.GET(options.BaseURL+"/user/logout", wrapper.LogoutUser)

	router.DELETE(options.BaseURL+"/user/:username", wrapper.DeleteUser)

	router.GET(options.BaseURL+"/user/:username", wrapper.GetUserByName)

	router.PUT(options.BaseURL+"/user/:username", wrapper.UpdateUser)

	return router
}