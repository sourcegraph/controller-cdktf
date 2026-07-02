package escalationpath

type EscalationPathPathIfElse struct {
	// The prerequisite conditions that must all be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#conditions EscalationPath#conditions}
	Conditions any `field:"required" json:"conditions" yaml:"conditions"`
	// Then path nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#then_path EscalationPath#then_path}
	ThenPath any `field:"required" json:"thenPath" yaml:"thenPath"`
	// The nodes that form the levels if our condition is not met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#else_path EscalationPath#else_path}
	ElsePath any `field:"optional" json:"elsePath" yaml:"elsePath"`
}
