package workflow


type WorkflowExpressionsOperationsBranchesBranchesResult struct {
	// The array of literal or reference parameter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#array_value Workflow#array_value}
	ArrayValue interface{} `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The literal or reference parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#value Workflow#value}
	Value *WorkflowExpressionsOperationsBranchesBranchesResultValue `field:"optional" json:"value" yaml:"value"`
}

