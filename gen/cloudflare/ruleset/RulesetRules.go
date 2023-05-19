package ruleset


type RulesetRules struct {
	// Criteria for an HTTP request to trigger the ruleset rule action.
	//
	// Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Action to perform in the ruleset rule.
	//
	// Available values: `allow`, `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// action_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#action_parameters Ruleset#action_parameters}
	ActionParameters interface{} `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// Brief summary of the ruleset rule and its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#description Ruleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// exposed_credential_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#exposed_credential_check Ruleset#exposed_credential_check}
	ExposedCredentialCheck interface{} `field:"optional" json:"exposedCredentialCheck" yaml:"exposedCredentialCheck"`
	// Unique rule identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#id Ruleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The most recent update to this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#last_updated Ruleset#last_updated}
	LastUpdated *string `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#logging Ruleset#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// ratelimit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#ratelimit Ruleset#ratelimit}
	Ratelimit interface{} `field:"optional" json:"ratelimit" yaml:"ratelimit"`
	// Rule reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#ref Ruleset#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Version of the ruleset to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#version Ruleset#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

