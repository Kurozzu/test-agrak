package routes

import (
	"log"

	cProduct "test-agrak/app/controllers"
	// mProduct "test-agrak/app/models"
	response "test-agrak/app/utils"

	"github.com/gin-gonic/gin"
)

var prod cProduct.Prod

//RoutesProduct set of routes about product for CRUD
func RoutesProduct(router *gin.Engine) {
	router.GET("/product", func(ctx *gin.Context) {
		resp, err := prod.GetAllProduct()
		if err != nil {
			response.JSONresponse(ctx, 400, err.Error())
			return
		}
		response.JSONresponse(ctx, 200, resp)
	})
	router.GET("/product/:sku", func(ctx *gin.Context) {
		sku := ctx.Param("sku")
		resp, err := prod.GetProduct(sku)
		if err != nil {
			response.JSONresponse(ctx, 400, err.Error())
			return
		}
		response.JSONresponse(ctx, 200, resp)
		// response.JSONresponse(ctx, 200, fmt.Sprintf("Obtener el producto %s", id))
	})
	router.POST("/product", func(ctx *gin.Context) {
		if err := ctx.ShouldBind(&prod); err != nil {
			response.JSONresponse(ctx, 400, "Formato de entrada es incorrecto")
			log.Fatal(err)
		}
		resp, err := prod.InsertProduct()
		if err != nil {
			response.JSONresponse(ctx, 400, err.Error())
			return
		}
		response.JSONresponse(ctx, 201, resp)
	})
	router.PUT("/product", func(ctx *gin.Context) {
		if err := ctx.ShouldBind(&prod); err != nil {
			response.JSONresponse(ctx, 400, "Formato de entrada es incorrecto")
			log.Fatal(err)
			return
		}
		resp, err := prod.UpdateProduct()
		if err != nil {
			response.JSONresponse(ctx, 400, err.Error())
			return
		}
		response.JSONresponse(ctx, 200, resp)
	})
	router.DELETE("/product/:sku", func(ctx *gin.Context) {
		sku := ctx.Param("sku")
		resp, err := prod.DeleteProduct(sku)
		if err != nil {
			response.JSONresponse(ctx, 400, err.Error())
			return
		}
		response.JSONresponse(ctx, 200, resp)
	})
}
