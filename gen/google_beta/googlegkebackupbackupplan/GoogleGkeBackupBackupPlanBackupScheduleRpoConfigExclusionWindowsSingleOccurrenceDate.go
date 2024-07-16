package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate struct {
	// Day of a month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#day GoogleGkeBackupBackupPlan#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Month of a year.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#month GoogleGkeBackupBackupPlan#month}
	Month *float64 `field:"optional" json:"month" yaml:"month"`
	// Year of the date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#year GoogleGkeBackupBackupPlan#year}
	Year *float64 `field:"optional" json:"year" yaml:"year"`
}

