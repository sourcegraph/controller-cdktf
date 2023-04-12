package googleendpointsserviceconsumersiambinding


type GoogleEndpointsServiceConsumersIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_consumers_iam_binding#expression GoogleEndpointsServiceConsumersIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_consumers_iam_binding#title GoogleEndpointsServiceConsumersIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_endpoints_service_consumers_iam_binding#description GoogleEndpointsServiceConsumersIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

