package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/CineQuery/config"
	"github.com/saleh-ghazimoradi/CineQuery/migrations"
	"github.com/saleh-ghazimoradi/CineQuery/utils"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// migrateUpCmd represents the migrateUp command
var migrateUpCmd = &cobra.Command{
	Use:   "migrateUp",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateUp called")

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

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

		migrator, err := migrations.NewMigrator(db, postgresql.Name)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		if err := migrator.Up(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		defer func() {
			if err := migrator.Close(); err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}()
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
}
