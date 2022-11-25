package route

import (
	"github.com/helmimuzkr/belajar-golang-restapi/controller"
	"github.com/julienschmidt/httprouter"
)

type categoryRouter struct {
	router     *httprouter.Router
	controller controller.CategoryController
}

func NewCategoryRouter(router *httprouter.Router, controller controller.CategoryController) *categoryRouter {
	return &categoryRouter{
		router:     router,
		controller: controller,
	}
}

func (cr *categoryRouter) CategoryRouter() *httprouter.Router {
	cr.router.POST("/api/categories/", cr.controller.CreateCategory)
	cr.router.PUT("/api/categories/:id", cr.controller.UpdateCategory)
	cr.router.DELETE("/api/categories/:id", cr.controller.DeleteCategory)
	cr.router.GET("/api/categories/", cr.controller.GetAllCategory)
	cr.router.GET("/api/categories/:id", cr.controller.GetCategory)

	return cr.router
}
