package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyUser struct {
	// Add device type to the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#device_type Ruleset#device_type}
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Add geo data to the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#geo Ruleset#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// Add language data to the custom key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#lang Ruleset#lang}
	Lang interface{} `field:"optional" json:"lang" yaml:"lang"`
}

