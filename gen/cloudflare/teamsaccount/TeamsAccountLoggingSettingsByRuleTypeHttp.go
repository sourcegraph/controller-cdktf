package teamsaccount

type TeamsAccountLoggingSettingsByRuleTypeHttp struct {
	// Whether to log all activity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#log_all TeamsAccount#log_all}
	LogAll any `field:"required" json:"logAll" yaml:"logAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#log_blocks TeamsAccount#log_blocks}.
	LogBlocks any `field:"required" json:"logBlocks" yaml:"logBlocks"`
}
