package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#hotword_regex GoogleDataLossPreventionJobTrigger#hotword_regex}
	HotwordRegex *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex `field:"optional" json:"hotwordRegex" yaml:"hotwordRegex"`
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#likelihood_adjustment GoogleDataLossPreventionJobTrigger#likelihood_adjustment}
	LikelihoodAdjustment *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `field:"optional" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#proximity GoogleDataLossPreventionJobTrigger#proximity}
	Proximity *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity `field:"optional" json:"proximity" yaml:"proximity"`
}

