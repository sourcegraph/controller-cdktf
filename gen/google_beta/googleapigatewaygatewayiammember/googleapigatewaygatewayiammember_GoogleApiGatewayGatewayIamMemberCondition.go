package googleapigatewaygatewayiammember


type GoogleApiGatewayGatewayIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_gateway_iam_member#expression GoogleApiGatewayGatewayIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_gateway_iam_member#title GoogleApiGatewayGatewayIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_gateway_iam_member#description GoogleApiGatewayGatewayIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

