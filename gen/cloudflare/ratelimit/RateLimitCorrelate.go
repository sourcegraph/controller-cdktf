package ratelimit


type RateLimitCorrelate struct {
	// If set to 'nat', NAT support will be enabled for rate limiting. Available values: `nat`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#by RateLimit#by}
	By *string `field:"optional" json:"by" yaml:"by"`
}

