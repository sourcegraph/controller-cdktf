package googlebigtableinstanceiammember


type GoogleBigtableInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_instance_iam_member#expression GoogleBigtableInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_instance_iam_member#title GoogleBigtableInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_instance_iam_member#description GoogleBigtableInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

