package ratelimit


type RateLimitAction struct {
	// The type of action to perform. Available values: `simulate`, `ban`, `challenge`, `js_challenge`, `managed_challenge`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#mode RateLimit#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#response RateLimit#response}
	Response *RateLimitActionResponse `field:"optional" json:"response" yaml:"response"`
	// The time in seconds as an integer to perform the mitigation action.
	//
	// This field is required if the `mode` is either `simulate` or `ban`. Must be the same or greater than the period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#timeout RateLimit#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

