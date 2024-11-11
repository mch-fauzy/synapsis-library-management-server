package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/synapsis-library-management-server/microservices/books/models/dto"
	"github.com/synapsis-library-management-server/microservices/books/utils/constant"
	"github.com/synapsis-library-management-server/microservices/books/utils/response"
)

// CreateBook
// @Summary Create book
// @Description This endpoint for create a book
// @Tags books
// @Param request body dto.CreateBookRequest true "Required body to create a book"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/books [post]
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get(constant.EmailHeader)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.CreateBookRequest
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

	msg, err := h.Service.CreateBook(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusCreated, msg)
}

// GetBooksByFilter
// @Summary Get books
// @Description This endpoint for get list of books
// @Tags books
// @Param page query string false "Number of page"
// @Param pageSize query string false "Total data per Page"
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/books [get]
func (h *Handler) GetBooksByFilter(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	var request dto.GetBooksByFilterRequest
	request.Page = int64(page)
	request.PageSize = int64(pageSize)
	err := request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, metadata, err := h.Service.GetBooksByFilter(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMetadata(w, http.StatusOK, result, metadata)
}
