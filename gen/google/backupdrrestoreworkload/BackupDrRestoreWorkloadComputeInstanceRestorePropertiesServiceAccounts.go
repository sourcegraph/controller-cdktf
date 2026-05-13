package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#email BackupDrRestoreWorkload#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#scopes BackupDrRestoreWorkload#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

