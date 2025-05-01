package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindows struct {
	// Optional.
	//
	// Days of week. If left empty, all days of the week will be included. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#days_of_week GoogleClouddeployDeployPolicy#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#end_time GoogleClouddeployDeployPolicy#end_time}
	EndTime *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime `field:"optional" json:"endTime" yaml:"endTime"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#start_time GoogleClouddeployDeployPolicy#start_time}
	StartTime *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}

