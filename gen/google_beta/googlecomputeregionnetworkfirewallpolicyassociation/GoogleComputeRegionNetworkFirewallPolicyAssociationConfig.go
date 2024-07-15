package googlecomputeregionnetworkfirewallpolicyassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRegionNetworkFirewallPolicyAssociationConfig struct {
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
	// The target that the firewall policy is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#attachment_target GoogleComputeRegionNetworkFirewallPolicyAssociation#attachment_target}
	AttachmentTarget *string `field:"required" json:"attachmentTarget" yaml:"attachmentTarget"`
	// The firewall policy ID of the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#firewall_policy GoogleComputeRegionNetworkFirewallPolicyAssociation#firewall_policy}
	FirewallPolicy *string `field:"required" json:"firewallPolicy" yaml:"firewallPolicy"`
	// The name for an association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#name GoogleComputeRegionNetworkFirewallPolicyAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#id GoogleComputeRegionNetworkFirewallPolicyAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#project GoogleComputeRegionNetworkFirewallPolicyAssociation#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The location of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#region GoogleComputeRegionNetworkFirewallPolicyAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_firewall_policy_association#timeouts GoogleComputeRegionNetworkFirewallPolicyAssociation#timeouts}
	Timeouts *GoogleComputeRegionNetworkFirewallPolicyAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

