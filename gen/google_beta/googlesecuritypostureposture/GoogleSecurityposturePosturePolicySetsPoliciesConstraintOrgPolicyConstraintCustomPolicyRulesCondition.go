package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#expression GoogleSecurityposturePosture#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Description of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#description GoogleSecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#location GoogleSecurityposturePosture#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Title for the expression, i.e. a short string describing its purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#title GoogleSecurityposturePosture#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

