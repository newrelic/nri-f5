package definition

type LtmPoolMember struct {
	Kind  string              `json:"kind"`
	Items []LtmPoolMemberItem `json:"items"`
}

type LtmPoolMemberItem struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// ===============

type LtmPoolMemberStats struct {
	Kind    string                                  `json:"kind"`
	Entries map[string]LtmPoolMemberStatsEntryValue `json:"entries"`
}

type LtmPoolMemberStatsEntryValue struct {
	NestedStats LtmPoolMemberStatsEntryValueNestedStats `json:"nestedStats"`
}

// TODO add metric names and types when those are determined
type LtmPoolMemberStatsEntryValueNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		AvailabilityState struct {
			Description string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
			Value int
		} `json:"serverside.curConns"`
		CurrentSessions struct {
			Value int
		} `json:"curSessions"`
		DataIn struct {
			Value int
		} `json:"serverside.bitsIn"`
		DataOut struct {
			Value int
		} `json:"serverside.bitsOut"`
		EnabledState struct {
			Description string
		} `json:"status.enabledState"`
		MonitorStatus struct {
			Description string
		} `json:"monitorStatus"`
		PacketsIn struct {
			Value int
		} `json:"serverside.pktsIn"`
		PacketsOut struct {
			Value int
		} `json:"serverside.pktsOut"`
		Requests struct {
			Value int
		} `json:"totRequests"`
		SessionStatus struct {
			Description string
		} `json:"sessionStatus"`
		StatusReason struct {
			Description string
		} `json:"status.statusReason"`
		// inventory
		MaximumConnections struct {
			Value int
		} `json:"serverside.maxConns"`
		MonitorRule struct {
			Description string
		} `json:"monitorRule"`
		NodeName struct {
			Description string
		} `json:"nodeName"`
		PoolName struct {
			Description string
		} `json:"poolName"`
		Port struct {
			Value int
		} `json:"port"`
	}
}
