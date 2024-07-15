package googlelookerinstance


type GoogleLookerInstanceMaintenanceWindow struct {
	// Required. Day of the week for this MaintenanceWindow (in UTC).
	//
	// - MONDAY: Monday
	// - TUESDAY: Tuesday
	// - WEDNESDAY: Wednesday
	// - THURSDAY: Thursday
	// - FRIDAY: Friday
	// - SATURDAY: Saturday
	// - SUNDAY: Sunday Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#day_of_week GoogleLookerInstance#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#start_time GoogleLookerInstance#start_time}
	StartTime *GoogleLookerInstanceMaintenanceWindowStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

