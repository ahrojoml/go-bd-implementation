package repository

import (
	"app/internal"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

const (
	GetWarehouseByID = "SELECT `id`, `name`, `address`, `telephone`, `capacity` FROM `warehouses` WHERE `id` = ?"
	GetAllWarehouses = "SELECT `id`, `name`, `address`, `telephone`, `capacity` FROM `warehouses`"
	CreateWarehouse  = "INSERT INTO `warehouses` (`name`, `address`, `telephone`, `capacity`) VALUES (?, ?, ?, ?)"
	GetProductCount  = "SELECT SUM(p.`quantity`) FROM `products` AS p WHERE `warehouse_id` = ?"
)

func NewWarehouseMySQL(db *sql.DB) *WarehouseMySQL {
	return &WarehouseMySQL{
		db: db,
	}
}

type WarehouseMySQL struct {
	db *sql.DB
}

func (r *WarehouseMySQL) GetByID(id int) (w internal.Warehouse, err error) {
	row := r.db.QueryRow(GetWarehouseByID, id)
	if err = row.Err(); err != nil {
		return internal.Warehouse{}, err
	}

	var warehouse internal.Warehouse
	if err := row.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.Warehouse{}, internal.ErrWarehouseNotFound
		}
		return internal.Warehouse{}, err
	}
	return warehouse, nil
}

func (r *WarehouseMySQL) GetAll() ([]internal.Warehouse, error) {
	rows, err := r.db.Query(GetAllWarehouses)
	if err != nil {
		return nil, err
	}

	var warehouses []internal.Warehouse = make([]internal.Warehouse, 0)
	for rows.Next() {
		var warehouse internal.Warehouse
		err = rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}

	fmt.Println(len(warehouses))

	return warehouses, nil
}

func (r *WarehouseMySQL) GetProductCount(id int) (int, error) {
	row := r.db.QueryRow(GetProductCount, id)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var count int
	if err := row.Scan(&count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}

func (r *WarehouseMySQL) Create(w *internal.Warehouse) (err error) {
	result, err := r.db.Exec(
		CreateWarehouse,
		w.Name,
		w.Address,
		w.Telephone,
		w.Capacity,
	)
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			return err
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	w.ID = int(id)

	return nil
}
