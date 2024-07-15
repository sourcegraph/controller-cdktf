package googleiamaccessboundarypolicy


type GoogleIamAccessBoundaryPolicyRules struct {
	// access_boundary_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_access_boundary_policy#access_boundary_rule GoogleIamAccessBoundaryPolicy#access_boundary_rule}
	AccessBoundaryRule *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule `field:"optional" json:"accessBoundaryRule" yaml:"accessBoundaryRule"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_access_boundary_policy#description GoogleIamAccessBoundaryPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

