package googleclouddeployautomation


type GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhases struct {
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#retry GoogleClouddeployAutomation#retry}
	Retry *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry `field:"optional" json:"retry" yaml:"retry"`
	// rollback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#rollback GoogleClouddeployAutomation#rollback}
	Rollback *GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback `field:"optional" json:"rollback" yaml:"rollback"`
}

