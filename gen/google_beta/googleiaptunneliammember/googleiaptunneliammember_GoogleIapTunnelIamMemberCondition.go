package googleiaptunneliammember


type GoogleIapTunnelIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_iam_member#expression GoogleIapTunnelIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_iam_member#title GoogleIapTunnelIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_iam_member#description GoogleIapTunnelIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

