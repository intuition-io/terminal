/*
 * Intuition API client - package main wraps high level abstraction
 */
package main

import (
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitchellh/panicwrap"
)

var Log = NewUi()

func terminal() error {
	app := cli.NewApp()
	app.Name = "Intuition"
	app.Version = Version
	app.Usage = "Intuition API client"

	// Flags available for every subcommands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "http://localhost:5000",
			Usage: "Telepathy HTTP API address",
		},
	}

	app.Commands = setupCommands()
	//app.RunAndExitOnError()
	return app.Run(os.Args)
}

func appWrapper() int {
	// We also always send logs to a temporary file that we use in case
	// there is a panic. Otherwise, we delete it.
	logTempFile, err := ioutil.TempFile("", "packer-log")
	if err != nil {
		Log.Error("Couldn't setup logging tempfile: %s", err)
		return 1
	}
	defer os.Remove(logTempFile.Name())
	defer logTempFile.Close()

	// Create the configuration for panicwrap and wrap our executable
	wrapConfig := &panicwrap.WrapConfig{
		Handler: panicHandler(logTempFile),
		Writer:  logTempFile,
	}

	exitStatus, err := panicwrap.Wrap(wrapConfig)
	if err != nil {
		Log.Error("Couldn't start Terminal: %s", err)
		return 1
	}
	if exitStatus >= 0 {
		return 1
	}

	// We're the child, so just close the tempfile we made in order to
	// save file handles since the tempfile is only used by the parent.
	logTempFile.Close()

	if err := terminal(); err != nil {
		Log.Error("%s", err)
		return 1
	}

	return 0
}

func main() {
	// Call realMain instead of doing the work here so we can use
	// `defer` statements within the function and have them work properly.
	// (defers aren't called with os.Exit)
	os.Exit(appWrapper())
}
