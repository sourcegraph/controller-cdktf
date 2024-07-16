package googleorgpolicypolicy


type GoogleOrgPolicyPolicySpecRulesCondition struct {
	// Optional.
	//
	// Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_org_policy_policy#description GoogleOrgPolicyPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_org_policy_policy#expression GoogleOrgPolicyPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Optional.
	//
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_org_policy_policy#location GoogleOrgPolicyPolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Optional.
	//
	// Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_org_policy_policy#title GoogleOrgPolicyPolicy#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

