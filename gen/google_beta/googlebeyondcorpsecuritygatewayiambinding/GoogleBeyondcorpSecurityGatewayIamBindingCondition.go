package googlebeyondcorpsecuritygatewayiambinding


type GoogleBeyondcorpSecurityGatewayIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_security_gateway_iam_binding#expression GoogleBeyondcorpSecurityGatewayIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_security_gateway_iam_binding#title GoogleBeyondcorpSecurityGatewayIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_security_gateway_iam_binding#description GoogleBeyondcorpSecurityGatewayIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

