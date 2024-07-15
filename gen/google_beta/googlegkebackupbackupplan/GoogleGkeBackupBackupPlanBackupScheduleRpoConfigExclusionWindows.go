package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindows struct {
	// Specifies duration of the window in seconds with up to nine fractional digits, terminated by 's'.
	//
	// Example: "3.5s". Restrictions for duration based on the
	// recurrence type to allow some time for backup to happen:
	//   - single_occurrence_date:  no restriction
	//   - daily window: duration < 24 hours
	//   - weekly window:
	//     - days of week includes all seven days of a week: duration < 24 hours
	//     - all other weekly window: duration < 168 hours (i.e., 24 * 7 hours)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#duration GoogleGkeBackupBackupPlan#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#start_time GoogleGkeBackupBackupPlan#start_time}
	StartTime *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
	// The exclusion window occurs every day if set to "True".
	//
	// Specifying this field to "False" is an error.
	// Only one of singleOccurrenceDate, daily and daysOfWeek may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#daily GoogleGkeBackupBackupPlan#daily}
	Daily interface{} `field:"optional" json:"daily" yaml:"daily"`
	// days_of_week block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#days_of_week GoogleGkeBackupBackupPlan#days_of_week}
	DaysOfWeek *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// single_occurrence_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#single_occurrence_date GoogleGkeBackupBackupPlan#single_occurrence_date}
	SingleOccurrenceDate *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate `field:"optional" json:"singleOccurrenceDate" yaml:"singleOccurrenceDate"`
}

