package googlebigtabletableiammember


type GoogleBigtableTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table_iam_member#expression GoogleBigtableTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table_iam_member#title GoogleBigtableTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table_iam_member#description GoogleBigtableTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

