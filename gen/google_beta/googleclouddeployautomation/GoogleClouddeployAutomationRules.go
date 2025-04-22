package googleclouddeployautomation


type GoogleClouddeployAutomationRules struct {
	// advance_rollout_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#advance_rollout_rule GoogleClouddeployAutomation#advance_rollout_rule}
	AdvanceRolloutRule *GoogleClouddeployAutomationRulesAdvanceRolloutRule `field:"optional" json:"advanceRolloutRule" yaml:"advanceRolloutRule"`
	// promote_release_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#promote_release_rule GoogleClouddeployAutomation#promote_release_rule}
	PromoteReleaseRule *GoogleClouddeployAutomationRulesPromoteReleaseRule `field:"optional" json:"promoteReleaseRule" yaml:"promoteReleaseRule"`
	// repair_rollout_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#repair_rollout_rule GoogleClouddeployAutomation#repair_rollout_rule}
	RepairRolloutRule *GoogleClouddeployAutomationRulesRepairRolloutRule `field:"optional" json:"repairRolloutRule" yaml:"repairRolloutRule"`
	// timed_promote_release_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#timed_promote_release_rule GoogleClouddeployAutomation#timed_promote_release_rule}
	TimedPromoteReleaseRule *GoogleClouddeployAutomationRulesTimedPromoteReleaseRule `field:"optional" json:"timedPromoteReleaseRule" yaml:"timedPromoteReleaseRule"`
}

