package ruleset


type RulesetRulesActionParametersOverrides struct {
	// Action to perform in the rule-level override.
	//
	// Available values: `allow`, `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// categories block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#categories Ruleset#categories}
	Categories interface{} `field:"optional" json:"categories" yaml:"categories"`
	// Defines if the current ruleset-level override enables or disables the ruleset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#rules Ruleset#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Sensitivity level to override for all ruleset rules. Available values: `default`, `medium`, `low`, `eoff`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#sensitivity_level Ruleset#sensitivity_level}
	SensitivityLevel *string `field:"optional" json:"sensitivityLevel" yaml:"sensitivityLevel"`
}

