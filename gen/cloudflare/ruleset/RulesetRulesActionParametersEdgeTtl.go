package ruleset


type RulesetRulesActionParametersEdgeTtl struct {
	// Default edge TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
	// Mode of the edge TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// status_code_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#status_code_ttl Ruleset#status_code_ttl}
	StatusCodeTtl interface{} `field:"optional" json:"statusCodeTtl" yaml:"statusCodeTtl"`
}

