package ruleset


type RulesetRules struct {
	// Criteria for an HTTP request to trigger the ruleset rule action.
	//
	// Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Action to perform in the ruleset rule.
	//
	// Available values: `block`, `challenge`, `compress_response`, `ddos_dynamic`, `ddos_mitigation`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `serve_error`, `set_cache_settings`, `set_config`, `skip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// action_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#action_parameters Ruleset#action_parameters}
	ActionParameters interface{} `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// Brief summary of the ruleset rule and its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#description Ruleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// exposed_credential_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#exposed_credential_check Ruleset#exposed_credential_check}
	ExposedCredentialCheck interface{} `field:"optional" json:"exposedCredentialCheck" yaml:"exposedCredentialCheck"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#logging Ruleset#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// ratelimit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#ratelimit Ruleset#ratelimit}
	Ratelimit interface{} `field:"optional" json:"ratelimit" yaml:"ratelimit"`
	// Rule reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/ruleset#ref Ruleset#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

