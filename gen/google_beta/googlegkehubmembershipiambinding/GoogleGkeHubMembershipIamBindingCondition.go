package googlegkehubmembershipiambinding


type GoogleGkeHubMembershipIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_membership_iam_binding#expression GoogleGkeHubMembershipIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_membership_iam_binding#title GoogleGkeHubMembershipIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_membership_iam_binding#description GoogleGkeHubMembershipIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

