package definition

type LtmNode struct {
	Kind  string        `json:"kind"`
	Items []LtmNodeItem `json:"items"`
}

type LtmNodeItem struct {
	Name           string          `json:"name"`
	Partition      string          `json:"partition"`
	FullPath       string          `json:"fullPath"`
	MonitorRule    string          `json:"monitor"`
	Session        string          `json:"user-enabled"`
	State          string          `json:"state"`
	Kind           string          `json:"kind"`
	Address        string          `json:"address"`
	MaxConnections int             `json:"connectionLimit"`
	FQDN           LtmNodeItemFQDN `json:"fqdn"`
}

type LtmNodeItemFQDN struct {
	TMName string `json:"tmName"`
}

// ===============

type LtmNodeStats struct {
	Kind    string                            `json:"kind"`
	Entries map[string]LtmNodeStatsEntryValue `json:"entries"`
}

type LtmNodeStatsEntryValue struct {
	NestedStats LtmNodeStatsEntryValueNestedStats `json:"nestedStats"`
}

// TODO add metric names and types when those are determined
type LtmNodeStatsEntryValueNestedStats struct {
	Kind    string `json:"kind"`
	Entries struct {
		AvailabilityState struct {
      ProcessedDescription *int `metric_name:"node.availabilityState" source_type:"gauge"`
      Description string
		} `json:"status.availabilityState"`
		CurrentConnections struct {
      Value int `metric_name:"node.connections" source_type:"gauge"`
		} `json:"serverside.curConns"`
		CurrentSessions struct {
      Value int `metric_name:"node.sessions" source_type:"gauge"`
		} `json:"curSessions"`
		DataIn struct {
      ProcessedValue int `metric_name:"node.inDataInBytes" source_type:"rate"`
      Value int
		} `json:"serverside.bitsIn"`
		DataOut struct {
      ProcessedValue int `metric_name:"node.outDataInBytes" source_type:"rate"`
      Value int
		} `json:"serverside.bitsOut"`
		EnabledState struct {
      ProcessedDescription *int `metric_name:"node.enabled" source_type:"gauge"`
      Description string
		} `json:"status.enabledState"`
		MonitorStatus struct {
      ProcessedDescription *int `metric_name:"node.monitorStatus" source_type:"gauge"`
      Description string
		} `json:"monitorStatus"`
		PacketsIn struct {
      Value int `metric_name:"node.packetsReceived" source_type:"rate"`
		} `json:"serverside.pktsIn"`
		PacketsOut struct {
      Value int `metric_name:"node.packetsSent" source_type:"rate"`
		} `json:"serverside.pktsOut"`
		Requests struct {
      Value int `metric_name:"node.requests" source_type:"rate"`
		} `json:"totRequests"`
		SessionStatus struct {
      ProcessedDescription *int `metric_name:"node.sessionStatus" source_type:"gauge"`
		} `json:"sessionStatus"`
		StatusReason struct {
      Description string `metric_name:"node.statusReason" source_type:"attribute"`
		} `json:"status.statusReason"`
	}
}
