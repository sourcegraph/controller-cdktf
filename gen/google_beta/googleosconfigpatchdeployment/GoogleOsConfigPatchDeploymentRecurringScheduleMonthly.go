package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentRecurringScheduleMonthly struct {
	// One day of the month.
	//
	// 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.
	// Months without the target day will be skipped. For example, a schedule to run "every month on the 31st"
	// will not run in February, April, June, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#month_day GoogleOsConfigPatchDeployment#month_day}
	MonthDay *float64 `field:"optional" json:"monthDay" yaml:"monthDay"`
	// week_day_of_month block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#week_day_of_month GoogleOsConfigPatchDeployment#week_day_of_month}
	WeekDayOfMonth *GoogleOsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth `field:"optional" json:"weekDayOfMonth" yaml:"weekDayOfMonth"`
}

