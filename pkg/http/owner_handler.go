package http

import (
	"encoding/json"
	"gc-hexa-go/pkg/domain"
	"gc-hexa-go/pkg/ports/usecase"
	"gc-hexa-go/pkg/utils/customerror"
	"gc-hexa-go/pkg/utils/logger"
	"gc-hexa-go/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type OwnerHandler struct {
	OwnerUseCase usecase.OwnerUseCase
}

func NewOwnerHandler(ownerUseCase usecase.OwnerUseCase, r chi.Router) http.Handler {
	handler := &OwnerHandler{
		OwnerUseCase: ownerUseCase,
	}

	// owner routes
	r.Route("/owner", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetOwner)
		r.Post("/", handler.NewOwner)
		r.Put("/{id}", handler.UpdateOwner)
		r.Delete("/{id}", handler.DeleteOwner)
		r.Post("/{id}", handler.DeactivateOwner)
	})

	return r
}

func (h *OwnerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get owners data
	users, err := h.OwnerUseCase.FindAllOwners(ctx)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error())
		default:
			if err == customerror.ErrOwnersEmpty {
				response.WriteErrorResponse(w, http.StatusOK, err.Error())
			} else {
				response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			}
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, users)
}

func (h *OwnerHandler) GetOwner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse ID from URL params /{id}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed to parse Owner ID.")
	}

	// find owner
	owner, err := h.OwnerUseCase.FindOwnerById(id, ctx)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error())
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, owner)
}

func (h *OwnerHandler) NewOwner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody domain.Owner

	// parse user request body and map it to `owner` struct
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, customerror.ErrInvalidRequestBody.Error())
		return
	}

	newOwner, err := h.OwnerUseCase.CreateNewOwner(&reqBody, ctx)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error())
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusCreated, newOwner)
}

func (h *OwnerHandler) UpdateOwner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody domain.Owner

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	// parse request body and map it to `owner` struct
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	owner, err := h.OwnerUseCase.UpdateOwnerDetail(id, &reqBody, ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error())
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, owner)
}

func (h *OwnerHandler) DeleteOwner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.OwnerUseCase.RemoveOwner(id, ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, err.Error())
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, nil)
}

func (h *OwnerHandler) DeactivateOwner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.OwnerUseCase.DeactivateOwner(id, ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error())
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, nil)
}
