package entities

import (
	"github.com/newrelic/infra-integrations-sdk/log"
)

func logOnError(itemName, entityName string, err error) {
	if err != nil {
		log.Error(`Failed to set inventory item %s for %s: %s`, itemName, entityName, err.Error())
	}
}

func convertMonitorStatus(rule string) *int {
	var i int
	switch rule {
	case "down":
		i = 0
		return &i
	case "unchecked":
		i = 1
		return &i
	default:
		i = 2
		return &i
	}
}

func convertEnabledState(rule string) *int {
	var i int
	switch rule {
	case "disabled":
		i = 0
		return &i
	case "enabled":
		i = 1
		return &i
	default:
		i = 2
		return &i
	}
}

func convertAvailabilityState(rule string) *int {
	var i int
	switch rule {
	case "offline":
		i = 0
		return &i
	case "unknown":
		i = 1
		return &i
	default:
		i = 2
		return &i
	}
}

func convertSessionStatus(rule string) *int {
	var i int
	switch rule {
	case "disabled":
		i = 0
		return &i
	case "enabled":
		i = 1
		return &i
	default:
		i = 2
		return &i
	}
}
