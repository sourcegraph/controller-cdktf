package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateScheduling struct {
	// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user).
	//
	// This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#automatic_restart GoogleComputeInstanceTemplate#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#instance_termination_action GoogleComputeInstanceTemplate#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#local_ssd_recovery_timeout GoogleComputeInstanceTemplate#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout interface{} `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#maintenance_interval GoogleComputeInstanceTemplate#maintenance_interval}
	MaintenanceInterval *string `field:"optional" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#max_run_duration GoogleComputeInstanceTemplate#max_run_duration}
	MaxRunDuration *GoogleComputeInstanceTemplateSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Minimum number of cpus for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#min_node_cpus GoogleComputeInstanceTemplate#min_node_cpus}
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#node_affinities GoogleComputeInstanceTemplate#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Defines the maintenance behavior for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#on_host_maintenance GoogleComputeInstanceTemplate#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#on_instance_stop_action GoogleComputeInstanceTemplate#on_instance_stop_action}
	OnInstanceStopAction *GoogleComputeInstanceTemplateSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Allows instance to be preempted. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#preemptible GoogleComputeInstanceTemplate#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#provisioning_model GoogleComputeInstanceTemplate#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
}

