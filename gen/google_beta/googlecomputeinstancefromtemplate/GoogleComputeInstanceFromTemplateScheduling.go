package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateScheduling struct {
	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#automatic_restart GoogleComputeInstanceFromTemplate#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#instance_termination_action GoogleComputeInstanceFromTemplate#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#local_ssd_recovery_timeout GoogleComputeInstanceFromTemplate#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout *GoogleComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#maintenance_interval GoogleComputeInstanceFromTemplate#maintenance_interval}
	MaintenanceInterval *string `field:"optional" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#max_run_duration GoogleComputeInstanceFromTemplate#max_run_duration}
	MaxRunDuration *GoogleComputeInstanceFromTemplateSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#min_node_cpus GoogleComputeInstanceFromTemplate#min_node_cpus}.
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#node_affinities GoogleComputeInstanceFromTemplate#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#on_host_maintenance GoogleComputeInstanceFromTemplate#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#on_instance_stop_action GoogleComputeInstanceFromTemplate#on_instance_stop_action}
	OnInstanceStopAction *GoogleComputeInstanceFromTemplateSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Whether the instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#preemptible GoogleComputeInstanceFromTemplate#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#provisioning_model GoogleComputeInstanceFromTemplate#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
}

