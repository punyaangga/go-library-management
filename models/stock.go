package models

import (
	"errors"
	"libraryManagement/config"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	IdProduct     int    `json:"IdProduct" binding:"required";gorm:"unique;not null"`
	StockBarcode  string `json:"StockBarcode" binding:"required";gorm:"unique;not null"`
	StockQty      int    `json:"StockQty" binding:"required";gorm:"unique;not null"`
	StockLocation string `gorm:"type:text; null"`
}

type StockResponse struct {
	StockBarcode  string `json:"StockBarcode"`
	StockQty      int    `json:"StockQty"`
	StockLocation string `json:"Description"`
}

type StockWithProductResponse struct {
	ID            uint   `json:"id"`
	NameProduct   string `json:"name_product"`
	StockBarcode  string `json:"stock_barcode"`
	StockQty      int    `json:"stock_qty"`
	StockLocation string `json:"stock_location"`
}

// Method to check if product exists
func (p *Product) ProductExists(idProduct int) bool {

	var product Product
	// Try to find the product by id
	if err := config.DB.First(&product, idProduct).Error; err != nil {
		return false // Product not found
	}
	return true // Product exists
}

// method get detail stock one or more data
func GetStocksWithProducts(db *gorm.DB, id *string, productName *string) ([]StockWithProductResponse, error) {
	var stocks []StockWithProductResponse
	var err error

	// Query dasar
	query := `
		SELECT
			a.id,
			a.stock_barcode,
			a.stock_qty,
			a.stock_location,
			b.name_product
		FROM stocks AS a
		LEFT JOIN products AS b
		ON a.id_product = b.id
	`

	// Parameter untuk query
	var queryParams []interface{}

	// Tambahkan kondisi untuk ID
	if id != nil && *id != "" {
		query += " WHERE a.id = ?"
		queryParams = append(queryParams, *id)
	}

	// Tambahkan kondisi untuk nama produk
	if productName != nil && *productName != "" {
		if len(queryParams) > 0 {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " b.name_product LIKE ?"
		queryParams = append(queryParams, "%"+*productName+"%")
	}

	// Eksekusi query
	err = db.Raw(query, queryParams...).Scan(&stocks).Error

	if err != nil {
		return nil, err
	}

	// Pengecekan jika data kosong
	if len(stocks) == 0 {
		if id != nil && *id != "" {
			return nil, errors.New("Stock ID not found")
		}
		if productName != nil && *productName != "" {
			return nil, errors.New("No stocks available for the given product name")
		}
		return nil, errors.New("No stocks available")
	}

	return stocks, nil
}
