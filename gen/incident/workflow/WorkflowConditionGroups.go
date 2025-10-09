package workflow


type WorkflowConditionGroups struct {
	// The prerequisite conditions that must all be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#conditions Workflow#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
}

