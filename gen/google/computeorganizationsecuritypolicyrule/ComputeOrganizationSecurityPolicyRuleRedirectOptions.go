package computeorganizationsecuritypolicyrule


type ComputeOrganizationSecurityPolicyRuleRedirectOptions struct {
	// Type of the redirect action. For organization security policies, only EXTERNAL_302 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_rule#type ComputeOrganizationSecurityPolicyRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Target for the redirect action. This is required if the type is EXTERNAL_302.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_rule#target ComputeOrganizationSecurityPolicyRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

