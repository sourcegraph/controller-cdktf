package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo struct {
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_uptime_check_config#password GoogleMonitoringUptimeCheckConfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username to authenticate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_uptime_check_config#username GoogleMonitoringUptimeCheckConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

