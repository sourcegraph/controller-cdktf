package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule struct {
	// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#matching_type GoogleDataLossPreventionInspectTemplate#matching_type}
	MatchingType *string `field:"required" json:"matchingType" yaml:"matchingType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#dictionary GoogleDataLossPreventionInspectTemplate#dictionary}
	Dictionary *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// exclude_info_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#exclude_info_types GoogleDataLossPreventionInspectTemplate#exclude_info_types}
	ExcludeInfoTypes *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#regex GoogleDataLossPreventionInspectTemplate#regex}
	Regex *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex `field:"optional" json:"regex" yaml:"regex"`
}

