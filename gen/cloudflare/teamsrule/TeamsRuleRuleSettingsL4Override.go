package teamsrule


type TeamsRuleRuleSettingsL4Override struct {
	// Override IP to forward traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#ip TeamsRule#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Override Port to forward traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#port TeamsRule#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

