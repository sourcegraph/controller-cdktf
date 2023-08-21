package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupSchedule struct {
	// A standard cron string that defines a repeating schedule for creating Backups via this BackupPlan.
	//
	// If this is defined, then backupRetainDays must also be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_gke_backup_backup_plan#cron_schedule GoogleGkeBackupBackupPlan#cron_schedule}
	CronSchedule *string `field:"optional" json:"cronSchedule" yaml:"cronSchedule"`
	// This flag denotes whether automatic Backup creation is paused for this BackupPlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_gke_backup_backup_plan#paused GoogleGkeBackupBackupPlan#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
}

