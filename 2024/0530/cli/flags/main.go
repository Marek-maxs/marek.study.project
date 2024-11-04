package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:    "Lang",
		Usage:   "language for the greeting",
		Version: "english",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Usage: "language for the greeting",
				Value: "english",
			},
		},
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			if c.String("lang") == "english" {
				fmt.Println("helle", name)
			} else {
				fmt.Println("你好", name)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("app run failed")
	}
}
