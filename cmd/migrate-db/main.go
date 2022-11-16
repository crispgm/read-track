package main

import (
	"os"

	"github.com/crispgm/read-track/internal/app"
)

func main() {
	var appl app.Application
	var err error
	path := "./"
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}
	err = appl.Init(path)
	if err != nil {
		panic(err)
	}
	err = appl.MigrateDB()
	if err != nil {
		panic(err)
	}
}
