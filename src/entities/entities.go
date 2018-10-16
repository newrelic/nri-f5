package entities

import (
	"github.com/newrelic/infra-integrations-sdk/log"
)

func logOnError(itemName, entityName string, err error) {
	if err != nil {
		log.Error(`Failed to set inventory item %s for %s: %s`, itemName, entityName, err.Error())
	}
}
