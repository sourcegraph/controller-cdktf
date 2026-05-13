package networkconnectivityhubiambinding


type NetworkConnectivityHubIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_hub_iam_binding#expression NetworkConnectivityHubIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_hub_iam_binding#title NetworkConnectivityHubIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_hub_iam_binding#description NetworkConnectivityHubIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

