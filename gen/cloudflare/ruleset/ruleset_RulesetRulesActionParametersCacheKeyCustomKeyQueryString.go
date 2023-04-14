package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyQueryString struct {
	// List of query string parameters to exclude from the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#exclude Ruleset#exclude}
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of query string parameters to include in the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#include Ruleset#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

