package instancecommands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jrperritt/rack/auth"
	"github.com/jrperritt/rack/output"
	"github.com/jrperritt/rack/util"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

var get = cli.Command{
	Name:        "get",
	Usage:       fmt.Sprintf("%s %s get %s [optional flags]", util.Name, commandPrefix, idOrNameUsage),
	Description: "Retrieves an existing server",
	Action:      commandGet,
	Flags:       util.CommandFlags(flagsGet, keysGet),
	BashComplete: func(c *cli.Context) {
		util.CompleteFlags(util.CommandFlags(flagsGet, keysGet))
	},
}

func flagsGet() []cli.Flag {
	return idAndNameFlags
}

var keysGet = []string{"ID", "Name", "Status", "Created", "Updated", "Image", "Flavor", "Public IPv4", "Public IPv6", "Private IPv4", "KeyName"}

func commandGet(c *cli.Context) {
	util.CheckArgNum(c, 0)
	client := auth.NewClient("compute")
	serverID := idOrName(c, client)
	o, err := servers.Get(client, serverID).Extract()
	if err != nil {
		fmt.Printf("Error retrieving server (%s): %s\n", serverID, err)
		os.Exit(1)
	}

	f := func() interface{} {
		return serverSingle(o)
	}
	output.Print(c, &f, keysGet)
}
