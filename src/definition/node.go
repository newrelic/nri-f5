package definition

type LtmNodeItem struct {
  Name           string `json:"name"`
  Partition      string `json:"partition"`
  FullPath       string `json:"fullPath"`
  MonitorRule    string `json:"monitor"`
  Session        string `json:"user-enabled"`
  State          string `json:"state"`
  Kind           string `json:"kind"`
  Address        string `json:"address"`
  MaxConnections int    `json:"connectionLimit"`
}

// ===============


type LtmNodeStatsEntryValue struct {
  NestedStats LtmNodeStatsEntryValueNestedStats `json:"nestedStats"`
}

// TODO add metric names and types when those are determined
type LtmNodeStatsEntryValueNestedStats struct {
  Kind string `json:"kind"`
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
    EnabledState struct {
      Description string
    } `json:"status.enabledState"`
    StatusReason struct {
      Description string
    } `json:"status.statusReason"`
  }
}

