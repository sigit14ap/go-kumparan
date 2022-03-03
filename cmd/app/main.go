package main

import (
	"flag"
	"github.com/sigit14ap/go-kumparan/internal/app"
	"github.com/sigit14ap/go-kumparan/internal/domain"
)

func main() {

	var seeds bool

	// flags declaration using flag package
	flag.BoolVar(&seeds, "seeds", false, "Running seeders")

	flag.Parse() // after declaring flags we need to call it

	command := domain.Command{
		Seeds: seeds,
	}

	app.Run("config/config.yml", command)
}
