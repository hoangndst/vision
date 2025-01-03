package server

import (
	"os"

	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Port        int
	Database    DatabaseOptions
	LogFilePath string
	GithubToken string
}

type Options interface {
	// Validate checks Options and return a slice of found error(s)
	Validate() error
	// AddFlags adds flags for a specific Option to the specified FlagSet
	AddFlags(fs *pflag.FlagSet)
}

const (
	MaskString           = "******"
	DefaultAdminUsername = "admin"
	DefaultDBPort        = 5432
	DefaultPort          = 3000
)

var (
	DBHostEnv        = os.Getenv("DB_HOST")
	DBPortEnv        = os.Getenv("DB_PORT")
	DBUserEnv        = os.Getenv("DB_USER")
	DBPassEnv        = os.Getenv("DB_PASS")
	DBNameEnv        = os.Getenv("DB_NAME")
	PortEnv          = os.Getenv("PORT")
	LogFilePathEnv   = os.Getenv("LOG_FILE_PATH")
	AutoMigrateEnv   = os.Getenv("AUTO_MIGRATE")
	MigrateFileEnv   = os.Getenv("MIGRATE_FILE")
	AdminUsernameEnv = os.Getenv("ADMIN_USERNAME")
	AdminPasswordEnv = os.Getenv("ADMIN_PASSWORD")
	GithubTokenEnv   = os.Getenv("GITHUB_TOKEN")
)
