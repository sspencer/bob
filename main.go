package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
	"github.com/sspencer/bob/api"
)

const (
	usage    = "Start BOB"
	mysqlDSN = "BOB_MYSQL"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	portPtr := flag.Int("p", 8080, "port to run server on")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}

	flag.Parse()

	dsn := os.Getenv(mysqlDSN)
	if len(dsn) == 0 {
		return fmt.Errorf("%s must be set", mysqlDSN)
	}

	db, err := dbConnect(dsn, time.Minute)
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}

	defer db.Close()

	server := api.Server(db)

	addr := fmt.Sprintf(":%d", *portPtr)
	log.Printf("Starting server on %s\n", addr)
	log.Fatal(server.Run(addr))

	return nil
}

// dbConnect tries to connect to the DB under given DSN using a give driver
// in a loop until connection succeeds. timeout specifies the timeout for the
// loop.
func dbConnect(dsn string, timeout time.Duration) (*sql.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := sql.Open("mysql", dsn)
			if err != nil {
				return nil, errors.Wrap(err, "Could not open database")
			}

			// Open doesn't open a connection. Validate DSN data:
			err = db.Ping()
			if err != nil {
				return nil, errors.Wrap(err, "Could not connection to database")
			}

			return db, nil
		}
	}
}
