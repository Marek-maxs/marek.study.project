package inflxudb

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb/v2"
	"github.com/influxdata/influxdb/v2/cmd/inspect"
	"github.com/influxdata/influxdb/v2/cmd/launcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	version = "dev"
	commit  = "none"
	date    = ""
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

	// upgrade binds options to env variables, so it must be added after rootCmd is isitialized
	upgradeCmd, err := upgrade.NewCommand(ctx, v)
	if err != nil {
		handleErr(err.Error())
	}
	rootCmd.AddCommand(upgradeCmd)
	inspectCmd, err := inspect.NewCommand(v)
	if err != nil {
		handleErr(err.Error())
	}

	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(recovery.NewCommand())
	downgradeCmd, err := downgrade.NewCommand(ctx, v)
}

func handleErr(err string) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func versionCmd() *cobra.Command {
	return &cobra.Command{}
}
