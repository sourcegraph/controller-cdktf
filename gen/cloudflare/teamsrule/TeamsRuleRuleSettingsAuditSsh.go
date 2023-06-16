package teamsrule


type TeamsRuleRuleSettingsAuditSsh struct {
	// Log all SSH commands.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#command_logging TeamsRule#command_logging}
	CommandLogging interface{} `field:"required" json:"commandLogging" yaml:"commandLogging"`
}

