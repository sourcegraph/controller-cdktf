package googlebigquerydatasetiambinding


type GoogleBigqueryDatasetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset_iam_binding#expression GoogleBigqueryDatasetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset_iam_binding#title GoogleBigqueryDatasetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset_iam_binding#description GoogleBigqueryDatasetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

