package googlebigquerytableiammember


type GoogleBigqueryTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table_iam_member#expression GoogleBigqueryTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table_iam_member#title GoogleBigqueryTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table_iam_member#description GoogleBigqueryTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

