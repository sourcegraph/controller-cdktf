package backupdrbackupplan


type BackupDrBackupPlanDiskBackupPlanProperties struct {
	// Indicates whether to perform a guest flush operation before taking a disk backup.
	//
	// When set to true, the system will attempt to ensure
	// application-consistent backups. When set to false, the system will
	// create crash-consistent backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_backup_plan#guest_flush BackupDrBackupPlan#guest_flush}
	GuestFlush interface{} `field:"required" json:"guestFlush" yaml:"guestFlush"`
}

