package seeder

import (
	"libraryManagement/config"
	"libraryManagement/models"
	"log"
)

func SeederProduct() {
	dummyProducts := []models.Product{
		{IdCategory: 1, NameProduct: "Product A", DesciprtionProduct: "Description for Product A"},
		{IdCategory: 2, NameProduct: "Product B", DesciprtionProduct: "Description for Product B"},
		{IdCategory: 3, NameProduct: "Product C", DesciprtionProduct: "Description for Product C"},
		{IdCategory: 3, NameProduct: "Product D", DesciprtionProduct: "Description for Product D"},
	}

	for _, product := range dummyProducts {
		if err := config.DB.Create(&product).Error; err != nil {
			log.Printf("Error inserting product %s: %v", product.NameProduct, err)
		} else {
			log.Printf("Inserted product: %s", product.NameProduct)
		}
	}
}
