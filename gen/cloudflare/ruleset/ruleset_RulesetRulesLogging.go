package ruleset


type RulesetRulesLogging struct {
	// Override the default logging behavior when a rule is matched.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

