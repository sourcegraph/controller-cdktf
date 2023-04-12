package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth struct {
	// A day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#day_of_week GoogleOsConfigPatchDeployment#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Week number in a month.
	//
	// 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#week_ordinal GoogleOsConfigPatchDeployment#week_ordinal}
	WeekOrdinal *float64 `field:"required" json:"weekOrdinal" yaml:"weekOrdinal"`
}

