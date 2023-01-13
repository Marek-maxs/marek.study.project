package toml

import "time"

// Duration is a TOML wrapper type for time.Duration.
type Duration time.Duration


// Size represents a TOML parsable file size.
// Users can specify size using "k" or "K" for kibibytes, "m" or "M" for mebibytes,
// and "g" or "G" for gibibytes. If a size suffix isn't specified then bytes are assumed.
type Size uint64