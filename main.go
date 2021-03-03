package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RolloCasanova/sample-clean/configuration"
	"github.com/RolloCasanova/sample-clean/controller"
	"github.com/RolloCasanova/sample-clean/router"
	pokemonSQLService "github.com/RolloCasanova/sample-clean/service/sql"
	"github.com/RolloCasanova/sample-clean/usecase"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
)

func main() {
	// Set logger
	l := logrus.New()

	// enable following line if wanted to set logger to debug level
	// l.SetLevel(logrus.DebugLevel)

	// Configure project's configuration file
	viper.SetConfigFile("./config.yml")
	viper.AutomaticEnv()

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		l.Fatalf("Error reading config file: %v", err)

		return
	}

	// Unmarshal configuration into cfg variable
	var cfg configuration.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		l.Fatalf("Unable to decode into struct: %v", err)

		return
	}

	// To avoid having long dependencies (cfg.PostgreSQL.Name, cfg.PostgreSQL.User, etc.)
	db := cfg.PostgreSQL

	pokemonService, _ := pokemonSQLService.NewPokemonService(db.Name, db.User, db.Password, db.Host, db.Port, l)
	pokemonUsecase, _ := usecase.NewPokemonUsecase(pokemonService, l)
	pokemonController, _ := controller.NewPokemonController(render.New(), pokemonUsecase, l)

	// Error is ignored as this function never returns an error (at this moment)
	// Additional comments on router.Setup
	httpRouter, _ := router.Setup(pokemonController)

	// Build HTTP server object
	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:           handlers.LoggingHandler(os.Stdout, httpRouter),
		WriteTimeout:      cfg.Server.WriteTimeout * time.Second,
		ReadTimeout:       cfg.Server.ReadTimeout * time.Second,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout * time.Second,
	}

	// Launch server in a goroutine
	go func() {
		l.Infof("starting server in address, %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			l.Fatalf("starting server: %v", err)
		}
	}()

	// Catch application's interrupt signals (Kill, Hang up and Interrupt)
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	<-stop

	l.Infoln("shutting down the server...")

	ctxTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)

	// Gracefully shutdown server by using Shutdown() method
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("shutting down server gracefully: %w", err)

		return
	}

	defer cancel()

	l.Infoln("server gracefully stopped")
}
