package googleiaptunneliambinding


type GoogleIapTunnelIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_iam_binding#expression GoogleIapTunnelIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_iam_binding#title GoogleIapTunnelIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_iam_binding#description GoogleIapTunnelIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

