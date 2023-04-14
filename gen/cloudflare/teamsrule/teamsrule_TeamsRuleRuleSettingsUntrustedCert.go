package teamsrule


type TeamsRuleRuleSettingsUntrustedCert struct {
	// Action to be taken when the SSL certificate of upstream is invalid. Available values: `pass_through`, `block`, `error`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#action TeamsRule#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
}

