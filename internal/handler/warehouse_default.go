package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewWarehouseDefault(rw internal.RepositoryWarehouse) *WarehouseDefault {
	return &WarehouseDefault{
		rw: rw,
	}
}

type WarehouseDefault struct {
	rw internal.RepositoryWarehouse
}

type WarehouseJSON struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

func (h *WarehouseDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
		}

		warehouse, err := h.rw.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := WarehouseJSON{
			ID:        warehouse.ID,
			Name:      warehouse.Name,
			Address:   warehouse.Address,
			Telephone: warehouse.Telephone,
			Capacity:  warehouse.Capacity,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *WarehouseDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.rw.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}

		data := make([]WarehouseJSON, len(warehouses))
		for _, warehouse := range warehouses {
			data = append(data, WarehouseJSON{
				ID:        warehouse.ID,
				Name:      warehouse.Name,
				Address:   warehouse.Address,
				Telephone: warehouse.Telephone,
				Capacity:  warehouse.Capacity,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *WarehouseDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody WarehouseJSON
		if err := request.JSON(r, &reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, "could not read body")
			return
		}

		warehouse := internal.Warehouse{
			Name:      reqBody.Name,
			Address:   reqBody.Address,
			Telephone: reqBody.Telephone,
			Capacity:  reqBody.Capacity,
		}
		err := h.rw.Create(&warehouse)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseNotUnique):
				response.Error(w, http.StatusBadRequest, "warehouse not unique")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := WarehouseJSON{
			ID:        warehouse.ID,
			Name:      warehouse.Name,
			Address:   warehouse.Address,
			Telephone: warehouse.Telephone,
			Capacity:  warehouse.Capacity,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

// Im lazy and dont want to define my service layer right now
func (h *WarehouseDefault) ProductReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
