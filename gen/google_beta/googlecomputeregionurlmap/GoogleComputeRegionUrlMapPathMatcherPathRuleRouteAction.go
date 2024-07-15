package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherPathRuleRouteAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#cors_policy GoogleComputeRegionUrlMap#cors_policy}
	CorsPolicy *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#fault_injection_policy GoogleComputeRegionUrlMap#fault_injection_policy}
	FaultInjectionPolicy *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#request_mirror_policy GoogleComputeRegionUrlMap#request_mirror_policy}
	RequestMirrorPolicy *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#retry_policy GoogleComputeRegionUrlMap#retry_policy}
	RetryPolicy *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#timeout GoogleComputeRegionUrlMap#timeout}
	Timeout *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#url_rewrite GoogleComputeRegionUrlMap#url_rewrite}
	UrlRewrite *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#weighted_backend_services GoogleComputeRegionUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

