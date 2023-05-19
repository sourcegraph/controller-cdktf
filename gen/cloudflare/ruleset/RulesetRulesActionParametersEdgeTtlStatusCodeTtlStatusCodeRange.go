package ruleset


type RulesetRulesActionParametersEdgeTtlStatusCodeTtlStatusCodeRange struct {
	// From status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#from Ruleset#from}
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// To status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#to Ruleset#to}
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

