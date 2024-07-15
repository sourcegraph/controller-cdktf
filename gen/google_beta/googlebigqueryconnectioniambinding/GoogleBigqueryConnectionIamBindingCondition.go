package googlebigqueryconnectioniambinding


type GoogleBigqueryConnectionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection_iam_binding#expression GoogleBigqueryConnectionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection_iam_binding#title GoogleBigqueryConnectionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection_iam_binding#description GoogleBigqueryConnectionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

