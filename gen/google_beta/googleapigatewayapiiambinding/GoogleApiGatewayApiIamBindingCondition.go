package googleapigatewayapiiambinding


type GoogleApiGatewayApiIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_iam_binding#expression GoogleApiGatewayApiIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_iam_binding#title GoogleApiGatewayApiIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_iam_binding#description GoogleApiGatewayApiIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

