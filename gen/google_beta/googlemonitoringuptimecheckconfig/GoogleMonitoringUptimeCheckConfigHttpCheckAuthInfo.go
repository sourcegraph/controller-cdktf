package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo struct {
	// The username to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_uptime_check_config#username GoogleMonitoringUptimeCheckConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_uptime_check_config#password GoogleMonitoringUptimeCheckConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_uptime_check_config#password_wo GoogleMonitoringUptimeCheckConfig#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// The password write-only version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_uptime_check_config#password_wo_version GoogleMonitoringUptimeCheckConfig#password_wo_version}
	PasswordWoVersion *string `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
}

