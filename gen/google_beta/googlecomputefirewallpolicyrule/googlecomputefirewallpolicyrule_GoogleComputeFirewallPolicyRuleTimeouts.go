package googlecomputefirewallpolicyrule


type GoogleComputeFirewallPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_firewall_policy_rule#create GoogleComputeFirewallPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_firewall_policy_rule#delete GoogleComputeFirewallPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_firewall_policy_rule#update GoogleComputeFirewallPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

