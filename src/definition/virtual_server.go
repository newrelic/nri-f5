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
		EvictedConns struct {
			Value int `metric_name:"virtualserver.evictedConnsPerSecond" source_type:"rate"`
		} `clientside.evictedConns`
		SlowKilled struct {
			Value int `metric_name:"virtualserver.slowKilledPerSecond" source_type:"rate"`
		} `clientside.slowKilled`
		TotalConns struct {
			Value int `metric_name:"virtualserver.clientsideConnectionsPerSecond" source_type:"rate"`
		} `json:"clientside.totConns"`
		CmpEnableMode struct {
			Description string `metric_name:"virtualserver.cmpEnableMode" source_type:"attribute"`
		} `json:"cmpEnableMode"`
		CmpEnabled struct {
			Description string `metric_name:"virtualserver.cmpEnabled" source_type:"attribute"`
		} `json:"cmpEnabled"`
		CsMaxConnDur struct {
			Value int `metric_name:"virtualserver.csMaxConnDur" source_type:"gauge"`
		} `json:"csMaxConnDur"`
		CsMinConnDur struct {
			Value int `metric_name:"virtualserver.csMinConnDur" source_type:"gauge"`
		} `json:"csMinConnDur"`
		CsMeanConnDur struct {
			Value int `metric_name:"virtualserver.csMeanConnDur" source_type:"gauge"`
		} `json:"csMeanConnDur"`
		EphemeralBytesIn struct {
			ProcessedValue *int `metric_name:"virtualserver.ephemeralBytesInPerSecond" source_type:"rate"`
			Value          int
		} `json:"ephemeral.bitsIn"`
		EphemeralBytesOut struct {
			ProcessedValue *int `metric_name:"virtualserver.ephemeralBytesOutPerSecond" source_type:"rate"`
			Value          int
		} `json:"ephemeral.bitsOut"`
		EphemeralCurrentConnections struct {
			Value int `metric_name:"virtualserver.ephemeralCurrentConnections" source_type:"gauge"`
		} `json:"ephemeral.curConns"`
		EphemeralEvictedConnections struct {
			Value int `metric_name:"virtualserver.ephemeralEvictedConnectionsPerSecond" source_type:"rate"`
		} `json:"ephemeral.evictedConns"`
		EphemeralMaxConnections struct {
			Value int `metric_name:"virtualserver.ephemeralMaxConnections" source_type:"gauge"`
		} `json:"ephemeral.maxConns"`
		EphemeralPacketsIn struct {
			Value int `metric_name:"virtualserver.ephemeralPacketsReceivedPerSecond" source_type:"rate"`
		} `json:"ephemeral.pktsIn"`
		EphemeralPacketsOut struct {
			Value int `metric_name:"virtualserver.ephemeralPacketsSentPerSecond" source_type:"rate"`
		} `json:"ephemeral.pktsOut"`
		EphemeralSlowKilled struct {
			Value int `metric_name:"virtualserver.ephemeralSlowKilledPerSecond" source_type:"rate"`
		} `json:"ephemeral.slowKilled"`
		EphemeralTotalConnections struct {
			Value int `metric_name:"virtualserver.ephemeralConnectionsPerSecond" source_type:"rate"`
		} `json:"ephemeral.totConns"`
		UsageRatio struct {
			Value int `metric_name:"virtualserver.usageRatio" source_type:"rate"`
		} `json:"oneMinAvgUsageRatio"`
		SyncookieAccepts struct {
			Value int `metric_name:"virtualserver.syncookieAcceptsPerSecond" source_type:"rate"`
		} `json:"syncookie.accepts"`
		SyncookieHwAccepts struct {
			Value int `metric_name:"virtualserver.syncookieHwAcceptsPerSecond" source_type:"rate"`
		} `json:"syncookie.hwAccepts"`
		SyncookieHwInstance struct {
			Value int `metric_name:"virtualserver.hwSyncookieInstance" source_type:"gauge"`
		} `json:"syncookie.hwsyncookieInstance"`
		SyncookieRejects struct {
			Value int `metric_name:"virtualserver.syncookieRejectsPerSecond" source_type:"rate"`
		} `json:"syncookie.rejects"`
		SyncookieSwInstance struct {
			Value int `metric_name:"virtualserver.swSyncookieInstance" source_type:"gauge"`
		} `json:"syncookie.swsyncookieInstance"`
		SyncookieSyncacheCurr struct {
			Value int `metric_name:"virtualserver.syncookieSyncacheCurr" source_type:"gauge"`
		} `json:"syncookie.syncacheCurr"`
		Syncookies struct {
			Value int `metric_name:"virtualserver.syncookies" source_type:"gauge"`
		} `json:"syncookie.syncookies"`
	}
}
