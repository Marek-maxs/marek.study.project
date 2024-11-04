package export_lp

import (
	"github.com/influxdata/influxdb/v2/kit/cli"
	"github.com/influxdata/influxdb/v2/kit/platform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)
// exportFlags contains CLI-compatible forms of export options.
type exportFlags struct {
	enginePath string
	bucketID platform.ID
	measurements []string
	startTime string
	endTime string

	outputPath string
	compress bool

	logLevel zapcore.Level
}

// NewExportLineProtocolCommand builds and registers the `export` subcommand of `influxdb inspect.`
func NewExportLineProtocolCommand(v *viper.Viper) (*cobra.Command, error) {
	flags := newFlags()

	cmd := &cobra.Command{
		Use:                        `export-lp`,
		Short:                      "Export TSM data as line protocol",
		Long:                       `
This command will export all TSM data stored in a bucket
to line protocol for inspection and re-ingestion.`,
		Args:                       cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}

	opts := []cli.Opt{
		{
			DestP:    &flags.enginePath,
			Flag:     "engine-path",
			Desc:     "path to persistent engine files",
			Required: true,
		},
		{
			DestP:    &flags.bucketID,
			Flag:     "bucket-id",
			Desc:     "ID of bucket containing data to export",
			Required: true,
		},
		{
			DestP: &flags.measurements,
			Flag:  "measurement",
			Desc:  "optional: name(s) of specific measurement to export",
		},
		{
			DestP: &flags.startTime,
			Flag:  "start",
			Desc:  "optional: the start time to export (RFC3339 format)",
		},
		{
			DestP: &flags.endTime,
			Flag:  "end",
			Desc:  "optional: the end time to export (RFC3339 format)",
		},
		{
			DestP:    &flags.outputPath,
			Flag:     "output-path",
			Desc:     "path where exported line-protocol should be written. Use '-' to write to standard out",
			Required: true,
		},
		{
			DestP: &flags.compress,
			Flag:  "compress",
			Desc:  "if true, compress output with GZIP",
		},
		{
			DestP:   &flags.logLevel,
			Flag:    "log-level",
			Default: flags.logLevel,
		},
	}

	if err := cli.BindOptions(v, cmd, opts); err != nil {
		return nil, err
	}
	return cmd, nil
}

func newFlags() *exportFlags {
	return &exportFlags{
		compress:     false,
		logLevel:     zapcore.InfoLevel,
	}
}