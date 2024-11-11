package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/synapsis-library-management-server/microservices/users/models/dto"
	"github.com/synapsis-library-management-server/microservices/users/utils/constant"
	"github.com/synapsis-library-management-server/microservices/users/utils/response"
)

// RegisterUser
// @Summary Register user
// @Description This endpoint for register an user
// @Tags register
// @Param request body dto.RegisterRequest true "Required body to register user"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/register [post]
func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.RegisterRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	request.Role = constant.RoleUser
	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	msg, err := h.Service.Register(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusCreated, msg)
}

// RegisterAdmin
// @Summary Register admin
// @Description This endpoint for register an admin
// @Tags admin register
// @Param request body dto.RegisterRequest true "Required body to register admin"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/admin/register [post]
func (h *Handler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.RegisterRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	request.Role = constant.RoleAdmin
	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	msg, err := h.Service.Register(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithMessage(w, http.StatusCreated, msg)
}

// Login
// @Summary Login
// @Description This endpoint for login
// @Tags login
// @Param request body dto.LoginRequest true "Required body to login user"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.LoginRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.Service.Login(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithData(w, http.StatusOK, result)
}
