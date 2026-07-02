package escalationpath

type EscalationPathPathIfElseConditions struct {
	// The logical operation to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#operation EscalationPath#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Bindings for the operation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#param_bindings EscalationPath#param_bindings}
	ParamBindings any `field:"required" json:"paramBindings" yaml:"paramBindings"`
	// The subject of the condition, on which the operation is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#subject EscalationPath#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}
