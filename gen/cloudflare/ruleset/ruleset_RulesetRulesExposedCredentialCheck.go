package ruleset


type RulesetRulesExposedCredentialCheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#password_expression Ruleset#password_expression}.
	PasswordExpression *string `field:"optional" json:"passwordExpression" yaml:"passwordExpression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#username_expression Ruleset#username_expression}.
	UsernameExpression *string `field:"optional" json:"usernameExpression" yaml:"usernameExpression"`
}

