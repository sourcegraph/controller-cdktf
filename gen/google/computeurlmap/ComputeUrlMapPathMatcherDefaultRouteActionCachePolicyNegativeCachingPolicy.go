package computeurlmap


type ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyNegativeCachingPolicy struct {
	// The HTTP status code to define a TTL against.
	//
	// Only HTTP status codes
	// 300, 301, 302, 307, 308, 404, 405, 410, 421, 451 and 501 can be
	// specified as values, and you cannot specify a status code more than
	// once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#code ComputeUrlMap#code}
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#ttl ComputeUrlMap#ttl}
	Ttl *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyNegativeCachingPolicyTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

