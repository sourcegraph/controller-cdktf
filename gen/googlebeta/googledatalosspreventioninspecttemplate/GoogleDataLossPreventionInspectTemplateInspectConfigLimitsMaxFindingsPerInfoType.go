package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#info_type GoogleDataLossPreventionInspectTemplate#info_type}
	InfoType *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `field:"required" json:"infoType" yaml:"infoType"`
	// Max findings limit for the given infoType.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#max_findings GoogleDataLossPreventionInspectTemplate#max_findings}
	MaxFindings *float64 `field:"required" json:"maxFindings" yaml:"maxFindings"`
}

