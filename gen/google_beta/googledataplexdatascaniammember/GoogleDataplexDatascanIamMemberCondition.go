package googledataplexdatascaniammember


type GoogleDataplexDatascanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan_iam_member#expression GoogleDataplexDatascanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan_iam_member#title GoogleDataplexDatascanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan_iam_member#description GoogleDataplexDatascanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

