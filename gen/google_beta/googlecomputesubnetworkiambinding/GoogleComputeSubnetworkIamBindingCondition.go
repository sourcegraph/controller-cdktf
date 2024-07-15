package googlecomputesubnetworkiambinding


type GoogleComputeSubnetworkIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_subnetwork_iam_binding#expression GoogleComputeSubnetworkIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_subnetwork_iam_binding#title GoogleComputeSubnetworkIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_subnetwork_iam_binding#description GoogleComputeSubnetworkIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

