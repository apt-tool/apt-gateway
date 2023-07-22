package main

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/cmd"
	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/storage/sql"

	"github.com/spf13/cobra"
)

func main() {
	// load configs
	cfg := config.Load("config.yml")

	// database connection
	db, err := sql.NewConnection(cfg.MySQL)
	if err != nil {
		log.Fatal(fmt.Errorf("[main] failed in connecting to mysql server error=%w", err))
	}

	// perform migrations if needed
	if cfg.MySQL.Migrate {
		migrateInstance := cmd.Migrate{
			Cfg: cfg.Migrate,
			Db:  db,
		}

		migrateInstance.Do()
	}

	// create root command
	root := cobra.Command{}

	// add sub commands to root
	root.AddCommand(
		cmd.API{
			Cfg: cfg,
			Db:  db,
		}.Command(),
		cmd.Core{
			Cfg: cfg,
			Db:  db,
		}.Command(),
	)

	// execute root command
	if er := root.Execute(); er != nil {
		log.Fatal(fmt.Errorf("[main] failed to execute command error=%w", er))
	}
}
