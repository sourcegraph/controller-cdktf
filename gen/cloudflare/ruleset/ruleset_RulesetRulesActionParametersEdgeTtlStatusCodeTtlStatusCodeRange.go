package ruleset


type RulesetRulesActionParametersEdgeTtlStatusCodeTtlStatusCodeRange struct {
	// From status code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#from Ruleset#from}
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// To status code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#to Ruleset#to}
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

