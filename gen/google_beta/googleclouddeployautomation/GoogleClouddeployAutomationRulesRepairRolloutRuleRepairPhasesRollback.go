package googleclouddeployautomation


type GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback struct {
	// Optional. The starting phase ID for the Rollout. If unspecified, the Rollout will start in the stable phase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#destination_phase GoogleClouddeployAutomation#destination_phase}
	DestinationPhase *string `field:"optional" json:"destinationPhase" yaml:"destinationPhase"`
	// Optional. If pending rollout exists on the target, the rollback operation will be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_automation#disable_rollback_if_rollout_pending GoogleClouddeployAutomation#disable_rollback_if_rollout_pending}
	DisableRollbackIfRolloutPending interface{} `field:"optional" json:"disableRollbackIfRolloutPending" yaml:"disableRollbackIfRolloutPending"`
}

