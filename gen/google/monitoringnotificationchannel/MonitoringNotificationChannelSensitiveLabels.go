package monitoringnotificationchannel


type MonitoringNotificationChannelSensitiveLabels struct {
	// An authorization token for a notification channel. Channel types that support this field include: slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#auth_token MonitoringNotificationChannel#auth_token}
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// An authorization token for a notification channel. Channel types that support this field include: slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#auth_token_wo MonitoringNotificationChannel#auth_token_wo}
	AuthTokenWo *string `field:"optional" json:"authTokenWo" yaml:"authTokenWo"`
	// Triggers update of 'auth_token_wo' write-only.
	//
	// Increment this value when an update to 'auth_token_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#auth_token_wo_version MonitoringNotificationChannel#auth_token_wo_version}
	AuthTokenWoVersion *string `field:"optional" json:"authTokenWoVersion" yaml:"authTokenWoVersion"`
	// An password for a notification channel. Channel types that support this field include: webhook_basicauth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#password MonitoringNotificationChannel#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// An password for a notification channel. Channel types that support this field include: webhook_basicauth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#password_wo MonitoringNotificationChannel#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Triggers update of 'password_wo' write-only.
	//
	// Increment this value when an update to 'password_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#password_wo_version MonitoringNotificationChannel#password_wo_version}
	PasswordWoVersion *string `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// An servicekey token for a notification channel. Channel types that support this field include: pagerduty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#service_key MonitoringNotificationChannel#service_key}
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
	// An servicekey token for a notification channel. Channel types that support this field include: pagerduty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#service_key_wo MonitoringNotificationChannel#service_key_wo}
	ServiceKeyWo *string `field:"optional" json:"serviceKeyWo" yaml:"serviceKeyWo"`
	// Triggers update of 'service_key_wo' write-only.
	//
	// Increment this value when an update to 'service_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/monitoring_notification_channel#service_key_wo_version MonitoringNotificationChannel#service_key_wo_version}
	ServiceKeyWoVersion *string `field:"optional" json:"serviceKeyWoVersion" yaml:"serviceKeyWoVersion"`
}

