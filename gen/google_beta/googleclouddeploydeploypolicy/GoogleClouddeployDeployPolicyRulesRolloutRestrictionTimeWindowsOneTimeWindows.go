package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindows struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#end_date GoogleClouddeployDeployPolicy#end_date}
	EndDate *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#end_time GoogleClouddeployDeployPolicy#end_time}
	EndTime *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime `field:"required" json:"endTime" yaml:"endTime"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#start_date GoogleClouddeployDeployPolicy#start_date}
	StartDate *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#start_time GoogleClouddeployDeployPolicy#start_time}
	StartTime *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

