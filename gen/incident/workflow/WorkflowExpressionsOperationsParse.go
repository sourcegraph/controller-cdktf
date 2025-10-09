package workflow


type WorkflowExpressionsOperationsParse struct {
	// The return type of an operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#returns Workflow#returns}
	Returns *WorkflowExpressionsOperationsParseReturns `field:"required" json:"returns" yaml:"returns"`
	// The ES5 Javascript expression to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#source Workflow#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

