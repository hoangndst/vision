package server

import (
	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Mode        string
	Port        int
	Database    DatabaseOptions
	LogFilePath string
}

type Options interface {
	// Validate checks Options and return a slice of found error(s)
	Validate() error
	// AddFlags adds flags for a specific Option to the specified FlagSet
	AddFlags(fs *pflag.FlagSet)
}

const (
	MaskString  = "******"
	DefaultPort = 3000
)
