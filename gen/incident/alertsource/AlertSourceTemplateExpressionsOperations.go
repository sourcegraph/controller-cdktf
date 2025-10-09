package alertsource


type AlertSourceTemplateExpressionsOperations struct {
	// Indicates which operation type to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#operation_type AlertSource#operation_type}
	OperationType *string `field:"required" json:"operationType" yaml:"operationType"`
	// An operation type that allows for a value to be set conditionally by a series of logical branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#branches AlertSource#branches}
	Branches *AlertSourceTemplateExpressionsOperationsBranches `field:"optional" json:"branches" yaml:"branches"`
	// An operation type that allows values to be filtered out by conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#filter AlertSource#filter}
	Filter *AlertSourceTemplateExpressionsOperationsFilter `field:"optional" json:"filter" yaml:"filter"`
	// An operation type that allows attributes of a type to be accessed by reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#navigate AlertSource#navigate}
	Navigate *AlertSourceTemplateExpressionsOperationsNavigate `field:"optional" json:"navigate" yaml:"navigate"`
	// An operation type that allows a value to parsed from within a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_source#parse AlertSource#parse}
	Parse *AlertSourceTemplateExpressionsOperationsParse `field:"optional" json:"parse" yaml:"parse"`
}

