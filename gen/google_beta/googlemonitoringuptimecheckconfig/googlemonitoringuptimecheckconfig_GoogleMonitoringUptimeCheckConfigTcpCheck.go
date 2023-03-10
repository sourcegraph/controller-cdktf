package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigTcpCheck struct {
	// The port to the page to run the check against.
	//
	// Will be combined with host (specified within the MonitoredResource) to construct the full URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_uptime_check_config#port GoogleMonitoringUptimeCheckConfig#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

