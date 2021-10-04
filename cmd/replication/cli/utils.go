package cli

import (
	"io/ioutil"

	"token-strike/config"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var (
	homeDir = "$HOME/.rpl/"

	cfg = &config.Config{}
)

func preRunDecorator(f func(ctx *cli.Context) error) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		// read global flags
		if home := ctx.GlobalString("home"); home != "" {
			homeDir = home
		}

		// read config file
		dataConfig, err := ioutil.ReadFile(homeDir)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(dataConfig, &cfg); err != nil {
			return err
		}

		return f(ctx)
	}
}
