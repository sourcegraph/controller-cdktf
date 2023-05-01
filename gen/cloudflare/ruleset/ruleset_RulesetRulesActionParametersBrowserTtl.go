package ruleset


type RulesetRulesActionParametersBrowserTtl struct {
	// Default browser TTL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
	// Mode of the browser TTL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#mode Ruleset#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

