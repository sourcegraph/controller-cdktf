package googleendpointsserviceconsumersiammember


type GoogleEndpointsServiceConsumersIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_consumers_iam_member#expression GoogleEndpointsServiceConsumersIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_consumers_iam_member#title GoogleEndpointsServiceConsumersIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_consumers_iam_member#description GoogleEndpointsServiceConsumersIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

