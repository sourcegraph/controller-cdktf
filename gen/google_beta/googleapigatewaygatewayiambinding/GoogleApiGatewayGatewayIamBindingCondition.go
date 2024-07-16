package googleapigatewaygatewayiambinding


type GoogleApiGatewayGatewayIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_gateway_iam_binding#expression GoogleApiGatewayGatewayIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_gateway_iam_binding#title GoogleApiGatewayGatewayIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_gateway_iam_binding#description GoogleApiGatewayGatewayIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

