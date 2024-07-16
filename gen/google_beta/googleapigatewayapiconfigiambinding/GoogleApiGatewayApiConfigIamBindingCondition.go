package googleapigatewayapiconfigiambinding


type GoogleApiGatewayApiConfigIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config_iam_binding#expression GoogleApiGatewayApiConfigIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config_iam_binding#title GoogleApiGatewayApiConfigIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config_iam_binding#description GoogleApiGatewayApiConfigIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

