package ruleset


type RulesetRulesActionParametersBrowserTtl struct {
	// Default browser TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
	// Mode of the browser TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

