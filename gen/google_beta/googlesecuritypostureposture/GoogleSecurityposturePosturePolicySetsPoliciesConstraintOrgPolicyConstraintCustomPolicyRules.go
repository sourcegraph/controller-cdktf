package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRules struct {
	// Setting this to true means that all values are allowed.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#allow_all GoogleSecurityposturePosture#allow_all}
	AllowAll interface{} `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#condition GoogleSecurityposturePosture#condition}
	Condition *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to true means that all values are denied.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#deny_all GoogleSecurityposturePosture#deny_all}
	DenyAll interface{} `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If 'true', then the policy is enforced.
	//
	// If 'false', then any configuration is acceptable.
	// This field can be set only in policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#enforce GoogleSecurityposturePosture#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#values GoogleSecurityposturePosture#values}
	Values *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues `field:"optional" json:"values" yaml:"values"`
}

