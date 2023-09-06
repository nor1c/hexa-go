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

type PetHandler struct {
	PetUseCase usecase.PetUseCase
}

func NewPetHandler(petUseCase usecase.PetUseCase, r chi.Router) http.Handler {
	handler := &PetHandler{
		PetUseCase: petUseCase,
	}

	r.Route("/pet", func(r chi.Router) {
		r.Get("/", handler.GetAllPets)
		r.Get("/{id}", handler.GetPetDetail)
		r.Post("/", handler.AddNewPet)
		r.Put("/{id}", handler.UpdatePet)
		r.Delete("/{id}", handler.DeletePet)
	})

	return r
}

func (h *PetHandler) GetAllPets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pets, err := h.PetUseCase.GetAllPets(ctx)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error(), err)
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, pets)
}

func (h *PetHandler) GetPetDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	petID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	pet, err := h.PetUseCase.GetPetDetail(petID, ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error(), err)
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, pet)
}

func (h *PetHandler) AddNewPet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody *domain.Pet

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, customerror.ErrInvalidRequestBody.Error(), err)
		return
	}

	pet, err := h.PetUseCase.AddNewPet(reqBody, ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error(), err)
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusCreated, pet)
}

func (h *PetHandler) UpdatePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse petID from URL param
	petID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	// parse request body
	var reqBody *domain.Pet
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, customerror.ErrInvalidRequestBody.Error(), err)
		return
	}

	err = h.PetUseCase.UpdatePet(reqBody, petID, ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, customerror.ErrTimedOut.Error(), err)
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, nil)
}

func (h *PetHandler) DeletePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	petID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err = h.PetUseCase.DeletePet(petID, ctx); err != nil {
		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusOK, customerror.ErrTimedOut.Error(), err)
		default:
			response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, nil)
}
