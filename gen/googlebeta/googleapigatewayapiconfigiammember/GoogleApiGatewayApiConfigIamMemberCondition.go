package googleapigatewayapiconfigiammember


type GoogleApiGatewayApiConfigIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config_iam_member#expression GoogleApiGatewayApiConfigIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config_iam_member#title GoogleApiGatewayApiConfigIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config_iam_member#description GoogleApiGatewayApiConfigIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

