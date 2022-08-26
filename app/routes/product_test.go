package routes

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	cProduct "test-agrak/app/controllers"
	mProduct "test-agrak/app/models"
	response "test-agrak/app/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
)

func ConfigRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestGetAllProduct(t *testing.T) {
	router := ConfigRoutes()
	router.GET("/product", func(ctx *gin.Context) { prod.GetAllProduct() })
	t.Run("GetAllProductSuccess", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/product", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetProduct(t *testing.T) {
	router := ConfigRoutes()
	router.GET("/product/:sku", func(ctx *gin.Context) { sku := ctx.Param("sku"); prod.GetProduct(sku) })
	t.Run("GetProduct", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/product/FAL-8406271", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestInsertProduct(t *testing.T) {
	router := ConfigRoutes()
	router.POST("/product", func(ctx *gin.Context) {
		if err := ctx.ShouldBind(&prod); err != nil {
			response.JSONresponse(ctx, 400, "Formato de entrada es incorrecto")
			log.Println(err)
			return
		}
		prod.InsertProduct()
	})
	t.Run("PostProductSuccess", func(t *testing.T) {
		cProduct.Create = func(value interface{}) (tx *gorm.DB) {
			return cProduct.Con
		}
		p := mProduct.Product{
			Sku:        "FAL-8406271",
			Name:       "501 Zapatilla Urbana Mujer",
			Brand:      "New Balance 1",
			Size:       "37",
			Price:      52990.00,
			Mainimage:  "https://ruta update 1",
			Otherimage: "https://other 1",
		}
		jsonValue, _ := json.Marshal(p)
		req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("PostProductError", func(t *testing.T) {
		cProduct.Create = func(value interface{}) (tx *gorm.DB) {
			cProduct.Con.Error = nil
			return cProduct.Con
		}
		p := mProduct.Product{
			Sku:        "FAL-8406271",
			Name:       "501 Zapatilla Urbana Mujer",
			Brand:      "New Balance 1",
			Size:       "37",
			Price:      52990.00,
			Mainimage:  "https://ruta update 1",
			Otherimage: "https://other 1",
		}
		jsonValue, _ := json.Marshal(p)
		req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUpdateProduct(t *testing.T) {
	router := ConfigRoutes()
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
	t.Run("PutProductSuccess", func(t *testing.T) {
		cProduct.CustomUp = func(p *cProduct.Prod) int64 {
			return 1
		}
		p := mProduct.Product{
			Sku:        "FAL-8406271",
			Name:       "501 Zapatilla Urbana Mujer",
			Brand:      "New Balance 1",
			Size:       "37",
			Price:      52990.00,
			Mainimage:  "https://ruta update 1",
			Otherimage: "https://other 1",
		}
		jsonValue, _ := json.Marshal(p)
		req, _ := http.NewRequest("PUT", "/product", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("PutProductError", func(t *testing.T) {
		cProduct.CustomUp = func(p *cProduct.Prod) int64 {
			return 0
		}
		p := mProduct.Product{
			Sku:        "FAL-8406271",
			Name:       "501 Zapatilla Urbana Mujer",
			Brand:      "New Balance 1",
			Size:       "37",
			Price:      52990.00,
			Mainimage:  "https://ruta update 1",
			Otherimage: "https://other 1",
		}
		jsonValue, _ := json.Marshal(p)
		req, _ := http.NewRequest("PUT", "/product", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteProduct(t *testing.T) {
	router := ConfigRoutes()
	router.DELETE("/product/:sku", func(ctx *gin.Context) {
		sku := ctx.Param("sku")
		resp, err := prod.DeleteProduct(sku)
		if err != nil {
			response.JSONresponse(ctx, 400, err.Error())
			return
		}
		response.JSONresponse(ctx, 200, resp)
	})
	t.Run("DeleteProductoSuccess", func(t *testing.T) {
		cProduct.CustomDel = func(sku string, p *cProduct.Prod) int64 {
			return 1
		}
		req, _ := http.NewRequest("DELETE", "/product/FAL-8406271", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("DeleteProductoError", func(t *testing.T) {
		cProduct.CustomDel = func(sku string, p *cProduct.Prod) int64 {
			return 0
		}
		req, _ := http.NewRequest("DELETE", "/product/FAL-8406271", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
