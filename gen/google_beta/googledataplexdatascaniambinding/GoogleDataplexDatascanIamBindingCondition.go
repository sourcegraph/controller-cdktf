package googledataplexdatascaniambinding


type GoogleDataplexDatascanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan_iam_binding#expression GoogleDataplexDatascanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan_iam_binding#title GoogleDataplexDatascanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan_iam_binding#description GoogleDataplexDatascanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

