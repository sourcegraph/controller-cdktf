package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesComplianceStandards struct {
	// Mapping of security controls for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#control GoogleSecurityposturePosture#control}
	Control *string `field:"optional" json:"control" yaml:"control"`
	// Mapping of compliance standards for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#standard GoogleSecurityposturePosture#standard}
	Standard *string `field:"optional" json:"standard" yaml:"standard"`
}

