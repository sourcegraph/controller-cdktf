package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRules struct {
	// exclusion_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#exclusion_rule GoogleDataLossPreventionJobTrigger#exclusion_rule}
	ExclusionRule *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule `field:"optional" json:"exclusionRule" yaml:"exclusionRule"`
	// hotword_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#hotword_rule GoogleDataLossPreventionJobTrigger#hotword_rule}
	HotwordRule *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule `field:"optional" json:"hotwordRule" yaml:"hotwordRule"`
}

