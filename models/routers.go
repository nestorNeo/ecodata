/*
 * ecodata - OpenAPI 3.0
 *
 * API to store contamination events and data from different sources
 *
 * API version: 1.0.0
 * Contact: nestor.salvador@gdl.cinvestav.mx
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/api/v1/",
		Index,
	},

	{
		"CreateBundle",
		http.MethodPost,
		"/api/v1/bundle",
		CreateBundle,
	},

	{
		"GetBundleStatus",
		http.MethodGet,
		"/api/v1/bundle/:bundleId",
		GetBundleStatus,
	},

	{
		"CreateUser",
		http.MethodPost,
		"/api/v1/user",
		CreateUser,
	},

	{
		"CreateUsersWithListInput",
		http.MethodPost,
		"/api/v1/user/createWithList",
		CreateUsersWithListInput,
	},

	{
		"DeleteUser",
		http.MethodDelete,
		"/api/v1/user/:username",
		DeleteUser,
	},

	{
		"GetUserByName",
		http.MethodGet,
		"/api/v1/user/:username",
		GetUserByName,
	},

	{
		"LoginUser",
		http.MethodGet,
		"/api/v1/user/login",
		LoginUser,
	},

	{
		"LogoutUser",
		http.MethodGet,
		"/api/v1/user/logout",
		LogoutUser,
	},

	{
		"UpdateUser",
		http.MethodPut,
		"/api/v1/user/:username",
		UpdateUser,
	},
}
