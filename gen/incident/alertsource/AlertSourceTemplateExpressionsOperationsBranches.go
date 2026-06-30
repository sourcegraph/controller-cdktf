package alertsource

type AlertSourceTemplateExpressionsOperationsBranches struct {
	// The branches to apply for this operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#branches AlertSource#branches}
	Branches any `field:"required" json:"branches" yaml:"branches"`
	// The return type of an operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#returns AlertSource#returns}
	Returns *AlertSourceTemplateExpressionsOperationsBranchesReturns `field:"required" json:"returns" yaml:"returns"`
}
