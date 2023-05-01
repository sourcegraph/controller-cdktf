package ruleset


type RulesetRulesExposedCredentialCheck struct {
	// Firewall Rules expression language based on Wireshark display filters for where to check for the "password" value.
	//
	// Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#password_expression Ruleset#password_expression}
	PasswordExpression *string `field:"optional" json:"passwordExpression" yaml:"passwordExpression"`
	// Firewall Rules expression language based on Wireshark display filters for where to check for the "username" value.
	//
	// Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#username_expression Ruleset#username_expression}
	UsernameExpression *string `field:"optional" json:"usernameExpression" yaml:"usernameExpression"`
}

