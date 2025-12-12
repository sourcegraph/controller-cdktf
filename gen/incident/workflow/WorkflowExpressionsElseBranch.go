package workflow


type WorkflowExpressionsElseBranch struct {
	// The result assumed if the else branch is reached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#result Workflow#result}
	Result *WorkflowExpressionsElseBranchResult `field:"required" json:"result" yaml:"result"`
}

