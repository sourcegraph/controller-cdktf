package alertsource


type AlertSourceTemplateExpressionsOperationsBranchesBranchesResult struct {
	// The array of literal or reference parameter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#array_value AlertSource#array_value}
	ArrayValue interface{} `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The literal or reference parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#value AlertSource#value}
	Value *AlertSourceTemplateExpressionsOperationsBranchesBranchesResultValue `field:"optional" json:"value" yaml:"value"`
}

