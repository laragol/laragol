package Routes

import (
	. "app/Http/Controllers"

	_routing "github.com/qiangxue/fasthttp-routing"
)

func Handle(router *_routing.Router) {
	api := router.Group("/api")

	api.Get("/users", UserController{}.Index)
	api.Get("/users/<userId>", UserController{}.Show)
	api.Post("/users", UserController{}.Store)
	api.Put("/users/<userId>", UserController{}.Update)

	api.Get("/countries", CountryController{}.Index)
	api.Get("/countries/<countryId>", CountryController{}.Show)
	api.Post("/countries", CountryController{}.Store)
	api.Put("/countries/<countryId>", CountryController{}.Update)
}
