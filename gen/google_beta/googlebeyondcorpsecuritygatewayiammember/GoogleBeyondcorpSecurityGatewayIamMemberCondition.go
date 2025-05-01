package googlebeyondcorpsecuritygatewayiammember


type GoogleBeyondcorpSecurityGatewayIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_beyondcorp_security_gateway_iam_member#expression GoogleBeyondcorpSecurityGatewayIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_beyondcorp_security_gateway_iam_member#title GoogleBeyondcorpSecurityGatewayIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_beyondcorp_security_gateway_iam_member#description GoogleBeyondcorpSecurityGatewayIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

