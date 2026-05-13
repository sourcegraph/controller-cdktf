package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#location_selector GoogleComputeRolloutPlan#location_selector}
	LocationSelector *GoogleComputeRolloutPlanWavesSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#resource_hierarchy_selector GoogleComputeRolloutPlan#resource_hierarchy_selector}
	ResourceHierarchySelector *GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

