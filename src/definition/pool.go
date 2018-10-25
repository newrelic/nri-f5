package definition

// LtmPool is an unmarshalling struct
type LtmPool struct {
	Kind  string        `json:"kind"`
	Items []LtmPoolItem `json:"items"`
}

// LtmPoolItem is an unmarshalling struct
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

// LtmPoolItemMember is an unmarshalling struct
type LtmPoolItemMember struct {
	Kind string `json:"kind"`
}

// LtmPoolItemMembers is an unmarshalling struct
type LtmPoolItemMembers struct {
}

// =============

// LtmPoolStats is an unmarshalling struct
type LtmPoolStats struct {
	Kind    string                            `json:"kind"`
	Entries map[string]LtmPoolStatsEntryValue `json:"entries"`
}

// LtmPoolStatsEntryValue is an unmarshalling struct
type LtmPoolStatsEntryValue struct {
	NestedStats LtmPoolStatsEntryValueNestedStats `json:"nestedStats"`
}

// LtmPoolStatsEntryValueNestedStats is an unmarshalling struct
type LtmPoolStatsEntryValueNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		FullPath struct {
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
			ProcessedValue *int `metric_name:"pool.inDataInBytesPerSecond" source_type:"rate"`
			Value          int
		} `json:"serverside.bitsIn"`
		DataOut struct {
			ProcessedValue *int `metric_name:"pool.outDataInBytesPerSecond" source_type:"rate"`
			Value          int
		} `json:"serverside.bitsOut"`
		EnabledState struct {
			ProcessedDescription *int `metric_name:"pool.enabled" source_type:"gauge"`
			Description          string
		} `json:"status.enabledState"`
		PacketsIn struct {
			Value int `metric_name:"pool.packetsReceivedPerSecond" source_type:"rate"`
		} `json:"serverside.pktsIn"`
		PacketsOut struct {
			Value int `metric_name:"pool.packetsSentPerSecond" source_type:"rate"`
		} `json:"serverside.pktsOut"`
		Requests struct {
			Value int `metric_name:"pool.requestsPerSecond" source_type:"rate"`
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
		ConnqAgeEdm struct {
			Value int `metric_name:"pool.connqAgeEdm" source_type:"gauge"`
		} `json:"connq.ageEdm"`
		ConnqAgeEma struct {
			Value int `metric_name:"pool.connqAgeEma" source_type:"gauge"`
		} `json:"connq.ageEma"`
		ConnqAgeHead struct {
			Value int `metric_name:"pool.connqAgeHead" source_type:"gauge"`
		} `json:"connq.ageHead"`
		ConnqAgeMax struct {
			Value int `metric_name:"pool.connqAgeMax" source_type:"gauge"`
		} `json:"connq.ageMax"`
		ConnqDepth struct {
			Value int `metric_name:"pool.connqDepth" source_type:"gauge"`
		} `json:"connq.depth"`
		ConnqServiced struct {
			Value int `metric_name:"pool.connqServiced" source_type:"gauge"`
		} `json:"connq.serviced"`
		ConnqAllAgeEdm struct {
			Value int `metric_name:"pool.connqAllAgeEdm" source_type:"gauge"`
		} `json:"connqAll.ageEdm"`
		ConnqAllAgeEma struct {
			Value int `metric_name:"pool.connqAllAgeEma" source_type:"gauge"`
		} `json:"connqAll.ageEma"`
		ConnqAllAgeHead struct {
			Value int `metric_name:"pool.connqAllAgeHead" source_type:"gauge"`
		} `json:"connqAll.ageHead"`
		ConnqAllAgeMax struct {
			Value int `metric_name:"pool.connqAllAgeMax" source_type:"gauge"`
		} `json:"connqAll.ageMax"`
		ConnqAllDepth struct {
			Value int `metric_name:"pool.connqAllDepth" source_type:"gauge"`
		} `json:"connqAll.depth"`
		ConnqAllServiced struct {
			Value int `metric_name:"pool.connqAllServiced" source_type:"gauge"`
		} `json:"connqAll.serviced"`
		CurSessions struct {
			Value int `metric_name:"pool.currentConnections" source_type:"gauge"`
		} `json:"curSessions"`
		MinActiveMembers struct {
			Value int `metric_name:"pool.minActiveMembers" source_type:"gauge"`
		} `json:"minActiveMembers"`
	}
}
