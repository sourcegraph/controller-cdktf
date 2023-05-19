package ratelimit


type RateLimitMatchRequest struct {
	// HTTP Methods to match traffic on. Available values: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#methods RateLimit#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// HTTP schemes to match traffic on. Available values: `HTTP`, `HTTPS`, `_ALL_`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#schemes RateLimit#schemes}
	Schemes *[]*string `field:"optional" json:"schemes" yaml:"schemes"`
	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#url_pattern RateLimit#url_pattern}
	UrlPattern *string `field:"optional" json:"urlPattern" yaml:"urlPattern"`
}

