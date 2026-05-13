package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesSelectorsLocationSelector struct {
	// Example: "us-central1-a".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#included_locations GoogleComputeRolloutPlan#included_locations}
	IncludedLocations *[]*string `field:"optional" json:"includedLocations" yaml:"includedLocations"`
}

