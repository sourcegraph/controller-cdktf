package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherDefaultRouteAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#cors_policy GoogleComputeUrlMap#cors_policy}
	CorsPolicy *GoogleComputeUrlMapPathMatcherDefaultRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#fault_injection_policy GoogleComputeUrlMap#fault_injection_policy}
	FaultInjectionPolicy *GoogleComputeUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#request_mirror_policy GoogleComputeUrlMap#request_mirror_policy}
	RequestMirrorPolicy *GoogleComputeUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#retry_policy GoogleComputeUrlMap#retry_policy}
	RetryPolicy *GoogleComputeUrlMapPathMatcherDefaultRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#timeout GoogleComputeUrlMap#timeout}
	Timeout *GoogleComputeUrlMapPathMatcherDefaultRouteActionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#url_rewrite GoogleComputeUrlMap#url_rewrite}
	UrlRewrite *GoogleComputeUrlMapPathMatcherDefaultRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#weighted_backend_services GoogleComputeUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

