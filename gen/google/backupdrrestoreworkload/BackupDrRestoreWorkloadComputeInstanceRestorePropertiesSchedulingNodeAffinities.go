package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#key BackupDrRestoreWorkload#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Possible values: ["OPERATOR_UNSPECIFIED", "IN", "NOT_IN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#operator BackupDrRestoreWorkload#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#values BackupDrRestoreWorkload#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

