package config

import "goAdmin/models"

var GlobalTableList = map[string]models.GlobalTable{
	"user": models.GetUserTable(),
}
