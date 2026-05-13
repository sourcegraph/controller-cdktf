package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_integrity_monitoring BackupDrRestoreWorkload#enable_integrity_monitoring}.
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_secure_boot BackupDrRestoreWorkload#enable_secure_boot}.
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_vtpm BackupDrRestoreWorkload#enable_vtpm}.
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

