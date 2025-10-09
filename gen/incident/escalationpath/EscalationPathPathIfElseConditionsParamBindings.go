package escalationpath


type EscalationPathPathIfElseConditionsParamBindings struct {
	// The array of literal or reference parameter values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#array_value EscalationPath#array_value}
	ArrayValue interface{} `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The literal or reference parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#value EscalationPath#value}
	Value *EscalationPathPathIfElseConditionsParamBindingsValue `field:"optional" json:"value" yaml:"value"`
}

