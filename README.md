# sample-clean

Sample clean architecture approach for a basic Go application, in the context of what a Pokemon information retrieval project might look like.

In this sample, we imagine we are retrieving information from a PostgreSQL DB (not yet implemented), but gives you an idea on how to organize your code in distinct layers, which responsibilities are well defined (sort of).

## Getting started

Clone this repo in your preferred location

``` bash
git clone https://github.com/RolloCasanova/sample-clean.git
```

### Prerequisites

- `Go` installed in your machine.

- Rename `config.yml.dist` to `config.yml` and replace `zero values` with the required values.

## Running the project

- Get/update dependencies by running:

``` bash
go mod tidy
```

- Run the project:

``` bash
go run main.go
```

## Basic endpoints

`hostname` and `port` variables are the ones inside `config.yml` file, under `server` field.

- Get all pokemons in DB:

``` bash
GET <hostname>:<port>/api/v1/pokemons
```

- Get a single Pokemon by it's ID:

``` bash
GET <hostname>:<port>/api/v1/pokemons/{id}
```

## TODOs

Add unit tests samples in the whole project.

## Contact

Roboam Casanova - r_casanova_c@hotmail.com
