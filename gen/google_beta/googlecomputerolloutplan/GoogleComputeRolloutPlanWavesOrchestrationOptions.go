package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesOrchestrationOptions struct {
	// delays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#delays GoogleComputeRolloutPlan#delays}
	Delays interface{} `field:"optional" json:"delays" yaml:"delays"`
	// Maximum number of locations to be orchestrated in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#max_concurrent_locations GoogleComputeRolloutPlan#max_concurrent_locations}
	MaxConcurrentLocations *float64 `field:"optional" json:"maxConcurrentLocations" yaml:"maxConcurrentLocations"`
	// Maximum number of resources to be orchestrated per location in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#max_concurrent_resources_per_location GoogleComputeRolloutPlan#max_concurrent_resources_per_location}
	MaxConcurrentResourcesPerLocation *float64 `field:"optional" json:"maxConcurrentResourcesPerLocation" yaml:"maxConcurrentResourcesPerLocation"`
}

