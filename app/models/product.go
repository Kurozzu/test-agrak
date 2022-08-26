package models

import (
	"time"
)

//Product is the structure of product
type Product struct {
	Sku        string  `json:"sku" gorm:"type:varchar(255);not null;primaryKey" binding:"required"`
	Name       string  `json:"name" gorm:"type:varchar(255);not null" binding:"required"`
	Brand      string  `json:"brand" gorm:"type:varchar(255);not null" binding:"required"`
	Size       string  `json:"size, omitempty" gorm:"type:varchar(255)" `
	Price      float64 `json:"price" gorm:"type:decimal(10,3);not null" binding:"required"`
	Mainimage  string  `json:"mainImage" gorm:"type:varchar(255);not null" binding:"required"`
	Otherimage string  `json:"otherImage, omitempty" gorm:"type:varchar(255)" binding:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	//If descomment this line then "delete" sql sentences doesn't remove the row, instead update the row and change the delete date
	//Recommend use this with a status column 'cause product does not remove.
	// DeletedAt  gorm.DeletedAt `gorm:"index"`

}
