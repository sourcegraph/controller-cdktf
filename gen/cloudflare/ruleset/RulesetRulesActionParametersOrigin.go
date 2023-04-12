package ruleset


type RulesetRulesActionParametersOrigin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#host Ruleset#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#port Ruleset#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

