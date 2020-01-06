package main

import (
	"institute/controllers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type Routes []Route

var publicRoutes = Routes{
	Route{"SignUpNewUser", "POST", "/signup", controllers.SignUpNewUser},
	Route{"LogInUser", "POST", "/login", controllers.LogInUser},
	Route{"ConfirmationUserAccount", "GET", "/confirmation/:code", controllers.ConfirmationUserAccount},
	Route{"UpdateUser", "PUT", "/update-user", controllers.UpdateUser},
}

func NewRouter() {
	router := gin.Default()

	/* public routes */
	public := router.Group("/api/v1")
	for _, route := range publicRoutes {
		switch route.Method {
		case "GET":
			public.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			public.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			public.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			public.DELETE(route.Pattern, route.HandlerFunc)
		default:
			public.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}

	// public.StaticFile("/giftcard.png", "./utils/assets/images/giftcard.png")
	// router.NoRoute(controllers.UnauthorizedAccessResponse)
	router.Run(":8080")
	// router.RunTLS(":8080", "./certs/64bd560e702a84bb.crt", "./certs/star.bookingkoala.key")
}
