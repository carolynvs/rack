package blockstoragecommands

import (
	"github.com/codegangsta/cli"
	"github.com/jrperritt/rackcli/blockstoragecommands/snapshotcommands"
	"github.com/jrperritt/rackcli/blockstoragecommands/volumecommands"
)

// Get returns all the commands allowed for a `compute` request.
func Get() []cli.Command {
	return []cli.Command{
		{
			Name:        "snapshots",
			Usage:       "Used for BlockStorage Snapshot operations",
			Subcommands: snapshotcommands.Get(),
		},
		{
			Name:        "volumes",
			Usage:       "Used for BlockStorage Volume operations",
			Subcommands: volumecommands.Get(),
		},
	}
}
