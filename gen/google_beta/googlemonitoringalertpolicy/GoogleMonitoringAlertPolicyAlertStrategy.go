package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyAlertStrategy struct {
	// If an alert policy that was active has no data for this long, any open incidents will close.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_monitoring_alert_policy#auto_close GoogleMonitoringAlertPolicy#auto_close}
	AutoClose *string `field:"optional" json:"autoClose" yaml:"autoClose"`
	// notification_rate_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_monitoring_alert_policy#notification_rate_limit GoogleMonitoringAlertPolicy#notification_rate_limit}
	NotificationRateLimit *GoogleMonitoringAlertPolicyAlertStrategyNotificationRateLimit `field:"optional" json:"notificationRateLimit" yaml:"notificationRateLimit"`
}

