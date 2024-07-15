package googlelookerinstance


type GoogleLookerInstanceDenyMaintenancePeriod struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#end_date GoogleLookerInstance#end_date}
	EndDate *GoogleLookerInstanceDenyMaintenancePeriodEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#start_date GoogleLookerInstance#start_date}
	StartDate *GoogleLookerInstanceDenyMaintenancePeriodStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#time GoogleLookerInstance#time}
	Time *GoogleLookerInstanceDenyMaintenancePeriodTime `field:"required" json:"time" yaml:"time"`
}

