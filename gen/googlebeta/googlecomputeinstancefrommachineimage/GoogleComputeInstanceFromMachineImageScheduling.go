package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageScheduling struct {
	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#automatic_restart GoogleComputeInstanceFromMachineImage#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#instance_termination_action GoogleComputeInstanceFromMachineImage#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#min_node_cpus GoogleComputeInstanceFromMachineImage#min_node_cpus}.
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#node_affinities GoogleComputeInstanceFromMachineImage#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#on_host_maintenance GoogleComputeInstanceFromMachineImage#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// Whether the instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#preemptible GoogleComputeInstanceFromMachineImage#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_machine_image#provisioning_model GoogleComputeInstanceFromMachineImage#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
}

