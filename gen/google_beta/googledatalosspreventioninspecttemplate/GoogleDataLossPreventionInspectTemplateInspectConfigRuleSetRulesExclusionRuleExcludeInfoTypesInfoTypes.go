package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes struct {
	// Name of the information type.
	//
	// Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#name GoogleDataLossPreventionInspectTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#sensitivity_score GoogleDataLossPreventionInspectTemplate#sensitivity_score}
	SensitivityScore *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#version GoogleDataLossPreventionInspectTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

