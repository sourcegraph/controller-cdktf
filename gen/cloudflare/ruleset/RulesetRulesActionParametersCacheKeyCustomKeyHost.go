package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHost struct {
	// Resolve hostname to IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#resolved Ruleset#resolved}
	Resolved interface{} `field:"optional" json:"resolved" yaml:"resolved"`
}

