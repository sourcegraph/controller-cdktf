package ruleset


type RulesetRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#expression Ruleset#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#action Ruleset#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// action_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#action_parameters Ruleset#action_parameters}
	ActionParameters *RulesetRulesActionParameters `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#description Ruleset#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#enabled Ruleset#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// exposed_credential_check block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#exposed_credential_check Ruleset#exposed_credential_check}
	ExposedCredentialCheck *RulesetRulesExposedCredentialCheck `field:"optional" json:"exposedCredentialCheck" yaml:"exposedCredentialCheck"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#logging Ruleset#logging}
	Logging *RulesetRulesLogging `field:"optional" json:"logging" yaml:"logging"`
	// ratelimit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#ratelimit Ruleset#ratelimit}
	Ratelimit *RulesetRulesRatelimit `field:"optional" json:"ratelimit" yaml:"ratelimit"`
}

