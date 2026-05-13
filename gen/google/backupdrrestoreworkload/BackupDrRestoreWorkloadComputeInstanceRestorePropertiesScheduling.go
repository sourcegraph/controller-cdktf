package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#automatic_restart BackupDrRestoreWorkload#automatic_restart}.
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Possible values: ["INSTANCE_TERMINATION_ACTION_UNSPECIFIED", "DELETE", "STOP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#instance_termination_action BackupDrRestoreWorkload#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#local_ssd_recovery_timeout BackupDrRestoreWorkload#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingLocalSsdRecoveryTimeout `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#max_run_duration BackupDrRestoreWorkload#max_run_duration}
	MaxRunDuration *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#min_node_cpus BackupDrRestoreWorkload#min_node_cpus}.
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#node_affinities BackupDrRestoreWorkload#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Possible values: ["ON_HOST_MAINTENANCE_UNSPECIFIED", "TERMINATE", "MIGRATE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#on_host_maintenance BackupDrRestoreWorkload#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#preemptible BackupDrRestoreWorkload#preemptible}.
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Possible values: ["PROVISIONING_MODEL_UNSPECIFIED", "STANDARD", "SPOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#provisioning_model BackupDrRestoreWorkload#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#termination_time BackupDrRestoreWorkload#termination_time}.
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

