package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule struct {
	// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#matching_type GoogleDataLossPreventionInspectTemplate#matching_type}
	MatchingType *string `field:"required" json:"matchingType" yaml:"matchingType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#dictionary GoogleDataLossPreventionInspectTemplate#dictionary}
	Dictionary *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// exclude_by_hotword block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#exclude_by_hotword GoogleDataLossPreventionInspectTemplate#exclude_by_hotword}
	ExcludeByHotword *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByHotword `field:"optional" json:"excludeByHotword" yaml:"excludeByHotword"`
	// exclude_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#exclude_info_types GoogleDataLossPreventionInspectTemplate#exclude_info_types}
	ExcludeInfoTypes *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#regex GoogleDataLossPreventionInspectTemplate#regex}
	Regex *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex `field:"optional" json:"regex" yaml:"regex"`
}

