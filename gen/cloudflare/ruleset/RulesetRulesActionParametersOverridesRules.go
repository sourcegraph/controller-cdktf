package ruleset


type RulesetRulesActionParametersOverridesRules struct {
	// Action to perform in the rule-level override.
	//
	// Available values: `allow`, `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Defines if the current rule-level override enables or disables the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Rule ID to apply the override to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#id Ruleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Anomaly score threshold to apply in the ruleset rule override. Only applicable to modsecurity-based rulesets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#score_threshold Ruleset#score_threshold}
	ScoreThreshold *float64 `field:"optional" json:"scoreThreshold" yaml:"scoreThreshold"`
	// Sensitivity level for a ruleset rule override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#sensitivity_level Ruleset#sensitivity_level}
	SensitivityLevel *string `field:"optional" json:"sensitivityLevel" yaml:"sensitivityLevel"`
}

