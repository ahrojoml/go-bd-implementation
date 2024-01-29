package internal

import (
	"errors"
	"time"
)

// Product is an struct that represents a product
type Product struct {
	// ID is the unique identifier of the product
	ID int
	// Name is the name of the product
	Name string
	// Quantity is the quantity of the product
	Quantity int
	// CodeValue is the code value of the product
	CodeValue string
	// IsPublished is the published status of the product
	IsPublished bool
	// Expiration is the expiration date of the product
	Expiration time.Time
	// Price is the price of the product
	Price float64
	// WarehouseID is the warehouse id of the product
	WarehouseID int
}

var (
	// ErrProductNotFound is an error that will be returned when a product is not found
	ErrProductNotFound = errors.New("repository: product not found")
	// ErrProductNotUnique is an error that will be returned when a product is not unique
	ErrProductNotUnique = errors.New("repository: product not unique")
	// ErrProductRelation is an error that will be returned when a product relation fails
	ErrProductRelation = errors.New("repository: product relation error")
)

// RepositoryProducts is an interface that represents a product repository
type RepositoryProducts interface {
	// GetOne returns a product by id
	GetOne(id int) (p Product, err error)
	GetAll() (p []Product, err error)
	// Store stores a product
	Store(p *Product) (err error)
	// Update updates a product
	Update(p *Product) (err error)
	// Delete deletes a product by id
	Delete(id int) (err error)
}
