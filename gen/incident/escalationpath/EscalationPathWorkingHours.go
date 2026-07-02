package escalationpath

type EscalationPathWorkingHours struct {
	// The unique identifier for this set of working intervals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#id EscalationPath#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A human readable label for this set of working intervals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#name EscalationPath#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// How to interpret all the intervals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#timezone EscalationPath#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#weekday_intervals EscalationPath#weekday_intervals}.
	WeekdayIntervals any `field:"required" json:"weekdayIntervals" yaml:"weekdayIntervals"`
}
