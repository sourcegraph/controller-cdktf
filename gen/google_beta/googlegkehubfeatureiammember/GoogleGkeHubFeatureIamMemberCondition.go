package googlegkehubfeatureiammember


type GoogleGkeHubFeatureIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_iam_member#expression GoogleGkeHubFeatureIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_iam_member#title GoogleGkeHubFeatureIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_iam_member#description GoogleGkeHubFeatureIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

