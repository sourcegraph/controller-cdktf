package workflow


type WorkflowExpressions struct {
	// The human readable label of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#label Workflow#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The operations to execute in sequence for this expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#operations Workflow#operations}
	Operations interface{} `field:"required" json:"operations" yaml:"operations"`
	// A short ID that can be used to reference the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#reference Workflow#reference}
	Reference *string `field:"required" json:"reference" yaml:"reference"`
	// The root reference for this expression (i.e. where the expression starts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#root_reference Workflow#root_reference}
	RootReference *string `field:"required" json:"rootReference" yaml:"rootReference"`
	// The else branch to resort to if all operations fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#else_branch Workflow#else_branch}
	ElseBranch *WorkflowExpressionsElseBranch `field:"optional" json:"elseBranch" yaml:"elseBranch"`
}

