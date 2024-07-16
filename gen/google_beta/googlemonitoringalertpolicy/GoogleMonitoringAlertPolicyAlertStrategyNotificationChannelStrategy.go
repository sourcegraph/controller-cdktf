package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyAlertStrategyNotificationChannelStrategy struct {
	// The notification channels that these settings apply to.
	//
	// Each of these
	// correspond to the name field in one of the NotificationChannel objects
	// referenced in the notification_channels field of this AlertPolicy. The format is
	// 'projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#notification_channel_names GoogleMonitoringAlertPolicy#notification_channel_names}
	NotificationChannelNames *[]*string `field:"optional" json:"notificationChannelNames" yaml:"notificationChannelNames"`
	// The frequency at which to send reminder notifications for open incidents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#renotify_interval GoogleMonitoringAlertPolicy#renotify_interval}
	RenotifyInterval *string `field:"optional" json:"renotifyInterval" yaml:"renotifyInterval"`
}

