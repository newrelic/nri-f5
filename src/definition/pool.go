package definition

type LtmPool struct {
	Kind  string        `json:"kind"`
	Items []LtmPoolItem `json:"items"`
}

type LtmPoolItem struct {
	Name              string `json:"name"`
	Partition         string `json:"partition"`
	FullPath          string `json:"fullPath"`
	Kind              string `json:"kind"`
	LoadBalancingMode string `json:"loadBalancingMode"`
	Description       string `json:"description"`
}

// =============

type LtmPoolStats struct {
	Kind    string `json:"kind"`
	Entries map[string]LtmPoolStatsEntryValue
}

type LtmPoolStatsEntryValue struct {
	NestedStats LtmNodeStatsEntryValueNestedStats `json:"nestedStats"`
}

type LtmPoolStatsEntryValueNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		ActiveMemberCount struct {
			Value int
		} `json:"activeMemberCnt"`
		AvailabilityState struct {
			Description string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
			Value int
		} `json:"serverside.curConns"`
		DataIn struct {
			Value int
		} `json:"serverside.bitsIn"`
		DataOut struct {
			Value int
		} `json:"serverside.bitsOut"`
		EnabledState struct {
			Description string
		} `json:"status.enabledState"`
		PacketsIn struct {
			Value int
		} `json:"serverside.pktsIn"`
		PacketsOut struct {
			Value int
		} `json:"serverside.pktsOut"`
		Requests struct {
			Value int
		} `json:"totRequests"`
		StatusReason struct {
			Description string
		} `json:"status.statusReason"`
		MonitorRule struct {
			Description string
		} `json:"monitorRule"`
		MaxConnections struct {
			Value int
		} `json:"serverside.maxConns"`
	}
}
