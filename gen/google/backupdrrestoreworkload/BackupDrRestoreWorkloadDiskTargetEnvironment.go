package backupdrrestoreworkload


type BackupDrRestoreWorkloadDiskTargetEnvironment struct {
	// Required. Target project for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#project BackupDrRestoreWorkload#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Required. Target zone for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#zone BackupDrRestoreWorkload#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
}

