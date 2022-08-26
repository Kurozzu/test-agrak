package controllers

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

var prod Prod

func TestGetAllProduct(t *testing.T) {
	t.Run("GetAllProductSuccess", func(t *testing.T) {
		CustomAll = func(p *Prod) bool {
			return false
		}
		_, err := prod.GetAllProduct()
		if err != nil {
			t.Errorf("Got %v - Want %v", err, "nil")
		}
	})
	t.Run("GetAllProductError", func(t *testing.T) {
		CustomAll = func(p *Prod) bool {
			return true
		}
		_, err := prod.GetAllProduct()
		if err == nil {
			t.Errorf("Got %v - Want %v", err, "No se han encontrados datos relacioandos")
		}
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("GetproducSuccess", func(t *testing.T) {
		CustomGet = func(p *gorm.DB) int64 {
			return 1
		}
		_, err := prod.GetProduct("FAL-8406271")
		if err != nil {
			t.Errorf("Got %v - Want %v", err, "nil")
		}
	})
	t.Run("GetproducError", func(t *testing.T) {
		CustomGet = func(p *gorm.DB) int64 {
			return 0
		}
		_, err := prod.GetProduct("FAL-8406271")
		if err == nil {
			t.Errorf("Got %v - Want %v", err, "No se han encontrados datos relacioandos")
		}
	})
}

func TestInsertProduct(t *testing.T) {
	t.Run("InsertProductSuccess", func(t *testing.T) {
		Create = func(value interface{}) (tx *gorm.DB) {
			return Con
		}
		p := Prod{
			Sku:        "FAL-8406271",
			Name:       "501 Zapatilla Urbana Mujer",
			Brand:      "New Balance 1",
			Size:       "37",
			Price:      52990.00,
			Mainimage:  "https://ruta update 1",
			Otherimage: "https://other 1",
		}
		_, err := p.InsertProduct()
		if err != nil {
			t.Errorf("Got %v - Want %v", err, "nil")
		}
	})
	t.Run("InsertProductError", func(t *testing.T) {
		Create = func(value interface{}) (tx *gorm.DB) {
			Con.Error = errors.New("custom error")
			return Con
		}
		p := Prod{
			Sku:        "FAL-8406271",
			Name:       "501 Zapatilla Urbana Mujer",
			Brand:      "New Balance 1",
			Size:       "37",
			Price:      52990.00,
			Mainimage:  "https://ruta update 1",
			Otherimage: "https://other 1",
		}
		_, err := p.InsertProduct()
		if err == nil {
			t.Errorf("Got %v - Want %v", err, "Error al tratar de registrar el producto")
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("UpdateProductSuccess", func(t *testing.T) {
		CustomUp = func(p *Prod) int64 {
			return 1
		}
		_, err := prod.UpdateProduct()
		if err != nil {
			t.Errorf("Got %v - Want %v", err, "nil")
		}
	})
	t.Run("UpdateProductError", func(t *testing.T) {
		CustomUp = func(p *Prod) int64 {
			return 0
		}
		_, err := prod.UpdateProduct()
		if err == nil {
			t.Errorf("Got %v - Want %v", err, "Error al tratar de actualizar el producto")
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("DeleteProductSuccess", func(t *testing.T) {
		CustomDel = func(sku string, p *Prod) int64 {
			return 1
		}
		_, err := prod.DeleteProduct("FAL-8406271")
		if err != nil {
			t.Errorf("Got %v - Want %v", err, "nil")
		}
	})
	t.Run("DeleteProductError", func(t *testing.T) {
		CustomDel = func(sku string, p *Prod) int64 {
			return 0
		}
		_, err := prod.DeleteProduct("FAL-8406271")
		if err == nil {
			t.Errorf("Got %v - Want %v", err, "nil")
		}
	})
}
