package main

import (
	"github.com/codegangsta/cli"
	"github.com/jrperritt/rackcli/blockstoragecommands"
	"github.com/jrperritt/rackcli/computecommands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "rack"
	app.Usage = "An opinionated CLI for the Rackspace cloud"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:        "compute",
			Usage:       "Used for the Compute service",
			Subcommands: computecommands.Get(),
		},
		{
			Name:        "blockstorage",
			Usage:       "Used for the BlockStorage service",
			Subcommands: blockstoragecommands.Get(),
		},
	}
	app.Run(os.Args)
}
