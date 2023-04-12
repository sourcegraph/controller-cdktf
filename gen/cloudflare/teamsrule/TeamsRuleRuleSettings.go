package teamsrule


type TeamsRuleRuleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#add_headers TeamsRule#add_headers}.
	AddHeaders *map[string]*string `field:"optional" json:"addHeaders" yaml:"addHeaders"`
	// biso_admin_controls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#biso_admin_controls TeamsRule#biso_admin_controls}
	BisoAdminControls *TeamsRuleRuleSettingsBisoAdminControls `field:"optional" json:"bisoAdminControls" yaml:"bisoAdminControls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#block_page_enabled TeamsRule#block_page_enabled}.
	BlockPageEnabled interface{} `field:"optional" json:"blockPageEnabled" yaml:"blockPageEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#block_page_reason TeamsRule#block_page_reason}.
	BlockPageReason *string `field:"optional" json:"blockPageReason" yaml:"blockPageReason"`
	// check_session block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#check_session TeamsRule#check_session}
	CheckSession *TeamsRuleRuleSettingsCheckSession `field:"optional" json:"checkSession" yaml:"checkSession"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#insecure_disable_dnssec_validation TeamsRule#insecure_disable_dnssec_validation}.
	InsecureDisableDnssecValidation interface{} `field:"optional" json:"insecureDisableDnssecValidation" yaml:"insecureDisableDnssecValidation"`
	// l4override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#l4override TeamsRule#l4override}
	L4Override *TeamsRuleRuleSettingsL4Override `field:"optional" json:"l4Override" yaml:"l4Override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#override_host TeamsRule#override_host}.
	OverrideHost *string `field:"optional" json:"overrideHost" yaml:"overrideHost"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#override_ips TeamsRule#override_ips}.
	OverrideIps *[]*string `field:"optional" json:"overrideIps" yaml:"overrideIps"`
}

