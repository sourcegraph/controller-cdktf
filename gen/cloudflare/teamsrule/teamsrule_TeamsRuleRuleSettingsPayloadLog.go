package teamsrule


type TeamsRuleRuleSettingsPayloadLog struct {
	// Enable or disable DLP Payload Logging for this rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#enabled TeamsRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

