package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime struct {
	// Hours of day in 24 hour format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#hours GoogleGkeBackupBackupPlan#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of hour of day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#minutes GoogleGkeBackupBackupPlan#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#nanos GoogleGkeBackupBackupPlan#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#seconds GoogleGkeBackupBackupPlan#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

