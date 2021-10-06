package cli

import (
	"os"

	"token-strike/config"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var initCommand = cli.Command{
	Name:      "init",
	Usage:     "initializing home directory (default \"" + homeDir + "\")",
	ArgsUsage: "",
	Flags:     []cli.Flag{},
	Action:    initialization,
}

func initialization(ctx *cli.Context) error {

	data, err := yaml.Marshal(config.DefaultConfig)
	if err != nil {
		return err
	}

	_ = os.Mkdir(os.ExpandEnv(homeDir), 0777)
	return os.WriteFile(
		os.ExpandEnv(homeDir)+"config.yaml",
		data,
		0777,
	)
}
