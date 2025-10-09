package escalationpath


type EscalationPathPathIfElseThenPathNotifyChannel struct {
	// The targets (Slack channels) for this level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#targets EscalationPath#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// If the time to ack is relative to a time window, this defines whether we move when the window is active or inactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#time_to_ack_interval_condition EscalationPath#time_to_ack_interval_condition}
	TimeToAckIntervalCondition *string `field:"optional" json:"timeToAckIntervalCondition" yaml:"timeToAckIntervalCondition"`
	// How long should we wait for this level to acknowledge before moving on to the next node in the path?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#time_to_ack_seconds EscalationPath#time_to_ack_seconds}
	TimeToAckSeconds *float64 `field:"optional" json:"timeToAckSeconds" yaml:"timeToAckSeconds"`
	// If the time to ack is relative to a time window, this identifies which window it is relative to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#time_to_ack_weekday_interval_config_id EscalationPath#time_to_ack_weekday_interval_config_id}
	TimeToAckWeekdayIntervalConfigId *string `field:"optional" json:"timeToAckWeekdayIntervalConfigId" yaml:"timeToAckWeekdayIntervalConfigId"`
}

