package command

import (
	"fmt"
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

// CmdBootstrap pulls and builds the services in the docker-compose file
func CmdBootstrap(c *cli.Context) error {
	CmdTeardown(c)

	if !c.GlobalBool("offline") {
		if err := helpers.RunCommand(helpers.DockerComposeCommand("build", "--pull")); err != nil {
			return err
		}
		pullCmd := helpers.DockerComposeCommand("pull", "--ignore-pull-failures")
		benignError := regexp.MustCompile(fmt.Sprintf(`Error: image library/%s(:latest|) not found`, c.String("docker-image-name")))
		helpers.RunCommandWithFilter(pullCmd, benignError.MatchString)
	}

	return CmdInstallDependencies(c)
}
