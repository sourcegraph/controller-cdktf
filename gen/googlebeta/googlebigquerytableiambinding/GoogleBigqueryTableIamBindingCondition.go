package googlebigquerytableiambinding


type GoogleBigqueryTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table_iam_binding#expression GoogleBigqueryTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table_iam_binding#title GoogleBigqueryTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table_iam_binding#description GoogleBigqueryTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

