package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#hotword_regex GoogleDataLossPreventionJobTrigger#hotword_regex}
	HotwordRegex *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordHotwordRegex `field:"optional" json:"hotwordRegex" yaml:"hotwordRegex"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#proximity GoogleDataLossPreventionJobTrigger#proximity}
	Proximity *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordProximity `field:"optional" json:"proximity" yaml:"proximity"`
}

