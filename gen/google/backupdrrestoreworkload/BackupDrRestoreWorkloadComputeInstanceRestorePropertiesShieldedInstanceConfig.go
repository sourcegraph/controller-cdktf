package backupdrrestoreworkload

type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/backup_dr_restore_workload#enable_integrity_monitoring BackupDrRestoreWorkload#enable_integrity_monitoring}.
	EnableIntegrityMonitoring any `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/backup_dr_restore_workload#enable_secure_boot BackupDrRestoreWorkload#enable_secure_boot}.
	EnableSecureBoot any `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/backup_dr_restore_workload#enable_vtpm BackupDrRestoreWorkload#enable_vtpm}.
	EnableVtpm any `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}
