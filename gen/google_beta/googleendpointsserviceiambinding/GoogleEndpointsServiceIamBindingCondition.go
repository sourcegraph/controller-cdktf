package googleendpointsserviceiambinding


type GoogleEndpointsServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_endpoints_service_iam_binding#expression GoogleEndpointsServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_endpoints_service_iam_binding#title GoogleEndpointsServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_endpoints_service_iam_binding#description GoogleEndpointsServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

