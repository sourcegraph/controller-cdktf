package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity struct {
	// Possible values: ["TYPE_UNSPECIFIED", "NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#consume_allocation_type BackupDrRestoreWorkload#consume_allocation_type}
	ConsumeAllocationType *string `field:"optional" json:"consumeAllocationType" yaml:"consumeAllocationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#key BackupDrRestoreWorkload#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#values BackupDrRestoreWorkload#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

