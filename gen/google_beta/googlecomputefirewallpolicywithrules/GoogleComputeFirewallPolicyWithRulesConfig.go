package googlecomputefirewallpolicywithrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeFirewallPolicyWithRulesConfig struct {
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
	// The parent of this FirewallPolicy in the Cloud Resource Hierarchy. Format: organizations/{organization_id} or folders/{folder_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#parent GoogleComputeFirewallPolicyWithRules#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#rule GoogleComputeFirewallPolicyWithRules#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// A textual name of the security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#short_name GoogleComputeFirewallPolicyWithRules#short_name}
	ShortName *string `field:"required" json:"shortName" yaml:"shortName"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#description GoogleComputeFirewallPolicyWithRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#id GoogleComputeFirewallPolicyWithRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#timeouts GoogleComputeFirewallPolicyWithRules#timeouts}
	Timeouts *GoogleComputeFirewallPolicyWithRulesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

