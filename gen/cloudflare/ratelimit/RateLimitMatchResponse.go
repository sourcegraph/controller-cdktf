package ratelimit


type RateLimitMatchResponse struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#headers RateLimit#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#origin_traffic RateLimit#origin_traffic}.
	OriginTraffic interface{} `field:"optional" json:"originTraffic" yaml:"originTraffic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#statuses RateLimit#statuses}.
	Statuses *[]*float64 `field:"optional" json:"statuses" yaml:"statuses"`
}

