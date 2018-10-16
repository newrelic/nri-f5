package definition

/*
 *type LtmPoolMember struct {
 *  Kind  string              `json:"kind"`
 *  Items []LtmPoolMemberItem `json:"items"`
 *}
 *
 *type LtmPoolMemberItem struct {
 *  Name string `json:"name"`
 *  Kind string `json:"kind"`
 *}
 *
 */
// ===============

// LtmPoolMemberStats is an unmarshalling struct
type LtmPoolMemberStats struct {
	Kind    string                                  `json:"kind"`
	Entries map[string]LtmPoolMemberStatsEntryValue `json:"entries"`
}

// LtmPoolMemberStatsEntryValue is an unmarshalling struct
type LtmPoolMemberStatsEntryValue struct {
	NestedStats LtmPoolMemberStatsEntryValueNestedStats `json:"nestedStats"`
}

// LtmPoolMemberStatsEntryValueNestedStats is an unmarshalling struct
type LtmPoolMemberStatsEntryValueNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		AvailabilityState struct {
			ProcessedDescription *int `metric_name:"member.availabilityState" source_type:"gauge"`
			Description          string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
			Value int `metric_name:"member.connections" source_type:"gauge"`
		} `json:"serverside.curConns"`
		CurrentSessions struct {
			Value int `metric_name:"member.sessions" source_type:"gauge"`
		} `json:"curSessions"`
		DataIn struct {
			ProcessedValue *int `metric_name:"member.inDataInBytes" source_type:"rate"`
			Value          int
		} `json:"serverside.bitsIn"`
		DataOut struct {
			ProcessedValue *int `metric_name:"member.outDataInBytes" source_type:"rate"`
		} `json:"serverside.bitsOut"`
		EnabledState struct {
			ProcessedDescription *int `metric_name:"member.enabled" source_type:"gauge"`
			Description          string
		} `json:"status.enabledState"`
		MonitorStatus struct {
			ProcessedDescription *string `metric_name:"member.monitorStatus" source_type:"gauge"`
			Description          string
		} `json:"monitorStatus"`
		PacketsIn struct {
			Value int `metric_name:"member.packetsReceived" source_type:"rate"`
		} `json:"serverside.pktsIn"`
		PacketsOut struct {
			Value int `metric_name:"member.packetsSent" source_type:"rate"`
		} `json:"serverside.pktsOut"`
		Requests struct {
			Value int `metric_name:"member.requests" source_type:"rate"`
		} `json:"totRequests"`
		SessionStatus struct {
			ProcessedDescription *int `metric_name:"member.sessionStatus" source_type:"gauge"`
			Description          string
		} `json:"sessionStatus"`
		StatusReason struct {
			Description string `metric_name:"member.statusReason" source_type:"attribute"`
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
