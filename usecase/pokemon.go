package usecase

import (
	"fmt"

	"github.com/RolloCasanova/sample-clean/model"
	"github.com/sirupsen/logrus"
)

// PokemonService interface defines methods expected to be used as a service
type PokemonService interface {
	GetAllPokemons() ([]model.Pokemon, error)
	GetPokemonByID(id int) (*model.Pokemon, error)
}

// PokemonUsecase defines usecase fields
type PokemonUsecase struct {
	service PokemonService
	log     *logrus.Logger
}

// NewPokemonUsecase creates a usecase for Pokemon
func NewPokemonUsecase(ps PokemonService, l *logrus.Logger) (*PokemonUsecase, error) {
	l.Debugln("in usecase.NewPokemonUsecase")
	p := &PokemonUsecase{
		service: ps,
		log:     l,
	}

	l.Debugln("NewPokemonUsecase successfully created")

	// We are not returning an error (yet), you could leave it as it is, or remove error return.
	return p, nil
}

// GetAllPokemons calls service to retrieve information about all Pokemons
func (p PokemonUsecase) GetAllPokemons() ([]model.Pokemon, error) {
	p.log.Debugln("in usecase.GetAllPokemons")

	pokemons, err := p.service.GetAllPokemons()
	if err != nil {
		return nil, fmt.Errorf("unable to get all pokemons: %w", err)
	}

	return pokemons, nil
}

// GetPokemonByID calls service to retrieve information about a specific pokemon with an specific ID
func (p PokemonUsecase) GetPokemonByID(id int) (*model.Pokemon, error) {
	p.log.WithField("ID", id).Debugln("in usecase.GetPokemonByID")

	pokemon, err := p.service.GetPokemonByID(id)
	if err != nil {
		p.log.WithError(err).WithField("ID", id).Error("unable to get pokemon by ID")

		return nil, fmt.Errorf("unable to get pokemon with ID [%d]: %w", id, err)
	}

	return pokemon, nil
}
