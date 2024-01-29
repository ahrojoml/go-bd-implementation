package internal

import "errors"

type Warehouse struct {
	ID        int
	Name      string
	Address   string
	Telephone string
	Capacity  int
}

type RepositoryWarehouse interface {
	GetByID(id int) (w Warehouse, err error)
	GetAll() (w []Warehouse, err error)
	Create(w *Warehouse) (err error)
}

var (
	// ErrWarehouseNotFound is an error that will be returned when a warehouse is not found
	ErrWarehouseNotFound = errors.New("repository: warehouse not found")
	// ErrWarehouseNotUnique is an error that will be returned when a warehouse is not unique
	ErrWarehouseNotUnique = errors.New("repository: warehouse not unique")
	// ErrWarehouseRelation is an error that will be returned when a warehouse relation fails
	ErrWarehouseRelation = errors.New("repository: warehouse relation error")
)
