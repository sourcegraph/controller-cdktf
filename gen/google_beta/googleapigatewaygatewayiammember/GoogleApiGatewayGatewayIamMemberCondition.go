package googleapigatewaygatewayiammember


type GoogleApiGatewayGatewayIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_gateway_iam_member#expression GoogleApiGatewayGatewayIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_gateway_iam_member#title GoogleApiGatewayGatewayIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_gateway_iam_member#description GoogleApiGatewayGatewayIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

