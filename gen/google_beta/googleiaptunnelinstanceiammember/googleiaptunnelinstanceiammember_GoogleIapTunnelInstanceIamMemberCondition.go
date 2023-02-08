package googleiaptunnelinstanceiammember


type GoogleIapTunnelInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_instance_iam_member#expression GoogleIapTunnelInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_instance_iam_member#title GoogleIapTunnelInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_instance_iam_member#description GoogleIapTunnelInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

