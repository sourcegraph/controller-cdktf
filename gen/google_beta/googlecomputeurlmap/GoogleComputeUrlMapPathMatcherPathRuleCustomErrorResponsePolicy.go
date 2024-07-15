package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherPathRuleCustomErrorResponsePolicy struct {
	// error_response_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#error_response_rule GoogleComputeUrlMap#error_response_rule}
	ErrorResponseRule interface{} `field:"optional" json:"errorResponseRule" yaml:"errorResponseRule"`
	// The full or partial URL to the BackendBucket resource that contains the custom error content. Examples are:.
	//
	// https://www.googleapis.com/compute/v1/projects/project/global/backendBuckets/myBackendBucket
	// compute/v1/projects/project/global/backendBuckets/myBackendBucket
	// global/backendBuckets/myBackendBucket
	//
	// If errorService is not specified at lower levels like pathMatcher, pathRule and routeRule, an errorService specified at a higher level in the UrlMap will be used. If UrlMap.defaultCustomErrorResponsePolicy contains one or more errorResponseRules[], it must specify errorService.
	// If load balancer cannot reach the backendBucket, a simple Not Found Error will be returned, with the original response code (or overrideResponseCode if configured).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#error_service GoogleComputeUrlMap#error_service}
	ErrorService *string `field:"optional" json:"errorService" yaml:"errorService"`
}

