package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/RolloCasanova/sample-clean/model"
	"github.com/sirupsen/logrus"
)

// PokemonPostgreSQLService defines service fields
type PokemonPostgreSQLService struct {
	db  *sql.DB
	log *logrus.Logger
}

// NewPokemonService creates a PostgreSQL connection with specified values
func NewPokemonService(
	name, user, password, host string,
	port int,
	l *logrus.Logger,
) (PokemonPostgreSQLService, error) {
	// Not implemented yet. You can do it ;)

	ps := PokemonPostgreSQLService{
		db:  nil,
		log: l,
	}

	return ps, nil
}

// GetAllPokemons queries PostgreSQL to fetch all pokemons on DB
func (p PokemonPostgreSQLService) GetAllPokemons() ([]model.Pokemon, error) {
	p.log.Debugln("in service.GetAllPokemons")

	return nil, errors.New("method GetAllPokemons not yet implemented")
}

// GetPokemonByID queries PostgreSQL to get an specific pokemon by it's ID
func (p PokemonPostgreSQLService) GetPokemonByID(id int) (*model.Pokemon, error) {
	p.log.Debugln("in service.GetPokemonByID")

	// should call function to query in DB by using
	// pokemon, err := queryPokemonByID(id)
	// if err != nil {
	// return nil, fmt.Errorf("unable to get Pokemon by ID: %w", err)
	// }

	pokemon := &model.Pokemon{
		ID:   id,
		Name: fmt.Sprintf("%d-mon", id),
		Type: []model.Type{
			model.Dark,
			model.Dragon,
		},
		Base: model.Base{
			HP:        id,
			Attack:    id,
			Defense:   id,
			SPAttack:  id,
			SPDefense: id,
			Speed:     id,
		},
	}

	return pokemon, nil
}
