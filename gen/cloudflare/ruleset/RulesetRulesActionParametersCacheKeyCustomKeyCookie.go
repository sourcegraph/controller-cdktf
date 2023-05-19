package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyCookie struct {
	// List of cookies to check for presence in the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#check_presence Ruleset#check_presence}
	CheckPresence *[]*string `field:"optional" json:"checkPresence" yaml:"checkPresence"`
	// List of cookies to include in the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#include Ruleset#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

