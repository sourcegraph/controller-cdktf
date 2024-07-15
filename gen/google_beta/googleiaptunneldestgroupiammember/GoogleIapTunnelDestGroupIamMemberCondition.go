package googleiaptunneldestgroupiammember


type GoogleIapTunnelDestGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_member#expression GoogleIapTunnelDestGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_member#title GoogleIapTunnelDestGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_member#description GoogleIapTunnelDestGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

