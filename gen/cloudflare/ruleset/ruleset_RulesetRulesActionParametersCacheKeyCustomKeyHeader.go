package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHeader struct {
	// List of headers to check for presence in the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#check_presence Ruleset#check_presence}
	CheckPresence *[]*string `field:"optional" json:"checkPresence" yaml:"checkPresence"`
	// Exclude the origin header from the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#exclude_origin Ruleset#exclude_origin}
	ExcludeOrigin interface{} `field:"optional" json:"excludeOrigin" yaml:"excludeOrigin"`
	// List of headers to include in the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#include Ruleset#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

