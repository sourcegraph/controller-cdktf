package teamsrule


type TeamsRuleRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#duration TeamsRule#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Enable session enforcement for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#enforce TeamsRule#enforce}
	Enforce interface{} `field:"required" json:"enforce" yaml:"enforce"`
}

