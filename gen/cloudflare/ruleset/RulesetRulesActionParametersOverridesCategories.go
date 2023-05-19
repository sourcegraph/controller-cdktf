package ruleset


type RulesetRulesActionParametersOverridesCategories struct {
	// Action to perform in the tag-level override.
	//
	// Available values: `allow`, `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Tag name to apply the ruleset rule override to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#category Ruleset#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Defines if the current tag-level override enables or disables the ruleset rules with the specified tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

