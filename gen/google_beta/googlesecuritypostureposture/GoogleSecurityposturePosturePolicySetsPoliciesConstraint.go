package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraint struct {
	// org_policy_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#org_policy_constraint GoogleSecurityposturePosture#org_policy_constraint}
	OrgPolicyConstraint *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint `field:"optional" json:"orgPolicyConstraint" yaml:"orgPolicyConstraint"`
	// org_policy_constraint_custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#org_policy_constraint_custom GoogleSecurityposturePosture#org_policy_constraint_custom}
	OrgPolicyConstraintCustom *GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustom `field:"optional" json:"orgPolicyConstraintCustom" yaml:"orgPolicyConstraintCustom"`
	// security_health_analytics_custom_module block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#security_health_analytics_custom_module GoogleSecurityposturePosture#security_health_analytics_custom_module}
	SecurityHealthAnalyticsCustomModule *GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModule `field:"optional" json:"securityHealthAnalyticsCustomModule" yaml:"securityHealthAnalyticsCustomModule"`
	// security_health_analytics_module block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#security_health_analytics_module GoogleSecurityposturePosture#security_health_analytics_module}
	SecurityHealthAnalyticsModule *GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsModule `field:"optional" json:"securityHealthAnalyticsModule" yaml:"securityHealthAnalyticsModule"`
}

