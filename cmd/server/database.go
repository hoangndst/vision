package server

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/hoangndst/vision/models"
	"github.com/hoangndst/vision/server/util/credentials"
	"gorm.io/gorm"

	"github.com/hoangndst/vision/server"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var _ Options = &DatabaseOptions{}

// DatabaseOptions is a Database options struct
type DatabaseOptions struct {
	DatabaseAccessOptions `json:",inline" yaml:",inline"`
	// AutoMigrate will attempt to automatically migrate all tables
	AutoMigrate bool   `json:"autoMigrate,omitempty" yaml:"autoMigrate,omitempty"`
	MigrateFile string `json:"migrateFile,omitempty" yaml:"migrateFile,omitempty"`
}

// NewDatabaseOptions returns a DatabaseOptions instance with the default values
func NewDatabaseOptions() *DatabaseOptions {
	return &DatabaseOptions{
		DatabaseAccessOptions: DatabaseAccessOptions{},
		AutoMigrate:           false,
	}
}

// Validate checks DatabaseOptions and return a slice of found error(s)
func (o *DatabaseOptions) Validate() error {
	if o == nil {
		return errors.Errorf("options is nil")
	}

	// if o.AutoMigrate && len(o.MigrateFile) == 0 {
	// 	return errors.Errorf("when --auto-migrate is true, --migrate-file must be specified")
	// }

	return o.DatabaseAccessOptions.Validate()
}

// ApplyTo apply database options to the server config
func (o *DatabaseOptions) ApplyTo(config *server.Config) {
	if err := o.DatabaseAccessOptions.ApplyTo(&config.DB); err != nil {
		logrus.Fatalf("Failed to apply database options to server.Config as: %+v", err)
	}

	config.AutoMigrate = o.AutoMigrate

	// AutoMigrate will attempt to automatically migrate all tables
	if o.AutoMigrate && len(o.MigrateFile) > 0 {
		logrus.Debugf("AutoMigrate will attempt to automatically migrate all tables from [%s]", o.MigrateFile)
		// Read all content by migrate file
		migrateSQL, err := os.ReadFile(o.MigrateFile)
		if err != nil {
			logrus.Fatalf("Failed to read migrate file: %+v", err)
		}

		// Split multiple SQL statements into individual statements
		stmts := strings.Split(string(migrateSQL), ";")

		// Iterate over all statements and execute them
		for _, stmt := range stmts {
			// Ignore empty statements
			if len(strings.TrimSpace(stmt)) == 0 {
				continue
			}

			// Use gorm.Exec() function to execute SQL statement
			if err = config.DB.Exec(stmt).Error; err != nil {
				logrus.Warnf("Failed to exec migrate sql: %+v", err)
			}
		}
	}

	if config.DB != nil && config.AutoMigrate {
		err := models.AutoMigrate(config.DB)
		if err != nil {
			logrus.Fatalf("Failed to auto migrate: %+v", err)
		}
	}

	if err := InitUserAdminIfNotExist(config.DB); err != nil {
		logrus.Fatalf("Failed to init user admin: %+v", err)
	}
}

// AddFlags adds flags for a specific Option to the specified FlagSet
func (o *DatabaseOptions) AddFlags(fs *pflag.FlagSet) {
	o.DatabaseAccessOptions.AddFlags(fs)

	autoMigrate, err := strconv.ParseBool(AutoMigrateEnv)
	if err != nil {
		autoMigrate = false
	}
	fs.BoolVar(&o.AutoMigrate, "auto-migrate", autoMigrate, "Whether to enable automatic migration")
	fs.StringVar(&o.MigrateFile, "migrate-file", MigrateFileEnv, "The migrate sql file")
}

// MarshalJSON is custom marshalling function for masking sensitive field values
func (o DatabaseOptions) MarshalJSON() ([]byte, error) {
	type tempOptions DatabaseOptions
	o2 := tempOptions(o)
	o2.DBPassword = MaskString
	return json.Marshal(&o2)
}

func InitUserAdminIfNotExist(db *gorm.DB) error {
	var adminCount int64
	err := db.Model(&models.UserModel{}).Count(&adminCount).Error
	// adminCount == 0 or user table is empty or error that user table not exist
	if adminCount == 0 || errors.Is(err, gorm.ErrRecordNotFound) {
		if len(AdminPasswordEnv) == 0 {
			return errors.New("ADMIN_PASSWORD is not found")
		}
		argon2iHash := credentials.NewDefaultArgon2idHash()
		password, err := argon2iHash.HashPassword(AdminPasswordEnv, nil)
		if err != nil {
			return errors.Wrap(err, "failed to hash admin password")
		}
		username := AdminUsernameEnv
		if len(username) == 0 {
			username = DefaultAdminUsername
		}
		user := &models.UserModel{
			Name:     username,
			Username: username,
			Email:    "admin@vision.com",
			Password: password,
		}
		if err = db.Create(user).Error; err != nil {
			return errors.Wrap(err, "failed to create default admin user")
		}
	}
	return nil
}
