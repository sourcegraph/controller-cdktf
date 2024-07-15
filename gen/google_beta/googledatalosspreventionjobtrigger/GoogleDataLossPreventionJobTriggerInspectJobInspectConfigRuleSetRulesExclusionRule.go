package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule struct {
	// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#matching_type GoogleDataLossPreventionJobTrigger#matching_type}
	MatchingType *string `field:"required" json:"matchingType" yaml:"matchingType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#dictionary GoogleDataLossPreventionJobTrigger#dictionary}
	Dictionary *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// exclude_by_hotword block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#exclude_by_hotword GoogleDataLossPreventionJobTrigger#exclude_by_hotword}
	ExcludeByHotword *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword `field:"optional" json:"excludeByHotword" yaml:"excludeByHotword"`
	// exclude_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#exclude_info_types GoogleDataLossPreventionJobTrigger#exclude_info_types}
	ExcludeInfoTypes *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#regex GoogleDataLossPreventionJobTrigger#regex}
	Regex *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex `field:"optional" json:"regex" yaml:"regex"`
}

