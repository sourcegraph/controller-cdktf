package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesOrchestrationOptionsDelays struct {
	// Controls whether the delay should only be added between batches of projects corresponding to different locations, or also between batches of projects corresponding to the same location.
	//
	// Possible values: ["DELIMITER_UNSPECIFIED", "DELIMITER_LOCATION", "DELIMITER_BATCH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#delimiter GoogleComputeRolloutPlan#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The duration of the delay, if any, to be added between batches of projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#duration GoogleComputeRolloutPlan#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Controls whether the specified duration is to be added at the end of each batch, or if the total processing time for each batch will be padded if needed to meet the specified duration.
	//
	// Possible values: ["TYPE_UNSPECIFIED", "TYPE_OFFSET", "TYPE_MINIMUM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#type GoogleComputeRolloutPlan#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

