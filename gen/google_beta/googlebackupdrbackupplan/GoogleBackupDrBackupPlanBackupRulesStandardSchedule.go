package googlebackupdrbackupplan


type GoogleBackupDrBackupPlanBackupRulesStandardSchedule struct {
	// RecurrenceType enumerates the applicable periodicity for the schedule. Possible values: ["HOURLY", "DAILY", "WEEKLY", "MONTHLY", "YEARLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#recurrence_type GoogleBackupDrBackupPlan#recurrence_type}
	RecurrenceType *string `field:"required" json:"recurrenceType" yaml:"recurrenceType"`
	// The time zone to be used when interpreting the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#time_zone GoogleBackupDrBackupPlan#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// backup_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#backup_window GoogleBackupDrBackupPlan#backup_window}
	BackupWindow *GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindow `field:"optional" json:"backupWindow" yaml:"backupWindow"`
	// Specifies days of months like 1, 5, or 14 on which jobs will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#days_of_month GoogleBackupDrBackupPlan#days_of_month}
	DaysOfMonth *[]*float64 `field:"optional" json:"daysOfMonth" yaml:"daysOfMonth"`
	// Specifies days of week like MONDAY or TUESDAY, on which jobs will run.
	//
	// This is required for 'recurrence_type', 'WEEKLY' and is not applicable otherwise. Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#days_of_week GoogleBackupDrBackupPlan#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Specifies frequency for hourly backups.
	//
	// An hourly frequency of 2 means jobs will run every 2 hours from start time till end time defined.
	// This is required for 'recurrence_type', 'HOURLY' and is not applicable otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#hourly_frequency GoogleBackupDrBackupPlan#hourly_frequency}
	HourlyFrequency *float64 `field:"optional" json:"hourlyFrequency" yaml:"hourlyFrequency"`
	// Specifies values of months Possible values: ["MONTH_UNSPECIFIED", "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#months GoogleBackupDrBackupPlan#months}
	Months *[]*string `field:"optional" json:"months" yaml:"months"`
	// week_day_of_month block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#week_day_of_month GoogleBackupDrBackupPlan#week_day_of_month}
	WeekDayOfMonth *GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth `field:"optional" json:"weekDayOfMonth" yaml:"weekDayOfMonth"`
}

