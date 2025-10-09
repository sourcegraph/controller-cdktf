package workflow


type WorkflowExpressionsOperationsBranchesReturns struct {
	// Whether the return value should be single or multi-value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#array Workflow#array}
	Array interface{} `field:"required" json:"array" yaml:"array"`
	// Expected return type of this expression (what to try casting the result to).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#type Workflow#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

