package googlenetworksecurityaddressgroupiambinding


type GoogleNetworkSecurityAddressGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_address_group_iam_binding#expression GoogleNetworkSecurityAddressGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_address_group_iam_binding#title GoogleNetworkSecurityAddressGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_address_group_iam_binding#description GoogleNetworkSecurityAddressGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

