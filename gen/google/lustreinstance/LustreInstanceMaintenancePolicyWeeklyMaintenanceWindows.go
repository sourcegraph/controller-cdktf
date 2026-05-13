package lustreinstance


type LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows struct {
	// Possible values: MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#day_of_week LustreInstance#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#start_time LustreInstance#start_time}
	StartTime *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

