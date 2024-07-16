package googleclouddeployautomation


type GoogleClouddeployAutomationRules struct {
	// advance_rollout_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_automation#advance_rollout_rule GoogleClouddeployAutomation#advance_rollout_rule}
	AdvanceRolloutRule *GoogleClouddeployAutomationRulesAdvanceRolloutRule `field:"optional" json:"advanceRolloutRule" yaml:"advanceRolloutRule"`
	// promote_release_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_automation#promote_release_rule GoogleClouddeployAutomation#promote_release_rule}
	PromoteReleaseRule *GoogleClouddeployAutomationRulesPromoteReleaseRule `field:"optional" json:"promoteReleaseRule" yaml:"promoteReleaseRule"`
}

