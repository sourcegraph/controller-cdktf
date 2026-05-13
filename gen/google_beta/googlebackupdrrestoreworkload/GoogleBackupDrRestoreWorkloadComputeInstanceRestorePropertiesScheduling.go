package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#automatic_restart GoogleBackupDrRestoreWorkload#automatic_restart}.
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Possible values: ["INSTANCE_TERMINATION_ACTION_UNSPECIFIED", "DELETE", "STOP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#instance_termination_action GoogleBackupDrRestoreWorkload#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#local_ssd_recovery_timeout GoogleBackupDrRestoreWorkload#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingLocalSsdRecoveryTimeout `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#max_run_duration GoogleBackupDrRestoreWorkload#max_run_duration}
	MaxRunDuration *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#min_node_cpus GoogleBackupDrRestoreWorkload#min_node_cpus}.
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#node_affinities GoogleBackupDrRestoreWorkload#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Possible values: ["ON_HOST_MAINTENANCE_UNSPECIFIED", "TERMINATE", "MIGRATE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#on_host_maintenance GoogleBackupDrRestoreWorkload#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#preemptible GoogleBackupDrRestoreWorkload#preemptible}.
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Possible values: ["PROVISIONING_MODEL_UNSPECIFIED", "STANDARD", "SPOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#provisioning_model GoogleBackupDrRestoreWorkload#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#termination_time GoogleBackupDrRestoreWorkload#termination_time}.
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

