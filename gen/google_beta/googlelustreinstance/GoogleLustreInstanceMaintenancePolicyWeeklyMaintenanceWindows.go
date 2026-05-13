package googlelustreinstance


type GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows struct {
	// Possible values: MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#day_of_week GoogleLustreInstance#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#start_time GoogleLustreInstance#start_time}
	StartTime *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

