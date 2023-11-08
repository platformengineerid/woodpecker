package user

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"go.woodpecker-ci.org/woodpecker/cli/common"
	"go.woodpecker-ci.org/woodpecker/cli/internal"
	"go.woodpecker-ci.org/woodpecker/woodpecker-go/woodpecker"
)

var userAddCmd = &cli.Command{
	Name:      "add",
	Usage:     "adds a user",
	ArgsUsage: "<username>",
	Action:    userAdd,
	Flags:     common.GlobalFlags,
}

func userAdd(c *cli.Context) error {
	login := c.Args().First()

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	user, err := client.UserPost(&woodpecker.User{Login: login})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added user %s\n", user.Login)
	return nil
}
