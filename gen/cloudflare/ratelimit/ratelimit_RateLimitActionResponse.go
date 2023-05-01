package ratelimit


type RateLimitActionResponse struct {
	// The body to return, the content here should conform to the `content_type`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#body RateLimit#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// The content-type of the body. Available values: `text/plain`, `text/xml`, `application/json`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#content_type RateLimit#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
}

