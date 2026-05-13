package googlebeyondcorpsecuritygatewayapplicationiammember


type GoogleBeyondcorpSecurityGatewayApplicationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application_iam_member#expression GoogleBeyondcorpSecurityGatewayApplicationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application_iam_member#title GoogleBeyondcorpSecurityGatewayApplicationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application_iam_member#description GoogleBeyondcorpSecurityGatewayApplicationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

