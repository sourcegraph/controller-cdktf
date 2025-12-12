package schedule


type ScheduleRotationsVersionsWorkingIntervals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#end_time Schedule#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#start_time Schedule#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#weekday Schedule#weekday}.
	Weekday *string `field:"required" json:"weekday" yaml:"weekday"`
}

