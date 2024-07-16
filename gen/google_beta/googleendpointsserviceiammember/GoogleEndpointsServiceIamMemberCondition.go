package googleendpointsserviceiammember


type GoogleEndpointsServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_endpoints_service_iam_member#expression GoogleEndpointsServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_endpoints_service_iam_member#title GoogleEndpointsServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_endpoints_service_iam_member#description GoogleEndpointsServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

