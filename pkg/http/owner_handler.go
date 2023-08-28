package http

import (
	"context"
	"fmt"
	"gc-hexa-go/pkg/ports/usecase"
	"gc-hexa-go/pkg/utils/customerror"
	"gc-hexa-go/pkg/utils/logger"
	"gc-hexa-go/pkg/utils/response"
	"net/http"
	"time"
)

type OwnerHandler struct {
	OwnerUseCase usecase.OwnerUseCase
}

func NewOwnerHandler(ownerUseCase usecase.OwnerUseCase, mux *http.ServeMux) http.Handler {
	handler := &OwnerHandler{
		OwnerUseCase: ownerUseCase,
	}

	// owner routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Owner routes!")
	})
	mux.HandleFunc("/all", handler.GetAll)

	return mux
}

func (h *OwnerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	users, err := h.OwnerUseCase.FindAllOwners(ctx)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())

		select {
		case <-ctx.Done():
			response.WriteErrorResponse(w, http.StatusRequestTimeout, "Request timed out.", err)
		default:
			if err == customerror.ErrOwnersEmpty {
				response.WriteErrorResponse(w, http.StatusOK, err.Error(), err)
			} else {
				response.WriteErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
			}
		}

		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, users)
}
