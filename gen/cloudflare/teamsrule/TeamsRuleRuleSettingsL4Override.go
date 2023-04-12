package teamsrule


type TeamsRuleRuleSettingsL4Override struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#ip TeamsRule#ip}.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_rule#port TeamsRule#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

