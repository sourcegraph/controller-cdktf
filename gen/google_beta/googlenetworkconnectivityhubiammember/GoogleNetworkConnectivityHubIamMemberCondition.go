package googlenetworkconnectivityhubiammember


type GoogleNetworkConnectivityHubIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_hub_iam_member#expression GoogleNetworkConnectivityHubIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_hub_iam_member#title GoogleNetworkConnectivityHubIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_hub_iam_member#description GoogleNetworkConnectivityHubIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

