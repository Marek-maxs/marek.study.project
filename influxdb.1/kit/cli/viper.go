package cli

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type Opt struct {
	DestP interface{} // pointer to the desctnation

	EnvVar string
	Flag string
	Hidden bool
	Persistent bool
	Required bool
	Short rune // using rune b/c it guarantees correctness. a short must always be a string of length 1

	Default interface{}
	Desc string
}

// BindOptions adds opts to the specified command and automatically
// registers those options with viper.
func BindOptions(v *viper.Viper, cmd *cobra.Command, opts []Opt) error {
	for _, o := range opts {
		flagset := cmd.Flags()
		if o.Persistent {
			flagset = cmd.PersistentFlags()
		}
		envVal := lookupEnv(v, &o)
		hasShort := o.Short != 0

		switch destP := o.DestP.(type) {
		case *string:
			var d string
			if o.Default != nil {
				d = o.Default.(string)
			}
			if hasShort {
				flagset.StringVarP(destP, o.Flag, string(o.Short), d, o.Desc)
			} else {
				flagset.StringVar(destP, o.Flag, d, o.Desc)
			}
			if err := v.BindPFlag(o.Flag, flagset.Lookup(o.Flag)); err != nil {
				return fmt.Errorf("failed to bind flag %q: %w", o.Flag, err)
			}
			if envVal != nil {
				if s, err := cast.ToStringE(envVal); err == nil {
					*destP = s
				}
			}
		}
	}
}

// lookupEnv returns the value for a CLI option found in the environment, if any
func lookupEnv(v *viper.Viper, o *Opt) interface{} {
	envVar := o.Flag
	if o.EnvVar != "" {
		envVar = o.EnvVar
	}
	return v.Get(envVar)
}