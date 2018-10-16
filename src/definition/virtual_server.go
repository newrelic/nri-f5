package definition

type LtmVirtual struct {
	Kind  string           `json:"kind"`
	Items []LtmVirtualItem `json:"items"`
}

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

type LtmVirtualStats struct {
	Kind    string `json:"kind"`
	Entries map[string]LtmVirtualStatsEntryValue
}

type LtmVirtualStatsEntryValue struct {
	NestedStats LtmVirtualStatsNestedStats
}

type LtmVirtualStatsNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		AvailabilityState struct {
			ParsedDescription *int `metric_name:"virtualserver.availabilityState" source_type:"gauge"`
			Description       string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
			Value int `metric_name:"virtualserver.connections" source_type:"gauge"`
		} `json:"clientside.curConns"`
		DataIn struct {
			ParsedValue *int `metric_name:"virtualserver.inDataInBytes" source_type:"rate"`
			Value       int
		} `json:"clientside.bitsIn"`
		DataOut struct {
			ParsedValue *int `metric_name:"virtualserver.outDataInBytes" source_type:"rate"`
			Value       int
		} `json:"clientside.bitsOut"`
		EnabledState struct {
			ParsedDescription *int `metric_name:"virtualserver.enabled" source_type:"gauge"`
		} `json:"status.enabledState"`
		PacketsIn struct {
			Value int `metric_name:"virtualserver.packetsReceived" source_type:"rate"`
		} `json:"clientside.pktsIn"`
		PacketsOut struct {
			Value int `metric_name:"virtualserver.packetsSent" source_type:"rate"`
		} `json:"clientside.pktsOut"`
		Requests struct {
			Value int `metric_name:"virtualserver.requests" source_type:"rate"`
		} `json:"totRequests"`
		StatusReason struct {
			Description string `metric_name:"virtualserver.statusReason" source_type:"attribute"`
		} `json:"status.statusReason"`
		TmName struct {
			Description string
		} `json:"tmName"`
	}
}
