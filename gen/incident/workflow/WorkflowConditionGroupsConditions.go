package workflow


type WorkflowConditionGroupsConditions struct {
	// The logical operation to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#operation Workflow#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Bindings for the operation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#param_bindings Workflow#param_bindings}
	ParamBindings interface{} `field:"required" json:"paramBindings" yaml:"paramBindings"`
	// The subject of the condition, on which the operation is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/workflow#subject Workflow#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

