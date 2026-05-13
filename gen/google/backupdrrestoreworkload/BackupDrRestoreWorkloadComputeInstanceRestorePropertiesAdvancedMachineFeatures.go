package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures struct {
	// Optional. Whether to enable nested virtualization or not (default is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_nested_virtualization BackupDrRestoreWorkload#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
	// Optional. Whether to enable UEFI networking for instance creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_uefi_networking BackupDrRestoreWorkload#enable_uefi_networking}
	EnableUefiNetworking interface{} `field:"optional" json:"enableUefiNetworking" yaml:"enableUefiNetworking"`
	// Optional. The number of threads per physical core.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#threads_per_core BackupDrRestoreWorkload#threads_per_core}
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// Optional. The number of physical cores to expose to an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#visible_core_count BackupDrRestoreWorkload#visible_core_count}
	VisibleCoreCount *float64 `field:"optional" json:"visibleCoreCount" yaml:"visibleCoreCount"`
}

