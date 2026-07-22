package collector

type NetworkStats struct {
	TotalUploadSpeedKbitsSec   float64
	TotalDownloadSpeedKbitsSec float64
	PublicIp                   string
	LocalIp                    string
	Interface                  string
	MACAddress                 string
	TopProcesses               []ProcessNetworkStats
}

type ProcessNetworkStats struct {
	Name                  string
	UploadSpeedKBitsSec   float64
	DownloadSpeedKBitsSec float64
}
