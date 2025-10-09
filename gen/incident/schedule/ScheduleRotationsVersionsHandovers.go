package schedule


type ScheduleRotationsVersionsHandovers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/schedule#interval Schedule#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// How often a handover occurs. Possible values are: `hourly`, `daily`, `weekly`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/schedule#interval_type Schedule#interval_type}
	IntervalType *string `field:"required" json:"intervalType" yaml:"intervalType"`
}

