package googlespannerdatabaseiambinding


type GoogleSpannerDatabaseIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_database_iam_binding#expression GoogleSpannerDatabaseIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_database_iam_binding#title GoogleSpannerDatabaseIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_database_iam_binding#description GoogleSpannerDatabaseIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

