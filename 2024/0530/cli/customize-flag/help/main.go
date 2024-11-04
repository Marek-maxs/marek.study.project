package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	cli.HelpFlag = &cli.BoolFlag{
		Name:        "Haaalp",
		Usage:       "HALP",
		EnvVar:      "",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Destination: nil,
	}

	(&cli.App{}).Run(os.Args)
}
