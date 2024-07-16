package googlefirestorebackupschedule


type GoogleFirestoreBackupScheduleWeeklyRecurrence struct {
	// The day of week to run. Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#day GoogleFirestoreBackupSchedule#day}
	Day *string `field:"optional" json:"day" yaml:"day"`
}

