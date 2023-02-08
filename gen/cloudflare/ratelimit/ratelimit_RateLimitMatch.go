package ratelimit


type RateLimitMatch struct {
	// request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#request RateLimit#request}
	Request *RateLimitMatchRequest `field:"optional" json:"request" yaml:"request"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#response RateLimit#response}
	Response *RateLimitMatchResponse `field:"optional" json:"response" yaml:"response"`
}

