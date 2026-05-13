package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#key GoogleBackupDrRestoreWorkload#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Possible values: ["OPERATOR_UNSPECIFIED", "IN", "NOT_IN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#operator GoogleBackupDrRestoreWorkload#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#values GoogleBackupDrRestoreWorkload#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

