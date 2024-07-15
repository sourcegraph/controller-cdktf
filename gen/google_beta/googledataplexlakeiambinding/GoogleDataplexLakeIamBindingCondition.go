package googledataplexlakeiambinding


type GoogleDataplexLakeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_lake_iam_binding#expression GoogleDataplexLakeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_lake_iam_binding#title GoogleDataplexLakeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_lake_iam_binding#description GoogleDataplexLakeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

