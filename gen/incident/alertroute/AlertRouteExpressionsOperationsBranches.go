package alertroute

type AlertRouteExpressionsOperationsBranches struct {
	// The branches to apply for this operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#branches AlertRoute#branches}
	Branches any `field:"required" json:"branches" yaml:"branches"`
	// The return type of an operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#returns AlertRoute#returns}
	Returns *AlertRouteExpressionsOperationsBranchesReturns `field:"required" json:"returns" yaml:"returns"`
}
