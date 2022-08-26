package main

import (
	rProduct "test-agrak/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	rProduct.RoutesProduct(router)
	router.Run()
}
