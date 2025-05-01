package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageSchedulingGracefulShutdown struct {
	// Opts-in for graceful shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_machine_image#enabled GoogleComputeInstanceFromMachineImage#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_machine_image#max_duration GoogleComputeInstanceFromMachineImage#max_duration}
	MaxDuration *GoogleComputeInstanceFromMachineImageSchedulingGracefulShutdownMaxDuration `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

