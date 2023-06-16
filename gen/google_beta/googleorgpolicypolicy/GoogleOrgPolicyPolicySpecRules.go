package googleorgpolicypolicy


type GoogleOrgPolicyPolicySpecRules struct {
	// Setting this to true means that all values are allowed.
	//
	// This field can be set only in Policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_org_policy_policy#allow_all GoogleOrgPolicyPolicy#allow_all}
	AllowAll *string `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_org_policy_policy#condition GoogleOrgPolicyPolicy#condition}
	Condition *GoogleOrgPolicyPolicySpecRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to true means that all values are denied.
	//
	// This field can be set only in Policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_org_policy_policy#deny_all GoogleOrgPolicyPolicy#deny_all}
	DenyAll *string `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If `true`, then the `Policy` is enforced.
	//
	// If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_org_policy_policy#enforce GoogleOrgPolicyPolicy#enforce}
	Enforce *string `field:"optional" json:"enforce" yaml:"enforce"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_org_policy_policy#values GoogleOrgPolicyPolicy#values}
	Values *GoogleOrgPolicyPolicySpecRulesValues `field:"optional" json:"values" yaml:"values"`
}

