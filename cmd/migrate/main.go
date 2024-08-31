package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dialect      = "pgx"
	dbConnString = "host=localhost port=5432 user=user password=passw0rd dbname=books sslmode=disable"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with migration files")
)

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
	
Example:
	migrate status
`
	usageCommands = `
Commands:
up                   Migrate the DB to the most recent version available
up-by-one            Migrate the DB up by 1
up-to VERSION        Migrate the DB to a specific VERSION
down                 Roll back the version by 1
down-to VERSION      Roll back to a specific VERSION
redo                 Re-run the latest migration
reset                Roll back all migrations
status               Dump the migration status for the current DB
version              Print the current version of the database
create NAME [sql|go] Creates new migration file with the current timestamp
fix                  Apply sequential ordering to migrations`
)

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	cmd := args[0]

	db, err := goose.OpenDBWithDriver(dialect, dbConnString)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed to close db: %v", err)
		}
	}()

	if err := goose.RunContext(context.Background(), cmd, db, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate failed: %v %v", cmd, err)
	}
}
