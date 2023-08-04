package main

import (
	"log"
	"xsis/movie/common/app"
	cf "xsis/movie/common/config"
)

func main() {
	cfg, err := cf.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
