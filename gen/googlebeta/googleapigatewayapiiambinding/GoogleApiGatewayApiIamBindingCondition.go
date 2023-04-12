package googleapigatewayapiiambinding


type GoogleApiGatewayApiIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_iam_binding#expression GoogleApiGatewayApiIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_iam_binding#title GoogleApiGatewayApiIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_iam_binding#description GoogleApiGatewayApiIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

