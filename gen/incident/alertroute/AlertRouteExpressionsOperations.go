package alertroute


type AlertRouteExpressionsOperations struct {
	// Indicates which operation type to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#operation_type AlertRoute#operation_type}
	OperationType *string `field:"required" json:"operationType" yaml:"operationType"`
	// An operation type that allows for a value to be set conditionally by a series of logical branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#branches AlertRoute#branches}
	Branches *AlertRouteExpressionsOperationsBranches `field:"optional" json:"branches" yaml:"branches"`
	// An operation type that allows values to be filtered out by conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#filter AlertRoute#filter}
	Filter *AlertRouteExpressionsOperationsFilter `field:"optional" json:"filter" yaml:"filter"`
	// An operation type that allows attributes of a type to be accessed by reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#navigate AlertRoute#navigate}
	Navigate *AlertRouteExpressionsOperationsNavigate `field:"optional" json:"navigate" yaml:"navigate"`
	// An operation type that allows a value to parsed from within a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#parse AlertRoute#parse}
	Parse *AlertRouteExpressionsOperationsParse `field:"optional" json:"parse" yaml:"parse"`
}

