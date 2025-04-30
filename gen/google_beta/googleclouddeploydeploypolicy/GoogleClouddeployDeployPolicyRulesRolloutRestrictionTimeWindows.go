package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows struct {
	// Required. The time zone in IANA format IANA Time Zone Database (e.g. America/New_York).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#time_zone GoogleClouddeployDeployPolicy#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// one_time_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#one_time_windows GoogleClouddeployDeployPolicy#one_time_windows}
	OneTimeWindows interface{} `field:"optional" json:"oneTimeWindows" yaml:"oneTimeWindows"`
	// weekly_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#weekly_windows GoogleClouddeployDeployPolicy#weekly_windows}
	WeeklyWindows interface{} `field:"optional" json:"weeklyWindows" yaml:"weeklyWindows"`
}

