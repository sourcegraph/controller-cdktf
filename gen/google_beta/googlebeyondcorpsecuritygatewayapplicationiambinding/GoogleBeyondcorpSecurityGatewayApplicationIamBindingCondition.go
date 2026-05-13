package googlebeyondcorpsecuritygatewayapplicationiambinding


type GoogleBeyondcorpSecurityGatewayApplicationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application_iam_binding#expression GoogleBeyondcorpSecurityGatewayApplicationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application_iam_binding#title GoogleBeyondcorpSecurityGatewayApplicationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application_iam_binding#description GoogleBeyondcorpSecurityGatewayApplicationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

