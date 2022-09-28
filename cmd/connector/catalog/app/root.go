package app

import (
	"sort"

	"gopkg.in/urfave/cli.v1"
)

func NewCliApp() *cli.App {
	app := cli.NewApp()

	app.Name = "mrs-service-cab-connector-catalog"
	app.Usage = "mrs-service-connector-catalog for Super Koch"

	opts := NewCatalogCmdOptions()
	opts.AddFlags(app)

	app.Action = func(c *cli.Context) error {

		proc := NewCatalogApp()
		return proc.Run(opts)
	}

	// sort flags by name
	sort.Sort(cli.FlagsByName(app.Flags))

	return app
}
