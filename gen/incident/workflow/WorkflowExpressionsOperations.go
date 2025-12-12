package workflow


type WorkflowExpressionsOperations struct {
	// Indicates which operation type to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#operation_type Workflow#operation_type}
	OperationType *string `field:"required" json:"operationType" yaml:"operationType"`
	// An operation type that allows for a value to be set conditionally by a series of logical branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#branches Workflow#branches}
	Branches *WorkflowExpressionsOperationsBranches `field:"optional" json:"branches" yaml:"branches"`
	// An operation type that allows values to be filtered out by conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#filter Workflow#filter}
	Filter *WorkflowExpressionsOperationsFilter `field:"optional" json:"filter" yaml:"filter"`
	// An operation type that allows attributes of a type to be accessed by reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#navigate Workflow#navigate}
	Navigate *WorkflowExpressionsOperationsNavigate `field:"optional" json:"navigate" yaml:"navigate"`
	// An operation type that allows a value to parsed from within a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#parse Workflow#parse}
	Parse *WorkflowExpressionsOperationsParse `field:"optional" json:"parse" yaml:"parse"`
}

