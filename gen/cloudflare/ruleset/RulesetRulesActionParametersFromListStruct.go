package ruleset


type RulesetRulesActionParametersFromListStruct struct {
	// Expression to use for the list lookup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#key Ruleset#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

