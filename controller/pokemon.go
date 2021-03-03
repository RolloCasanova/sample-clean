package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RolloCasanova/sample-clean/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// PokemonUsecase interface defines methods expected to be used as a usecase
type PokemonUsecase interface {
	GetAllPokemons() ([]model.Pokemon, error)
	GetPokemonByID(id int) (*model.Pokemon, error)
}

// PokemonController defines controller fields
type PokemonController struct {
	render  *render.Render
	usecase PokemonUsecase
	log     *logrus.Logger
}

// NewPokemonController creates a Controller for Pokemon
func NewPokemonController(r *render.Render, u PokemonUsecase, l *logrus.Logger) (PokemonController, error) {
	p := PokemonController{
		render:  r,
		usecase: u,
		log:     l,
	}

	// We are not returning an error (yet), you could leave it as it is, or remove error return.
	return p, nil
}

// GetAllPokemons handles request to fetch all Pokemons
func (p PokemonController) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	p.log.Debugln("in controller.GetAllPokemons")

	pokemons, err := p.usecase.GetAllPokemons()
	if err != nil {
		err = fmt.Errorf("usecase request failed: %w", err)
		p.render.Text(w, http.StatusBadRequest, err.Error())

		return
	}

	p.log.Infoln("Pokemons retrieved successfully")

	p.render.JSON(w, http.StatusOK, pokemons)
}

// GetPokemonByID handles request to get Pokemon by it's ID
func (p PokemonController) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	p.log.Debugln("in controller.GetPokemonByID")

	idStr := mux.Vars(r)["id"]

	if idStr == "" {
		p.render.Text(w, http.StatusBadRequest, "controller: path param {id} must be not empty")

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		p.render.Text(w, http.StatusBadRequest, "controller: path param {id} must be an integer")

		return
	}

	pokemon, err := p.usecase.GetPokemonByID(id)
	if err != nil {
		err = fmt.Errorf("usecase request failed: %w", err)
		p.render.Text(w, http.StatusBadRequest, err.Error())

		return
	}

	p.log.Infof("Pokemon with ID [%d] successfully retrieved\n", id)

	p.render.JSON(w, http.StatusOK, pokemon)
}
