package models

type ServerInfo struct {
	ServerVersion    string `json:"version"`
	MasterHash       string `json:"master_hash"`
	AssertHash       string `json:"assert_hash"`
	MaintenanceState int    `json:"maintenance_state"`
	ResponseTmp
}
