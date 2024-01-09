package ruleset


type RulesetRulesActionParametersEdgeTtl struct {
	// Mode of the edge TTL. Available values: `override_origin`, `respect_origin`, `bypass_by_default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Default edge TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
	// status_code_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#status_code_ttl Ruleset#status_code_ttl}
	StatusCodeTtl interface{} `field:"optional" json:"statusCodeTtl" yaml:"statusCodeTtl"`
}

