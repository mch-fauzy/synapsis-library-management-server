package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/synapsis-library-management-server/microservices/authors/models/dto"
	"github.com/synapsis-library-management-server/microservices/authors/utils/constant"
	"github.com/synapsis-library-management-server/microservices/authors/utils/response"
)

// CreateAuthor
// @Summary Create Author
// @Description This endpoint for create an author
// @Tags authors
// @Param request body dto.CreateAuthorRequest true "Required body to create author"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/authors [post]
func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get(constant.EmailHeader)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.CreateAuthorRequest
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

	msg, err := h.Service.CreateAuthor(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusCreated, msg)
}

// GetAuthorsByFilter
// @Summary Get Authors
// @Description This endpoint for get list of author
// @Tags authors
// @Param page query string false "Number of page"
// @Param pageSize query string false "Total data per Page"
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/authors [get]
func (h *Handler) GetAuthorsByFilter(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	var request dto.GetAuthorsByFilterRequest
	request.Page = int64(page)
	request.PageSize = int64(pageSize)
	err := request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, metadata, err := h.Service.GetAuthorsByFilter(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMetadata(w, http.StatusOK, result, metadata)
}
