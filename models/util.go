package models

import (
	"fmt"
	"github.com/hoangndst/vision/domain/entity"
	"gorm.io/gorm"
)

type MultiString []string

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
	if err := db.AutoMigrate(&UserOrganizationModel{}); err != nil {
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
	//&UserOrganizationModel{},
	&ProjectModel{},
}
