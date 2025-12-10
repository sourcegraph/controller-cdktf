package escalationpath


type EscalationPathPathRepeat struct {
	// How many times to repeat these nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#repeat_times EscalationPath#repeat_times}
	RepeatTimes *float64 `field:"required" json:"repeatTimes" yaml:"repeatTimes"`
	// Which node ID we begin repeating from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#to_node EscalationPath#to_node}
	ToNode *string `field:"required" json:"toNode" yaml:"toNode"`
}

