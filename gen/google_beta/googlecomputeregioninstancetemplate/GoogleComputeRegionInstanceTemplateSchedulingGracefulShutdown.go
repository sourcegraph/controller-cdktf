package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateSchedulingGracefulShutdown struct {
	// Opts-in for graceful shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_instance_template#enabled GoogleComputeRegionInstanceTemplate#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_instance_template#max_duration GoogleComputeRegionInstanceTemplate#max_duration}
	MaxDuration *GoogleComputeRegionInstanceTemplateSchedulingGracefulShutdownMaxDuration `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

