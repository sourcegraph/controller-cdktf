package escalationpath


type EscalationPathPathIfElseThenPath struct {
	// The type of this node.
	//
	// Available types are:
	// * level: A set of targets (users or schedules) that should be paged, either all at once, or with a round-robin configuration.
	// * notify_channel: Send the escalation to a Slack channel, where it can be acked by anyone in the channel.
	// * if_else: Branch the escalation based on a set of conditions.
	// * repeat: Go back to a previous node and repeat the logic from there.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#type EscalationPath#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An ID for this node, unique within the escalation path.
	//
	// This allows you to reference the node in other nodes, such as when configuring a 'repeat' node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#id EscalationPath#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#if_else EscalationPath#if_else}.
	IfElse *EscalationPathPathIfElse `field:"optional" json:"ifElse" yaml:"ifElse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#level EscalationPath#level}.
	Level *EscalationPathPathIfElseThenPathLevel `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#notify_channel EscalationPath#notify_channel}.
	NotifyChannel *EscalationPathPathIfElseThenPathNotifyChannel `field:"optional" json:"notifyChannel" yaml:"notifyChannel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#repeat EscalationPath#repeat}.
	Repeat *EscalationPathPathIfElseThenPathRepeat `field:"optional" json:"repeat" yaml:"repeat"`
}

