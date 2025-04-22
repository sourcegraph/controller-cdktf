package googlecomputeregionnetworkfirewallpolicywithrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRegionNetworkFirewallPolicyWithRulesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// User-provided name of the Network firewall policy.
	//
	// The name should be unique in the project in which the firewall policy is created.
	// The name must be 1-63 characters long, and comply with RFC1035. Specifically,
	// the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?
	// which means the first character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#name GoogleComputeRegionNetworkFirewallPolicyWithRules#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#rule GoogleComputeRegionNetworkFirewallPolicyWithRules#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#description GoogleComputeRegionNetworkFirewallPolicyWithRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#id GoogleComputeRegionNetworkFirewallPolicyWithRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#project GoogleComputeRegionNetworkFirewallPolicyWithRules#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#region GoogleComputeRegionNetworkFirewallPolicyWithRules#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#timeouts GoogleComputeRegionNetworkFirewallPolicyWithRules#timeouts}
	Timeouts *GoogleComputeRegionNetworkFirewallPolicyWithRulesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

