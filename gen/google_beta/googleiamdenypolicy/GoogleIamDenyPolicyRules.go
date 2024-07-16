package googleiamdenypolicy


type GoogleIamDenyPolicyRules struct {
	// deny_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_deny_policy#deny_rule GoogleIamDenyPolicy#deny_rule}
	DenyRule *GoogleIamDenyPolicyRulesDenyRule `field:"optional" json:"denyRule" yaml:"denyRule"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_deny_policy#description GoogleIamDenyPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

