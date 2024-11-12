package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/synapsis-library-management-server/microservices/borrows/models/dto"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/constant"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/response"
)

// CreateBorrow
// @Summary Create borrow
// @Description This endpoint for create a borrow record
// @Tags borrows
// @Param request body dto.CreateBorrowRequest true "Required body to create borrow record"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/borrows [post]
func (h *Handler) CreateBorrow(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get(constant.EmailHeader)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.CreateBorrowRequest
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

	msg, err := h.Service.CreateBorrow(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusCreated, msg)
}

// GetBorrowsByFilter
// @Summary Get borrows
// @Description This endpoint for get list of borrows record
// @Tags borrows
// @Param page query string false "Number of page"
// @Param pageSize query string false "Total data per Page"
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/borrows [get]
func (h *Handler) GetBorrowsByFilter(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	var request dto.GetBorrowsByFilterRequest
	request.Page = int64(page)
	request.PageSize = int64(pageSize)
	err := request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, metadata, err := h.Service.GetBorrowsByFilter(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMetadata(w, http.StatusOK, result, metadata)
}
