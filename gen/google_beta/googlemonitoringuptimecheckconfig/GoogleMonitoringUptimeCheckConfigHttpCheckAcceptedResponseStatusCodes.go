package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodes struct {
	// A class of status codes to accept. Possible values: ["STATUS_CLASS_1XX", "STATUS_CLASS_2XX", "STATUS_CLASS_3XX", "STATUS_CLASS_4XX", "STATUS_CLASS_5XX", "STATUS_CLASS_ANY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#status_class GoogleMonitoringUptimeCheckConfig#status_class}
	StatusClass *string `field:"optional" json:"statusClass" yaml:"statusClass"`
	// A status code to accept.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#status_value GoogleMonitoringUptimeCheckConfig#status_value}
	StatusValue *float64 `field:"optional" json:"statusValue" yaml:"statusValue"`
}

