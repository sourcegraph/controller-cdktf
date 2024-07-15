package googlegkehubfeatureiambinding


type GoogleGkeHubFeatureIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_iam_binding#expression GoogleGkeHubFeatureIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_iam_binding#title GoogleGkeHubFeatureIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_iam_binding#description GoogleGkeHubFeatureIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

