package googlemonitoringnotificationchannel


type GoogleMonitoringNotificationChannelSensitiveLabels struct {
	// An authorization token for a notification channel. Channel types that support this field include: slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_notification_channel#auth_token GoogleMonitoringNotificationChannel#auth_token}
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// An password for a notification channel. Channel types that support this field include: webhook_basicauth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_notification_channel#password GoogleMonitoringNotificationChannel#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// An servicekey token for a notification channel. Channel types that support this field include: pagerduty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_notification_channel#service_key GoogleMonitoringNotificationChannel#service_key}
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
}

