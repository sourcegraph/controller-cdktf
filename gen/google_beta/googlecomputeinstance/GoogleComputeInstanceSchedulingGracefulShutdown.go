package googlecomputeinstance


type GoogleComputeInstanceSchedulingGracefulShutdown struct {
	// Opts-in for graceful shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance#enabled GoogleComputeInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance#max_duration GoogleComputeInstance#max_duration}
	MaxDuration *GoogleComputeInstanceSchedulingGracefulShutdownMaxDuration `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

