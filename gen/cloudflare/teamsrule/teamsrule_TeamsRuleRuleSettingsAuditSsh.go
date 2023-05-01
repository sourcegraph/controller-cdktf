package teamsrule


type TeamsRuleRuleSettingsAuditSsh struct {
	// Log all SSH commands.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#command_logging TeamsRule#command_logging}
	CommandLogging interface{} `field:"required" json:"commandLogging" yaml:"commandLogging"`
}

