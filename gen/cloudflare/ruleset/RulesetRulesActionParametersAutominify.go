package ruleset

type RulesetRulesActionParametersAutominify struct {
	// CSS minification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#css Ruleset#css}
	Css any `field:"optional" json:"css" yaml:"css"`
	// HTML minification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#html Ruleset#html}
	Html any `field:"optional" json:"html" yaml:"html"`
	// JS minification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#js Ruleset#js}
	Js any `field:"optional" json:"js" yaml:"js"`
}
