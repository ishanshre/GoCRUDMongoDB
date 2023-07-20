package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/connections"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/helpers"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/models"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/repository"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/repository/dbrepos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	MG repository.DbMethod
}

func NewHandler(mg connections.DbInterface) Handlers {
	return &handler{
		MG: dbrepos.NewMongoDbRepo(mg, context.Background()),
	}
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	products, err := h.MG.GetProducts(limit, page)
	if err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusOKAll(w, limit, page, products)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := &models.CreateUpdateProduct{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		helpers.StatusBadRequest(w, "error parsing json")
		return
	}
	product, err := h.MG.CreateProduct(p)
	if err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusCreated(w, product)
}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		helpers.StatusBadRequest(w, "invalid id")
		return
	}
	product, err := h.MG.GetProduct(id)
	if err != nil {
		helpers.StatusNotFound(w, err.Error())
		return
	}
	helpers.StatusOK(w, product)
}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		helpers.StatusBadRequest(w, "invalid id")
		return
	}
	if err := h.MG.DeleteProduct(id); err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusAcceptedMsg(w, "product deleted")
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		helpers.StatusBadRequest(w, "invalid id")
		return
	}
	update := &models.CreateUpdateProduct{}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		helpers.StatusBadRequest(w, "error in parsing json")
		return
	}
	if err := h.MG.UpdateProduct(id, update); err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	product, err := h.MG.GetProduct(id)
	if err != nil {
		helpers.StatusNotFound(w, err.Error())
		return
	}
	helpers.StatusAcceptedData(w, product)
}
