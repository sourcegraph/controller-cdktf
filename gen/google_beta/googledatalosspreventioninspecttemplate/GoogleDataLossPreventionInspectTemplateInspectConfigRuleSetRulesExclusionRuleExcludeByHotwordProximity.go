package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordProximity struct {
	// Number of characters after the finding to consider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#window_after GoogleDataLossPreventionInspectTemplate#window_after}
	WindowAfter *float64 `field:"optional" json:"windowAfter" yaml:"windowAfter"`
	// Number of characters before the finding to consider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#window_before GoogleDataLossPreventionInspectTemplate#window_before}
	WindowBefore *float64 `field:"optional" json:"windowBefore" yaml:"windowBefore"`
}

