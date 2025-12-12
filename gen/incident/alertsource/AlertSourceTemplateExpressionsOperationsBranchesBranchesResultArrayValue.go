package alertsource


type AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValue struct {
	// If set, this is the literal value of the step parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#literal AlertSource#literal}
	Literal *string `field:"optional" json:"literal" yaml:"literal"`
	// If set, this is the reference into the trigger scope that is the value of this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#reference AlertSource#reference}
	Reference *string `field:"optional" json:"reference" yaml:"reference"`
}

