package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateSchedulingGracefulShutdown struct {
	// Opts-in for graceful shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#enabled GoogleComputeInstanceFromTemplate#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#max_duration GoogleComputeInstanceFromTemplate#max_duration}
	MaxDuration *GoogleComputeInstanceFromTemplateSchedulingGracefulShutdownMaxDuration `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

