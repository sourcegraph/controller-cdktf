package ratelimit


type RateLimitCorrelate struct {
	// If set to 'nat', NAT support will be enabled for rate limiting. Available values: `nat`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#by RateLimit#by}
	By *string `field:"optional" json:"by" yaml:"by"`
}

