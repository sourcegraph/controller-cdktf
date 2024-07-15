package googleapigatewayapiconfigiammember


type GoogleApiGatewayApiConfigIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config_iam_member#expression GoogleApiGatewayApiConfigIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config_iam_member#title GoogleApiGatewayApiConfigIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config_iam_member#description GoogleApiGatewayApiConfigIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

