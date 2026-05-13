package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadDiskTargetEnvironment struct {
	// Required. Target project for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#project GoogleBackupDrRestoreWorkload#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Required. Target zone for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#zone GoogleBackupDrRestoreWorkload#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
}

