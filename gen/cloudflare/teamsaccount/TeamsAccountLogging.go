package teamsaccount


type TeamsAccountLogging struct {
	// Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#redact_pii TeamsAccount#redact_pii}
	RedactPii interface{} `field:"required" json:"redactPii" yaml:"redactPii"`
	// settings_by_rule_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#settings_by_rule_type TeamsAccount#settings_by_rule_type}
	SettingsByRuleType *TeamsAccountLoggingSettingsByRuleType `field:"required" json:"settingsByRuleType" yaml:"settingsByRuleType"`
}

