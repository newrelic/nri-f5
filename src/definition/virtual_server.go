package definition

// LtmVirtual is an unmarshalling struct
type LtmVirtual struct {
	Kind  string           `json:"kind"`
	Items []LtmVirtualItem `json:"items"`
}

// LtmVirtualItem is an unmarshalling struct
type LtmVirtualItem struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	FullPath       string `json:"fullPath"`
	Destination    string `json:"destination"`
	MaxConnections int    `json:"connectionLimit"`
	Pool           string `json:"pool"`
	AppService     string `json:"appService"`
}

// =================

// LtmVirtualStats is an unmarshalling struct
type LtmVirtualStats struct {
	Kind    string `json:"kind"`
	Entries map[string]LtmVirtualStatsEntryValue
}

// LtmVirtualStatsEntryValue is an unmarshalling struct
type LtmVirtualStatsEntryValue struct {
	NestedStats LtmVirtualStatsNestedStats
}

// LtmVirtualStatsNestedStats is an unmarshalling struct
type LtmVirtualStatsNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		AvailabilityState struct {
			ProcessedDescription *int `metric_name:"virtualserver.availabilityState" source_type:"gauge"`
			Description          string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
			Value int `metric_name:"virtualserver.connections" source_type:"gauge"`
		} `json:"clientside.curConns"`
		DataIn struct {
			ProcessedValue *int `metric_name:"virtualserver.inDataInBytesPerSecond" source_type:"rate"`
			Value          int
		} `json:"clientside.bitsIn"`
		DataOut struct {
			ProcessedValue *int `metric_name:"virtualserver.outDataInBytesPerSecond" source_type:"rate"`
			Value          int
		} `json:"clientside.bitsOut"`
		EnabledState struct {
			ProcessedDescription *int `metric_name:"virtualserver.enabled" source_type:"gauge"`
			Description          string
		} `json:"status.enabledState"`
		PacketsIn struct {
			Value int `metric_name:"virtualserver.packetsReceivedPerSecond" source_type:"rate"`
		} `json:"clientside.pktsIn"`
		PacketsOut struct {
			Value int `metric_name:"virtualserver.packetsSentPerSecond" source_type:"rate"`
		} `json:"clientside.pktsOut"`
		Requests struct {
			Value int `metric_name:"virtualserver.requestsPerSecond" source_type:"rate"`
		} `json:"totRequests"`
		StatusReason struct {
			Description string `metric_name:"virtualserver.statusReason" source_type:"attribute"`
		} `json:"status.statusReason"`
		TmName struct {
			Description string
		} `json:"tmName"`
	}
}
