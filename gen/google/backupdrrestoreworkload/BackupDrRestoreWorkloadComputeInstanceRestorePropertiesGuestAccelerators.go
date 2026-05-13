package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAccelerators struct {
	// Optional. The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#accelerator_count BackupDrRestoreWorkload#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Optional. Full or partial URL of the accelerator type resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#accelerator_type BackupDrRestoreWorkload#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
}

