package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigTcpCheckPingConfig struct {
	// Number of ICMP pings. A maximum of 3 ICMP pings is currently supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#pings_count GoogleMonitoringUptimeCheckConfig#pings_count}
	PingsCount *float64 `field:"required" json:"pingsCount" yaml:"pingsCount"`
}

