package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPolicies struct {
	// constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#constraint GoogleSecurityposturePosture#constraint}
	Constraint *GoogleSecurityposturePosturePolicySetsPoliciesConstraint `field:"required" json:"constraint" yaml:"constraint"`
	// ID of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#policy_id GoogleSecurityposturePosture#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// compliance_standards block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#compliance_standards GoogleSecurityposturePosture#compliance_standards}
	ComplianceStandards interface{} `field:"optional" json:"complianceStandards" yaml:"complianceStandards"`
	// Description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#description GoogleSecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

