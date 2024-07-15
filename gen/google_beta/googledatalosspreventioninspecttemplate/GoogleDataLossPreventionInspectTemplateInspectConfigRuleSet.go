package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSet struct {
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#info_types GoogleDataLossPreventionInspectTemplate#info_types}
	InfoTypes interface{} `field:"required" json:"infoTypes" yaml:"infoTypes"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#rules GoogleDataLossPreventionInspectTemplate#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

