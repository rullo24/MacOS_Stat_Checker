package collector

type NetworkStats struct {
	TotalUploadSpeedKbitsSec   float64
	TotalDownloadSpeedKbitsSec float64
	PublicIp                   string
	LocalIp                    string
	Interface                  string
	MACAddress                 string
}
