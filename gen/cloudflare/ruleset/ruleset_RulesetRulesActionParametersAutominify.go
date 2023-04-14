package ruleset


type RulesetRulesActionParametersAutominify struct {
	// CSS minification.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#css Ruleset#css}
	Css interface{} `field:"optional" json:"css" yaml:"css"`
	// HTML minification.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#html Ruleset#html}
	Html interface{} `field:"optional" json:"html" yaml:"html"`
	// JS minification.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#js Ruleset#js}
	Js interface{} `field:"optional" json:"js" yaml:"js"`
}

