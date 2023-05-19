package teamsaccount


type TeamsAccountLoggingSettingsByRuleTypeL4 struct {
	// Whether to log all activity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#log_all TeamsAccount#log_all}
	LogAll interface{} `field:"required" json:"logAll" yaml:"logAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#log_blocks TeamsAccount#log_blocks}.
	LogBlocks interface{} `field:"required" json:"logBlocks" yaml:"logBlocks"`
}

