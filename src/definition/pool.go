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
	MembersReference  struct {
		Items []LtmPoolItemMember `json:"items"`
	} `json:"membersReference"`
}

type LtmPoolItemMember struct {
	Kind string `json:"kind"`
}

type LtmPoolItemMembers struct {
}

// =============

type LtmPoolStats struct {
	Kind    string                            `json:"kind"`
	Entries map[string]LtmPoolStatsEntryValue `json:"entries"`
}

type LtmPoolStatsEntryValue struct {
	NestedStats LtmPoolStatsEntryValueNestedStats `json:"nestedStats"`
}

type LtmPoolStatsEntryValueNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		Name struct {
			Description string
		} `json:"tmName"`
		ActiveMemberCount struct {
			Value int `metric_name:"pool.activeMembers" source_type:"gauge"`
		} `json:"activeMemberCnt"`
		AvailabilityState struct {
			ProcessedDescription *int `metric_name:"pool.availabilityState" source_type:"gauge"`
			Description          string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
			Value int `metric_name:"pool.connections" source_type:"gauge"`
		} `json:"serverside.curConns"`
		DataIn struct {
			ProcessedValue *int `metric_name:"pool.inDataInBytes" source_type:"rate"`
			Value          int
		} `json:"serverside.bitsIn"`
		DataOut struct {
			ProcessedValue *int `metric_name:"pool.outDataInBytes" source_type:"rate"`
			Value          int
		} `json:"serverside.bitsOut"`
		EnabledState struct {
			ProcessedDescription *int `metric_name:"pool.enabled" source_type:"rate"`
			Description          string
		} `json:"status.enabledState"`
		PacketsIn struct {
			Value int `metric_name:"pool.packetsReceived" source_type:"rate"`
		} `json:"serverside.pktsIn"`
		PacketsOut struct {
			Value int `metric_name:"pool.packetsSent" source_type:"rate"`
		} `json:"serverside.pktsOut"`
		Requests struct {
			Value int `metric_name:"pool.requests" source_type:"rate"`
		} `json:"totRequests"`
		StatusReason struct {
			Description string `metric_name:"pool.statusReason" source_type:"attribute"`
		} `json:"status.statusReason"`
		MonitorRule struct {
			Description string
		} `json:"monitorRule"`
		MaxConnections struct {
			Value int
		} `json:"serverside.maxConns"`
	}
}
