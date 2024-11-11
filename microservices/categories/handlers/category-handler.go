package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/categories/models/dto"
	"github.com/synapsis-library-management-server/microservices/categories/utils/constant"
	"github.com/synapsis-library-management-server/microservices/categories/utils/response"
)

// CreateCategory
// @Summary Create category
// @Description This endpoint for create a category
// @Tags categories
// @Param request body dto.CreateCategoryRequest true "Required body to create categories"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/categories [post]
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get(constant.EmailHeader)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.CreateCategoryRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	request.Email = email
	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	msg, err := h.Service.CreateCategory(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusCreated, msg)
}

// GetCategoryById
// @Summary Get category by id
// @Description This endpoint for get a category id
// @Tags categories
// @Param id path string true "id of the category"
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/categories/{id} [get]
func (h *Handler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.GetCategoryByIdRequest
	request.Id = id
	result, err := h.Service.GetCategoryById(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithData(w, http.StatusOK, result)
}

// GetCategoriesByFilter
// @Summary Get categories
// @Description This endpoint for get list of categories
// @Tags categories
// @Param page query string false "Number of page"
// @Param pageSize query string false "Total data per Page"
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/categories [get]
func (h *Handler) GetCategoriesByFilter(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	var request dto.GetCategoriesByFilterRequest
	request.Page = int64(page)
	request.PageSize = int64(pageSize)
	err := request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, metadata, err := h.Service.GetCategoriesByFilter(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMetadata(w, http.StatusOK, result, metadata)
}
