package keypaircommands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/fatih/structs"
	"github.com/jrperritt/rack/auth"
	"github.com/jrperritt/rack/output"
	"github.com/jrperritt/rack/util"
	osKeypairs "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/keypairs"
)

var create = cli.Command{
	Name:        "create",
	Usage:       fmt.Sprintf("%s %s create <keypairName> [flags]", util.Name, commandPrefix),
	Description: "Creates a keypair",
	Action:      commandCreate,
	Flags:       util.CommandFlags(flagsCreate, keysCreate),
	BashComplete: func(c *cli.Context) {
		util.CompleteFlags(util.CommandFlags(flagsCreate, keysCreate))
	},
}

func flagsCreate() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: "publicKey",
			Usage: strings.Join([]string{"[optional] The public ssh key to associate with the user's account.",
				"It may be the actual key or the file containing the key. If empty,",
				"the key will be created for you and returned in the output."}, "\n\t"),
		},
	}
}

var keysCreate = []string{"Name", "Fingerprint", "PublicKey", "PrivateKey"}

func commandCreate(c *cli.Context) {
	util.CheckArgNum(c, 1)
	keypairName := c.Args()[0]
	client := auth.NewClient("compute")
	opts := osKeypairs.CreateOpts{
		Name: keypairName,
	}

	if c.IsSet("publicKey") {
		s := c.String("publicKey")
		pk, err := ioutil.ReadFile(s)
		if err != nil {
			opts.PublicKey = string(pk)
		} else {
			opts.PublicKey = s
		}
	}

	o, err := keypairs.Create(client, opts).Extract()
	if err != nil {
		fmt.Printf("Error creating keypair [%s]: %s\n", keypairName, err)
		os.Exit(1)
	}
	f := func() interface{} {
		return structs.Map(o)
	}
	output.Print(c, &f, keysCreate)
}
