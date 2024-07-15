package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization struct {
	// Target number of concurrent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#target_concurrent_requests GoogleAppEngineFlexibleAppVersion#target_concurrent_requests}
	TargetConcurrentRequests *float64 `field:"optional" json:"targetConcurrentRequests" yaml:"targetConcurrentRequests"`
	// Target requests per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#target_request_count_per_second GoogleAppEngineFlexibleAppVersion#target_request_count_per_second}
	TargetRequestCountPerSecond *string `field:"optional" json:"targetRequestCountPerSecond" yaml:"targetRequestCountPerSecond"`
}

