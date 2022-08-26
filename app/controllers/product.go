package controllers

import (
	"errors"
	"fmt"
	mProduct "test-agrak/app/models"
	conn "test-agrak/libs"

	"gorm.io/gorm"
)

type Prod mProduct.Product

var prods []mProduct.Product
var Con *gorm.DB = conn.InitMySql()
var Create = Con.Create
var CustomUp = CustomUpdate
var CustomDel = CustomDelete
var CustomAll = CustomAllproduct
var CustomGet = CustomGetProduct

//GetAllProduct allows to get all the products on a list sorted by sku asc
func (p *Prod) GetAllProduct() ([]mProduct.Product, error) {
	if CustomAll(p) == true {
		return nil, errors.New("No se han encontrados datos relacioandos")
	}
	return prods, nil
}

func CustomAllproduct(p *Prod) bool {
	Con.Model(&p).Order("CAST(sku AS UNSIGNED) ASC").Find(&prods)
	if len(prods) == 0 {
		return true
	} else {
		return false
	}
}

//GetProduct allows to get a specific product by sku
func (p *Prod) GetProduct(sku string) (prod mProduct.Product, err error) {
	tx := Con.Model(&p).Where("sku = ?", sku).Find(&prod)
	if CustomGet(tx) == 0 {
		return prod, errors.New("No se han encontrados datos relacioandos")
	}
	fmt.Println(prod)
	return prod, nil
}

func CustomGetProduct(p *gorm.DB) int64 {
	return p.RowsAffected
}

//InsertProduct allows to create or insert a new register into the data base
func (p *Prod) InsertProduct() (string, error) {
	tx := Create(&p)
	if tx.Error != nil {
		return "", errors.New("Error al tratar de registrar el producto")
	}
	return "Producto registrado correctamente", nil
}

//UpdateProduct allows to update one register by sku
func (p *Prod) UpdateProduct() (string, error) {
	if CustomUp(p) == 0 {
		return "", errors.New("Error al tratar de actualizar el producto")
	}
	return "Producto actualizado exitosamente", nil
}

func CustomUpdate(p *Prod) int64 {
	return Con.Model(&p).Where("sku = ?", p.Sku).Updates(Prod{
		Name:       p.Name,
		Brand:      p.Brand,
		Size:       p.Size,
		Price:      p.Price,
		Mainimage:  p.Mainimage,
		Otherimage: p.Otherimage,
	}).RowsAffected
}

//DeleteProduct allows to delete one register by sku
func (p *Prod) DeleteProduct(sku string) (string, error) {
	if CustomDel(sku, p) == 0 {
		return "", errors.New(fmt.Sprintf("No existe el registro %s", sku))
	}
	return "Registro eliminado exitosamente", nil
}

func CustomDelete(sku string, p *Prod) int64 {
	return Con.Debug().Where("sku = ?", sku).Delete(&p).RowsAffected
}
