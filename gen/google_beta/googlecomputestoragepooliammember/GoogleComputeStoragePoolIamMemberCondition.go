package googlecomputestoragepooliammember


type GoogleComputeStoragePoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool_iam_member#expression GoogleComputeStoragePoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool_iam_member#title GoogleComputeStoragePoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool_iam_member#description GoogleComputeStoragePoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

