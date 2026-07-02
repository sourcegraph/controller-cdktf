package ruleset

type RulesetRulesActionParametersUri struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#origin Ruleset#origin}.
	Origin any `field:"optional" json:"origin" yaml:"origin"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#path Ruleset#path}
	Path any `field:"optional" json:"path" yaml:"path"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#query Ruleset#query}
	Query any `field:"optional" json:"query" yaml:"query"`
}
