package ruleset


type RulesetRulesRatelimit struct {
	// List of parameters that define how Cloudflare tracks the request rate for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#characteristics Ruleset#characteristics}
	Characteristics *[]*string `field:"optional" json:"characteristics" yaml:"characteristics"`
	// Criteria for counting HTTP requests to trigger the Rate Limiting action.
	//
	// Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#counting_expression Ruleset#counting_expression}
	CountingExpression *string `field:"optional" json:"countingExpression" yaml:"countingExpression"`
	// Once the request rate is reached, the Rate Limiting rule blocks further requests for the period of time defined in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#mitigation_timeout Ruleset#mitigation_timeout}
	MitigationTimeout *float64 `field:"optional" json:"mitigationTimeout" yaml:"mitigationTimeout"`
	// The period of time to consider (in seconds) when evaluating the request rate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#period Ruleset#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The number of requests over the period of time that will trigger the Rate Limiting rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#requests_per_period Ruleset#requests_per_period}
	RequestsPerPeriod *float64 `field:"optional" json:"requestsPerPeriod" yaml:"requestsPerPeriod"`
	// Whether to include requests to origin within the Rate Limiting count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#requests_to_origin Ruleset#requests_to_origin}
	RequestsToOrigin interface{} `field:"optional" json:"requestsToOrigin" yaml:"requestsToOrigin"`
	// The maximum aggregate score over the period of time that will trigger Rate Limiting rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#score_per_period Ruleset#score_per_period}
	ScorePerPeriod *float64 `field:"optional" json:"scorePerPeriod" yaml:"scorePerPeriod"`
	// Name of HTTP header in the response, set by the origin server, with the score for the current request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#score_response_header_name Ruleset#score_response_header_name}
	ScoreResponseHeaderName *string `field:"optional" json:"scoreResponseHeaderName" yaml:"scoreResponseHeaderName"`
}

