package slo


type SloTimeWindowCalendar struct {
	// Date of the start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#start_time Slo#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Timezone name in IANA Time Zone Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#time_zone Slo#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

