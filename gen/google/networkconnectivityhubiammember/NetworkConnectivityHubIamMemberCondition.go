package networkconnectivityhubiammember


type NetworkConnectivityHubIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_hub_iam_member#expression NetworkConnectivityHubIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_hub_iam_member#title NetworkConnectivityHubIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_hub_iam_member#description NetworkConnectivityHubIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

