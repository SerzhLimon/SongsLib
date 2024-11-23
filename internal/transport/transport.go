package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc gin.HandlerFunc
}

func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
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

func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {
	DefaultAPI DefaultAPI
}

type DefaultAPI struct {
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"UsersGet",
			http.MethodGet,
			"/users",
			handleFunctions.DefaultAPI.GetUsers,
		},
	}
}

func (api *DefaultAPI) GetUsers(c *gin.Context) {

	c.Status(http.StatusOK)
}
