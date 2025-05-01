package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicyRulesRolloutRestriction struct {
	// Required.
	//
	// ID of the rule. This id must be unique in the 'DeployPolicy' resource to which this rule belongs. The format is 'a-z{0,62}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#id GoogleClouddeployDeployPolicy#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional.
	//
	// Rollout actions to be restricted as part of the policy. If left empty, all actions will be restricted. Possible values: ["ADVANCE", "APPROVE", "CANCEL", "CREATE", "IGNORE_JOB", "RETRY_JOB", "ROLLBACK", "TERMINATE_JOBRUN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#actions GoogleClouddeployDeployPolicy#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Optional. What invoked the action. If left empty, all invoker types will be restricted. Possible values: ["USER", "DEPLOY_AUTOMATION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#invokers GoogleClouddeployDeployPolicy#invokers}
	Invokers *[]*string `field:"optional" json:"invokers" yaml:"invokers"`
	// time_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#time_windows GoogleClouddeployDeployPolicy#time_windows}
	TimeWindows *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows `field:"optional" json:"timeWindows" yaml:"timeWindows"`
}

