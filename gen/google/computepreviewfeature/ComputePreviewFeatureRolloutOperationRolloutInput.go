package computepreviewfeature


type ComputePreviewFeatureRolloutOperationRolloutInput struct {
	// Predefined rollout plans. Possible values: ["ROLLOUT_PLAN_FAST_ROLLOUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#predefined_rollout_plan ComputePreviewFeature#predefined_rollout_plan}
	PredefinedRolloutPlan *string `field:"required" json:"predefinedRolloutPlan" yaml:"predefinedRolloutPlan"`
}

