package googlecomputesubnetworkiammember


type GoogleComputeSubnetworkIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork_iam_member#expression GoogleComputeSubnetworkIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork_iam_member#title GoogleComputeSubnetworkIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork_iam_member#description GoogleComputeSubnetworkIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

