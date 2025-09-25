package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/CineQuery/config"
	"github.com/saleh-ghazimoradi/CineQuery/internal/gateway/handlers"
	"github.com/saleh-ghazimoradi/CineQuery/internal/gateway/routes"
	"github.com/saleh-ghazimoradi/CineQuery/internal/helper"
	"github.com/saleh-ghazimoradi/CineQuery/internal/middleware"
	"github.com/saleh-ghazimoradi/CineQuery/internal/repository"
	"github.com/saleh-ghazimoradi/CineQuery/internal/server"
	"github.com/saleh-ghazimoradi/CineQuery/internal/service"
	"github.com/saleh-ghazimoradi/CineQuery/internal/validator"
	"github.com/saleh-ghazimoradi/CineQuery/utils"
	"log"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		customErr := helper.NewCustomErr()
		middlewares := middleware.NewMiddleware()
		validate := validator.NewValidator()

		cfg, err := config.NewConfig()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		postgresql := utils.NewPostgres(
			utils.WithHost(cfg.Postgres.Host),
			utils.WithPort(cfg.Postgres.Port),
			utils.WithUser(cfg.Postgres.User),
			utils.WithPassword(cfg.Postgres.Password),
			utils.WithName(cfg.Postgres.Name),
			utils.WithMaxOpenConn(cfg.Postgres.MaxOpenConn),
			utils.WithMaxIdleConn(cfg.Postgres.MaxIdleConn),
			utils.WithMaxIdleTime(cfg.Postgres.MaxIdleTime),
			utils.WithSSLMode(cfg.Postgres.SSLMode),
		)

		db, err := postgresql.Connect()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		defer func() {
			if err := db.Close(); err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}()

		healthHandler := handlers.NewHealthHandler(logger, cfg, customErr)
		healthRoutes := routes.NewHealthRoutes(healthHandler)

		movieRepository := repository.NewMovieRepository(db, db)
		movieService := service.NewMovieService(movieRepository)
		movieHandler := handlers.NewMovieHandler(customErr, movieService, validate)
		movieRoutes := routes.NewMovieRoutes(movieHandler)

		router := routes.NewRegisterRoutes(
			routes.WithCustomErr(customErr),
			routes.WithMiddleWares(middlewares),
			routes.WithHealthRoutes(healthRoutes),
			routes.WithMovieRoutes(movieRoutes),
		)

		s := server.NewServer(
			server.WithHost(cfg.Server.Host),
			server.WithPort(cfg.Server.Port),
			server.WithHandler(router.Register()),
			server.WithIdleTimeout(cfg.Server.IdleTimeout),
			server.WithReadTimeout(cfg.Server.ReadTimeout),
			server.WithWriteTimeout(cfg.Server.WriteTimeout),
			server.WithErrorLog(slog.NewLogLogger(logger.Handler(), slog.LevelError)),
		)

		logger.Info("starting server", "addr", cfg.Server.Host+":"+cfg.Server.Port, "env", cfg.Application.Environment)

		if err := s.Connect(); err != nil {
			log.Fatalf("error connecting to server: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
