package ratelimit


type RateLimitMatchResponse struct {
	// List of HTTP headers maps to match the origin response on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#headers RateLimit#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Only count traffic that has come from your origin servers.
	//
	// If true, cached items that Cloudflare serve will not count towards rate limiting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#origin_traffic RateLimit#origin_traffic}
	OriginTraffic interface{} `field:"optional" json:"originTraffic" yaml:"originTraffic"`
	// HTTP Status codes, can be one, many or indicate all by not providing this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#statuses RateLimit#statuses}
	Statuses *[]*float64 `field:"optional" json:"statuses" yaml:"statuses"`
}

