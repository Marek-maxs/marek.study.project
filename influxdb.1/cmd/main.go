package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb/v2"
	"github.com/influxdata/influxdb/v2/cmd/inspect"
	"os"
	"time"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit = "none"
	date = ""
)

func main() {
	if len(date) == 0 {
		date = time.Now().UTC().Format(time.RFC3339)
	}

	influxdb.SetBuildInfo(version, commit, date)

	ctx := context.Background()
	v := viper.New()

	rootCmd, err := launcher.NewInfluxdCommand(ctx, v)
	if err != nil {
		handleErr(err.Error())
	}

	inspectCmd, err := inspect.NewCommand(v)
	if err != nil {
		handleErr(err.Error())
	}
	rootCmd.
}

func handleErr(err string) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}