package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/dto"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/entity"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/infra/database"
	entityPkg "github.com/LucasBelusso1/GoExpert/7-APIS/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// CreateProduct godoc
// @Summary		Create product
// @Description	Create products
// @tags		products
// @Accept		json
// @Produce		json
// @Param		request body dto.CreateProductInput true "product request"
// @Success		201
// @Failure		500 {object} Error
// @Router		/products [post]
// @Security	ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product, err := entity.NewProduct(productDto.Name, productDto.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Create(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProducts godoc
// @Summary		List products
// @Description	Get all products
// @tags		products
// @Accept		json
// @Produce		json
// @Param		page query string false "page number"
// @Param		limit query string false "page limit"
// @Success		200 {array} entity.Product
// @Failure		404 {object} Error
// @Failure		500 {object} Error
// @Router		/products [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProduct godoc
// @Summary		Get a product
// @Description	Get a product
// @tags		products
// @Accept		json
// @Produce		json
// @Param		id path string true "product ID" Format(uuid)
// @Success		200 {object} entity.Product
// @Failure		404 {object} Error
// @Failure		500 {object} Error
// @Router		/products/{id} [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Product with given ID not found"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err := entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary		Update a product
// @Description	Update a product
// @tags		products
// @Accept		json
// @Produce		json
// @Param		id path string true "product ID" Format(uuid)
// @Param		request body dto.CreateProductInput true "product request"
// @Success		200
// @Failure		404 {object} Error
// @Failure		500 {object} Error
// @Router		/products/{id} [put]
// @Security	ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Product with given ID not found"}
		json.NewEncoder(w).Encode(error)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product.ID, err = entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err = h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary		Delete a product
// @Description	Delete a product
// @tags		products
// @Accept		json
// @Produce		json
// @Param		id path string true "product ID" Format(uuid)
// @Success		200
// @Failure		404 {object} Error
// @Failure		500 {object} Error
// @Router		/products/{id} [delete]
// @Security	ApiKeyAuth
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Product with given ID not found"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err := entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err = h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}
