package teamsaccount


type TeamsAccountLoggingSettingsByRuleType struct {
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#dns TeamsAccount#dns}
	Dns *TeamsAccountLoggingSettingsByRuleTypeDns `field:"required" json:"dns" yaml:"dns"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#http TeamsAccount#http}
	Http *TeamsAccountLoggingSettingsByRuleTypeHttp `field:"required" json:"http" yaml:"http"`
	// l4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#l4 TeamsAccount#l4}
	L4 *TeamsAccountLoggingSettingsByRuleTypeL4 `field:"required" json:"l4" yaml:"l4"`
}

