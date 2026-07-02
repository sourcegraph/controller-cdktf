package xraygroup

type XrayGroupInsightsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/xray_group#insights_enabled XrayGroup#insights_enabled}.
	InsightsEnabled any `field:"required" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/xray_group#notifications_enabled XrayGroup#notifications_enabled}.
	NotificationsEnabled any `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}
