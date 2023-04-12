package ratelimit


type RateLimitMatchRequest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#methods RateLimit#methods}.
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#schemes RateLimit#schemes}.
	Schemes *[]*string `field:"optional" json:"schemes" yaml:"schemes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#url_pattern RateLimit#url_pattern}.
	UrlPattern *string `field:"optional" json:"urlPattern" yaml:"urlPattern"`
}

