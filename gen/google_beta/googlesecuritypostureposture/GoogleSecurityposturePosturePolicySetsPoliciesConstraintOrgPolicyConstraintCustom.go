package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustom struct {
	// policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#policy_rules GoogleSecurityposturePosture#policy_rules}
	PolicyRules interface{} `field:"required" json:"policyRules" yaml:"policyRules"`
	// custom_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#custom_constraint GoogleSecurityposturePosture#custom_constraint}
	CustomConstraint *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomCustomConstraint `field:"optional" json:"customConstraint" yaml:"customConstraint"`
}

