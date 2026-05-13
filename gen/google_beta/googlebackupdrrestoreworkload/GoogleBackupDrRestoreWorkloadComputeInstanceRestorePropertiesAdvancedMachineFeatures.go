package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures struct {
	// Optional. Whether to enable nested virtualization or not (default is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#enable_nested_virtualization GoogleBackupDrRestoreWorkload#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
	// Optional. Whether to enable UEFI networking for instance creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#enable_uefi_networking GoogleBackupDrRestoreWorkload#enable_uefi_networking}
	EnableUefiNetworking interface{} `field:"optional" json:"enableUefiNetworking" yaml:"enableUefiNetworking"`
	// Optional. The number of threads per physical core.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#threads_per_core GoogleBackupDrRestoreWorkload#threads_per_core}
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// Optional. The number of physical cores to expose to an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#visible_core_count GoogleBackupDrRestoreWorkload#visible_core_count}
	VisibleCoreCount *float64 `field:"optional" json:"visibleCoreCount" yaml:"visibleCoreCount"`
}

