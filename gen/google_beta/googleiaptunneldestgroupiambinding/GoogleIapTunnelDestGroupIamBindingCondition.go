package googleiaptunneldestgroupiambinding


type GoogleIapTunnelDestGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#expression GoogleIapTunnelDestGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#title GoogleIapTunnelDestGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#description GoogleIapTunnelDestGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

