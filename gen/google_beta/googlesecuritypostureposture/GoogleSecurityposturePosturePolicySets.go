package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySets struct {
	// policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#policies GoogleSecurityposturePosture#policies}
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
	// ID of the policy set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#policy_set_id GoogleSecurityposturePosture#policy_set_id}
	PolicySetId *string `field:"required" json:"policySetId" yaml:"policySetId"`
	// Description of the policy set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#description GoogleSecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

