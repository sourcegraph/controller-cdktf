package teamsrule


type TeamsRuleRuleSettingsCheckSession struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#duration TeamsRule#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#enforce TeamsRule#enforce}.
	Enforce interface{} `field:"required" json:"enforce" yaml:"enforce"`
}

