package escalationpath


type EscalationPathWorkingHoursWeekdayIntervals struct {
	// End time of the interval, in 24hr format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#end_time EscalationPath#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Start time of the interval, in 24hr format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#start_time EscalationPath#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Weekdays for use within a schedule or escalation path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#weekday EscalationPath#weekday}
	Weekday *string `field:"required" json:"weekday" yaml:"weekday"`
}

