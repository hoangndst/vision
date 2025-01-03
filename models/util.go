package models

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/hoangndst/vision/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MultiString []string

// Scan implements the Scanner interface for the MultiString type.
func (s *MultiString) Scan(src any) error {
	switch src := src.(type) {
	case []byte:
		*s = strings.Split(string(src), ",")
	case string:
		*s = strings.Split(src, ",")
	case nil:
		*s = nil
	default:
		return fmt.Errorf("unsupported type %T", src)
	}
	return nil
}

// Value implements the Valuer interface for the MultiString type.
func (s MultiString) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return strings.Join(s, ","), nil
}

// GormDataType gorm common data type
func (s MultiString) GormDataType() string {
	return "text"
}

// GormDBDataType gorm db data type
func (s MultiString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "postgres", "mysql", "sqlite":
		return "text"
	}
	return ""
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserModel{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&OrganizationModel{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&ProjectModel{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&BlogModel{}); err != nil {
		return err
	}
	return nil
}

func GetProjectQuery(filter *entity.ProjectFilter) (string, []interface{}) {
	pattern := make([]string, 0)
	args := make([]interface{}, 0)
	if filter.OrgID.String() != "" {
		pattern = append(pattern, "organization_id = ?")
		args = append(args, fmt.Sprint(filter.OrgID))
	}
	if filter.Name != "" {
		pattern = append(pattern, "name = ?")
		args = append(args, filter.Name)
	}
	return CombineQueryParts(pattern), args
}

func CombineQueryParts(queryParts []string) string {
	queryString := ""
	if len(queryParts) > 0 {
		queryString = queryParts[0]
		for _, part := range queryParts[1:] {
			queryString += fmt.Sprintf(" AND %s", part)
		}
	}
	return queryString
}

var LoadModels = []interface{}{
	&UserModel{},
	&OrganizationModel{},
	&ProjectModel{},
	&BlogModel{},
}
