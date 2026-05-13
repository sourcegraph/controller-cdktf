package googlesaasruntimerolloutkind


type GoogleSaasRuntimeRolloutKindErrorBudget struct {
	// The maximum number of failed units allowed in a location without pausing the rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#allowed_count GoogleSaasRuntimeRolloutKind#allowed_count}
	AllowedCount *float64 `field:"optional" json:"allowedCount" yaml:"allowedCount"`
	// The maximum percentage of units allowed to fail (0, 100] within a location without pausing the rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#allowed_percentage GoogleSaasRuntimeRolloutKind#allowed_percentage}
	AllowedPercentage *float64 `field:"optional" json:"allowedPercentage" yaml:"allowedPercentage"`
}

