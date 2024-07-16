package googlednspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDnsPolicyConfig struct {
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
	// User assigned name for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#name GoogleDnsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// alternative_name_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#alternative_name_server_config GoogleDnsPolicy#alternative_name_server_config}
	AlternativeNameServerConfig *GoogleDnsPolicyAlternativeNameServerConfig `field:"optional" json:"alternativeNameServerConfig" yaml:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Terraform'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#description GoogleDnsPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections.
	//
	// When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#enable_inbound_forwarding GoogleDnsPolicy#enable_inbound_forwarding}
	EnableInboundForwarding interface{} `field:"optional" json:"enableInboundForwarding" yaml:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#enable_logging GoogleDnsPolicy#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#id GoogleDnsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#networks GoogleDnsPolicy#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#project GoogleDnsPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#timeouts GoogleDnsPolicy#timeouts}
	Timeouts *GoogleDnsPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

