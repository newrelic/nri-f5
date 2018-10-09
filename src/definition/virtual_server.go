package definition

type LtmVirtual struct {
  Kind string `json:"kind"`
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
}

// =================

type LtmVirtualStats struct {
  Kind string `json:"kind"`
  Entries map[string]LtmVirtualStatsEntryValue
}

type LtmVirtualStatsEntryValue struct {
  NestedStats LtmVirtualStatsNestedStats
}

type LtmVirtualStatsNestedStats struct {
  Kind string `json:"kind"`
  Entries struct {
    AvailabilityState struct {
      Description string
    }, `json:"status.availabilityState"`
    CurrentConnections struct {
      Value int
    }, `json:"clientside.curConns"`
    DataIn struct {
      Value int
    }, `json:"clientside.bitsIn"`
    DataOut struct {
      Value int
    }, `json:"clientside.bitsOut"`
    EnabledState struct {
      Description string
    }, `json:"status.enabledState"`
    PacketsIn struct {
      Value int
    }, `json:"clientside.pktsIn"`
    PacketsOut struct {
      Value int
    }, `json:"clientside.pktsOut"`
    Requests struct {
      Value int
    }, `json:"totRequests"`
    StatusReason struct {
      Description string
    }, `json:"status.statusReason"`
  }  
}
