package workflow

type WorkflowExpressionsElseBranchResult struct {
	// The array of literal or reference parameter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#array_value Workflow#array_value}
	ArrayValue any `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The literal or reference parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#value Workflow#value}
	Value *WorkflowExpressionsElseBranchResultValue `field:"optional" json:"value" yaml:"value"`
}
