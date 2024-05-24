package ruleset


type RulesetRulesLogging struct {
	// Override the default logging behavior when a rule is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

