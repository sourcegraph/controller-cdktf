package googleiaptunnelinstanceiambinding


type GoogleIapTunnelInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_instance_iam_binding#expression GoogleIapTunnelInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_instance_iam_binding#title GoogleIapTunnelInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_instance_iam_binding#description GoogleIapTunnelInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

