package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PokemonController - interface for Pokemon requests
type PokemonController interface {
	GetPokemonByID(w http.ResponseWriter, r *http.Request)
	GetAllPokemons(w http.ResponseWriter, r *http.Request)
}

// Setup returns a router instance
func Setup(s PokemonController) (http.Handler, error) {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Sample endpoints
	v1.HandleFunc("/pokemons", s.GetAllPokemons).
		Methods(http.MethodGet).Name("GetAllPokemons")

	v1.HandleFunc("/pokemons/{id}", s.GetPokemonByID).
		Methods(http.MethodGet).Name("GetSinglePokemon")

	// Even Setup function never returns an error, is a good practice to have this set for in case
	// you decide to add more logic to the function and it's possible to return one.
	// It's up to you to have this or not
	return r, nil
}
