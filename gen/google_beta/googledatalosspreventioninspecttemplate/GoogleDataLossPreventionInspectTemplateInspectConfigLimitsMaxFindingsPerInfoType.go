package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType struct {
	// Max findings limit for the given infoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#max_findings GoogleDataLossPreventionInspectTemplate#max_findings}
	MaxFindings *float64 `field:"required" json:"maxFindings" yaml:"maxFindings"`
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#info_type GoogleDataLossPreventionInspectTemplate#info_type}
	InfoType *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `field:"optional" json:"infoType" yaml:"infoType"`
}

