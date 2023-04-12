package ratelimit


type RateLimitAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#mode RateLimit#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#response RateLimit#response}
	Response *RateLimitActionResponse `field:"optional" json:"response" yaml:"response"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#timeout RateLimit#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

