package teamsaccount


type TeamsAccountLogging struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#redact_pii TeamsAccount#redact_pii}.
	RedactPii interface{} `field:"required" json:"redactPii" yaml:"redactPii"`
	// settings_by_rule_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#settings_by_rule_type TeamsAccount#settings_by_rule_type}
	SettingsByRuleType *TeamsAccountLoggingSettingsByRuleType `field:"required" json:"settingsByRuleType" yaml:"settingsByRuleType"`
}

