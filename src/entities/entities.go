package entities

func convertMonitorStatus(rule string) *int {
	var i int
	switch rule {
	case "down":
		i = 0
	case "unchecked":
		i = 1
	default:
		i = 2
	}
	return &i
}

func convertEnabledState(rule string) *int {
	var i int
	switch rule {
	case "disabled":
		i = 0
	case "enabled":
		i = 1
	default:
		i = 2
	}
	return &i
}

func convertAvailabilityState(rule string) *int {
	var i int
	switch rule {
	case "offline":
		i = 0
	case "unknown":
		i = 1
	default:
		i = 2
	}
	return &i
}

func convertSessionStatus(rule string) *int {
	var i int
	switch rule {
	case "disabled":
		i = 0
	case "enabled":
		i = 1
	default:
		i = 2
	}
	return &i
}
