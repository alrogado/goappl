package main

import (
	"log"

	"github.com/gobuffalo/buffalo/goappl/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
