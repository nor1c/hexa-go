package http

import (
	"fmt"
	r "gc-hexa-go/internal/application/repository"
	u "gc-hexa-go/internal/application/usecase"
	db "gc-hexa-go/internal/database/mysql"
	"net/http"

)

func InitRoutes() {
	mux := http.NewServeMux()

	// repos
	ownerRepo := r.NewOwnerRepository(db.DB)
	// petRepo := r.NewPetRepository(db.DB)

	// usecases
	ownerUseCase := u.NewOwnerUseCase(ownerRepo)
	// petUseCase := u.NewPetUseCase(petRepo)

	// handlers
	ownerHandler := NewOwnerHandler(ownerUseCase, mux)
	// petHandler = NewPetHandler(petUseCase, mux)

	http.Handle("/owner/", http.StripPrefix("/owner", ownerHandler))

	// root/default
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
}
