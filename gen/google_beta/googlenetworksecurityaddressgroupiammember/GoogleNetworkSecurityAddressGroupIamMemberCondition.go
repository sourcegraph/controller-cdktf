package googlenetworksecurityaddressgroupiammember


type GoogleNetworkSecurityAddressGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_address_group_iam_member#expression GoogleNetworkSecurityAddressGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_address_group_iam_member#title GoogleNetworkSecurityAddressGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_address_group_iam_member#description GoogleNetworkSecurityAddressGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

