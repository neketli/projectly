package main

import (
	"errors"
	"flag"
	"projectly-server/config"
	"projectly-server/migrations"
)

var (
	up       = flag.Bool("up", false, "Runs up migrations")
	down     = flag.Bool("down", false, "Runs down migrations")
	helpFlag = flag.Bool("help", false, "View default help message")
)

func main() {
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
	}

	if *up && *down {
		panic(errors.New("up and down flags cannot be used together"))
	}

	cfg := config.New()

	if *up {
		migrations.Up(cfg.PG.DSN)
	} else if *down {
		migrations.Down(cfg.PG.DSN)
	} else {
		flag.PrintDefaults()
		panic(errors.New("not a single flag has been applied"))
	}
}
