package main

import (
	"github.com/codegangsta/cli"
)

// setupCommands wraps main functions with cli documentation
// and return a suitable data structure for cli package to register
func setupCommands() []cli.Command {
	return []cli.Command{

		{
			Name:  "health",
			Usage: "Check Intuition API status",
			Action: func(c *cli.Context) {
				api := NewIntuitionAPI(c.GlobalString("addr"))
				if health, err := api.Health(); err == nil {
					Log.Info("[API version %v] Health state:", health.Version)
					for i := 0; i < len(health.Components); i++ {
						Log.Info("%v | %v | %v",
							health.Components[i].Name,
							health.Components[i].Version,
							health.Components[i].State,
						)
					}
				} else {
					Log.Error("Unable to read API health: %v", err)
				}
			},
		},

		{
			Name:  "docs",
			Usage: "Inspect API documentation",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "about",
					Value: "all",
					Usage: "Specific documentation",
				},
			},
			Action: func(c *cli.Context) {
				api := NewIntuitionAPI(c.GlobalString("addr"))
				if docs, err := api.Doc(c.String("about")); err == nil {
					Log.Info("[API version %v] Docs:", docs.Version)
					for i := 0; i < len(docs.Resources); i++ {
						Log.Info("%v", docs.Resources[i])
					}
				} else {
					Log.Error("Unable to read API docs: %v", err)
				}
			},
		},
	}
}
