package escalationpath

type EscalationPathPathLevel struct {
	// The targets (users or schedules) for this level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#targets EscalationPath#targets}
	Targets any `field:"required" json:"targets" yaml:"targets"`
	// Controls the behaviour of acknowledgements for this level, with 'first' cancelling all other escalations on the same level when someone acks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#ack_mode EscalationPath#ack_mode}
	AckMode *string `field:"optional" json:"ackMode" yaml:"ackMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#round_robin_config EscalationPath#round_robin_config}.
	RoundRobinConfig *EscalationPathPathLevelRoundRobinConfig `field:"optional" json:"roundRobinConfig" yaml:"roundRobinConfig"`
	// If the time to ack is relative to a time window, this defines whether we move when the window is active or inactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#time_to_ack_interval_condition EscalationPath#time_to_ack_interval_condition}
	TimeToAckIntervalCondition *string `field:"optional" json:"timeToAckIntervalCondition" yaml:"timeToAckIntervalCondition"`
	// How long should we wait for this level to acknowledge before proceeding to the next node in the path?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#time_to_ack_seconds EscalationPath#time_to_ack_seconds}
	TimeToAckSeconds *float64 `field:"optional" json:"timeToAckSeconds" yaml:"timeToAckSeconds"`
	// If the time to ack is relative to a time window, this identifies which window it is relative to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#time_to_ack_weekday_interval_config_id EscalationPath#time_to_ack_weekday_interval_config_id}
	TimeToAckWeekdayIntervalConfigId *string `field:"optional" json:"timeToAckWeekdayIntervalConfigId" yaml:"timeToAckWeekdayIntervalConfigId"`
}
