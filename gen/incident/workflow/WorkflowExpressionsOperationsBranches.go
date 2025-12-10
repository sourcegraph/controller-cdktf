package workflow


type WorkflowExpressionsOperationsBranches struct {
	// The branches to apply for this operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#branches Workflow#branches}
	Branches interface{} `field:"required" json:"branches" yaml:"branches"`
	// The return type of an operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#returns Workflow#returns}
	Returns *WorkflowExpressionsOperationsBranchesReturns `field:"required" json:"returns" yaml:"returns"`
}

