package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHost struct {
	// Resolve hostname to IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#resolved Ruleset#resolved}
	Resolved interface{} `field:"optional" json:"resolved" yaml:"resolved"`
}

