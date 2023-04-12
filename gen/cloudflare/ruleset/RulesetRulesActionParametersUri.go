package ruleset


type RulesetRulesActionParametersUri struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#origin Ruleset#origin}.
	Origin interface{} `field:"optional" json:"origin" yaml:"origin"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#path Ruleset#path}
	Path *RulesetRulesActionParametersUriPath `field:"optional" json:"path" yaml:"path"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#query Ruleset#query}
	Query *RulesetRulesActionParametersUriQuery `field:"optional" json:"query" yaml:"query"`
}

