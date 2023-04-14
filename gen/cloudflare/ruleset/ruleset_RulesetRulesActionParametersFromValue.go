package ruleset


type RulesetRulesActionParametersFromValue struct {
	// Preserve query string for redirect URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#preserve_query_string Ruleset#preserve_query_string}
	PreserveQueryString interface{} `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// Status code for redirect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// target_url block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#target_url Ruleset#target_url}
	TargetUrl interface{} `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}

