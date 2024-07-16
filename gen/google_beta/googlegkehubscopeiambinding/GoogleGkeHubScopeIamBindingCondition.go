package googlegkehubscopeiambinding


type GoogleGkeHubScopeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_scope_iam_binding#expression GoogleGkeHubScopeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_scope_iam_binding#title GoogleGkeHubScopeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_scope_iam_binding#description GoogleGkeHubScopeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

