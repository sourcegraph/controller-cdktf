package teamsaccount


type TeamsAccountLoggingSettingsByRuleTypeHttp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#log_all TeamsAccount#log_all}.
	LogAll interface{} `field:"required" json:"logAll" yaml:"logAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#log_blocks TeamsAccount#log_blocks}.
	LogBlocks interface{} `field:"required" json:"logBlocks" yaml:"logBlocks"`
}

