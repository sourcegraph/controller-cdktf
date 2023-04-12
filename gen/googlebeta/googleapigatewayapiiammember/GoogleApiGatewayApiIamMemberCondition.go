package googleapigatewayapiiammember


type GoogleApiGatewayApiIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_iam_member#expression GoogleApiGatewayApiIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_iam_member#title GoogleApiGatewayApiIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_iam_member#description GoogleApiGatewayApiIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

