package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyAlertStrategyNotificationRateLimit struct {
	// Not more than one notification per period.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example "60.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#period GoogleMonitoringAlertPolicy#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
}

