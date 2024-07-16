package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypes struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#info_type GoogleDataLossPreventionInspectTemplate#info_type}
	InfoType *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType `field:"required" json:"infoType" yaml:"infoType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#dictionary GoogleDataLossPreventionInspectTemplate#dictionary}
	Dictionary *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned.
	//
	// It still can be used for rules matching. Possible values: ["EXCLUSION_TYPE_EXCLUDE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#exclusion_type GoogleDataLossPreventionInspectTemplate#exclusion_type}
	ExclusionType *string `field:"optional" json:"exclusionType" yaml:"exclusionType"`
	// Likelihood to return for this CustomInfoType.
	//
	// This base value can be altered by a detection rule if the finding meets the criteria
	// specified by the rule. Default value: "VERY_LIKELY" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#likelihood GoogleDataLossPreventionInspectTemplate#likelihood}
	Likelihood *string `field:"optional" json:"likelihood" yaml:"likelihood"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#regex GoogleDataLossPreventionInspectTemplate#regex}
	Regex *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex `field:"optional" json:"regex" yaml:"regex"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#sensitivity_score GoogleDataLossPreventionInspectTemplate#sensitivity_score}
	SensitivityScore *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// stored_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#stored_type GoogleDataLossPreventionInspectTemplate#stored_type}
	StoredType *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType `field:"optional" json:"storedType" yaml:"storedType"`
	// surrogate_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#surrogate_type GoogleDataLossPreventionInspectTemplate#surrogate_type}
	SurrogateType *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateType `field:"optional" json:"surrogateType" yaml:"surrogateType"`
}

