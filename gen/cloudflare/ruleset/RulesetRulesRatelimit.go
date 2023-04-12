package ruleset


type RulesetRulesRatelimit struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#characteristics Ruleset#characteristics}.
	Characteristics *[]*string `field:"optional" json:"characteristics" yaml:"characteristics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#counting_expression Ruleset#counting_expression}.
	CountingExpression *string `field:"optional" json:"countingExpression" yaml:"countingExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#mitigation_timeout Ruleset#mitigation_timeout}.
	MitigationTimeout *float64 `field:"optional" json:"mitigationTimeout" yaml:"mitigationTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#period Ruleset#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#requests_per_period Ruleset#requests_per_period}.
	RequestsPerPeriod *float64 `field:"optional" json:"requestsPerPeriod" yaml:"requestsPerPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#requests_to_origin Ruleset#requests_to_origin}.
	RequestsToOrigin interface{} `field:"optional" json:"requestsToOrigin" yaml:"requestsToOrigin"`
}

