package googledataplexassetiambinding


type GoogleDataplexAssetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_asset_iam_binding#expression GoogleDataplexAssetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_asset_iam_binding#title GoogleDataplexAssetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_asset_iam_binding#description GoogleDataplexAssetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

