package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Log *log.Logger

func main() {
	app := cli.NewApp()
	app.Name = "Intuition"
	app.Version = Version
	app.Usage = "Intuition API client"

	app.Flags = []cli.Flag{
		// TODO Impact log setup
		cli.BoolFlag{Name: "verbose", Usage: "Extends log output to debug level"},
		cli.StringFlag{Name: "addr", Value: "http://localhost:5000", Usage: "Telepathy HTTP API address"},
	}

	app.Action = func(c *cli.Context) {
		Log = log.New(os.Stdout, "[intuition] ", log.Ldate|log.Ltime|log.Lshortfile)

		api := NewIntuitionAPI(c.String("addr"))

		health, err := api.Health()
		Log.Printf("Error: %v\n", err)
		Log.Printf("State: %v\n", health.State)
		Log.Printf("Versions: %v\n", health.Version)
	}

	app.RunAndExitOnError()
}
