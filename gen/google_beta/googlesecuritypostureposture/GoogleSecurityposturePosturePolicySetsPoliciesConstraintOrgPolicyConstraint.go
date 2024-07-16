package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint struct {
	// Organization policy canned constraint Id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#canned_constraint_id GoogleSecurityposturePosture#canned_constraint_id}
	CannedConstraintId *string `field:"required" json:"cannedConstraintId" yaml:"cannedConstraintId"`
	// policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#policy_rules GoogleSecurityposturePosture#policy_rules}
	PolicyRules interface{} `field:"required" json:"policyRules" yaml:"policyRules"`
}

