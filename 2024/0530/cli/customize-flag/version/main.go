package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:        "print-version",
		Usage:       "print only the version",
		EnvVar:      "",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Destination: nil,
	}

	app := &cli.App{
		Name:        "Version",
		HelpName:    "",
		Usage:       "",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "v1.1.0",
		Description: "marek study cli program",
		Compiled:    time.Now(),
		Author:      "Marek",
		Email:       "364021318@qq.com",
	}

	app.Run(os.Args)
}
