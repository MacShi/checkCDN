package modle

type CdnResult struct {
	HaveCDN string `json:"have_cdn"`
	IpAddr []string `json:"ip_addr"`
	LossRate []string `json:"loss_rate"`
}
