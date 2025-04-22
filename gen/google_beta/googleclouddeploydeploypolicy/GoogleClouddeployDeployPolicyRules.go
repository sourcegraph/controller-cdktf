package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicyRules struct {
	// rollout_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#rollout_restriction GoogleClouddeployDeployPolicy#rollout_restriction}
	RolloutRestriction *GoogleClouddeployDeployPolicyRulesRolloutRestriction `field:"optional" json:"rolloutRestriction" yaml:"rolloutRestriction"`
}

