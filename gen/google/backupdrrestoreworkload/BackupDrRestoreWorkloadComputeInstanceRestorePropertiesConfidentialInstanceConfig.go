package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig struct {
	// Optional. Defines whether the instance should have confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_confidential_compute BackupDrRestoreWorkload#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

