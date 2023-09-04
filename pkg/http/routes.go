package http

import (
	repo "gc-hexa-go/internal/application/repository"
	uc "gc-hexa-go/internal/application/usecase"
	db "gc-hexa-go/internal/database/mysql"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(r chi.Router) chi.Router {
	// repos
	ownerRepo := repo.NewOwnerRepository(db.DB)
	// petRepo := r.NewPetRepository(db.DB)

	// usecases
	ownerUseCase := uc.NewOwnerUseCase(ownerRepo)
	// petUseCase := u.NewPetUseCase(petRepo)

	// handlers
	NewOwnerHandler(ownerUseCase, r)
	// NewPetHandler(petUseCase, r)

	return r
}
