package ratelimit


type RateLimitActionResponse struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#body RateLimit#body}.
	Body *string `field:"required" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#content_type RateLimit#content_type}.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
}

