package googlenetworksecurityfirewallendpointassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityFirewallEndpointAssociationConfig struct {
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
	// The URL of the firewall endpoint that is being associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#firewall_endpoint GoogleNetworkSecurityFirewallEndpointAssociation#firewall_endpoint}
	FirewallEndpoint *string `field:"required" json:"firewallEndpoint" yaml:"firewallEndpoint"`
	// The location (zone) of the firewall endpoint association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#location GoogleNetworkSecurityFirewallEndpointAssociation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the firewall endpoint association resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#name GoogleNetworkSecurityFirewallEndpointAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL of the network that is being associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#network GoogleNetworkSecurityFirewallEndpointAssociation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Whether the association is disabled. True indicates that traffic will not be intercepted.
	//
	// ~> **Note:** The API will reject the request if this value is set to true when creating the resource,
	// otherwise on an update the association can be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#disabled GoogleNetworkSecurityFirewallEndpointAssociation#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#id GoogleNetworkSecurityFirewallEndpointAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A map of key/value label pairs to assign to the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#labels GoogleNetworkSecurityFirewallEndpointAssociation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#parent GoogleNetworkSecurityFirewallEndpointAssociation#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#timeouts GoogleNetworkSecurityFirewallEndpointAssociation#timeouts}
	Timeouts *GoogleNetworkSecurityFirewallEndpointAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The URL of the TlsInspectionPolicy that is being associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_firewall_endpoint_association#tls_inspection_policy GoogleNetworkSecurityFirewallEndpointAssociation#tls_inspection_policy}
	TlsInspectionPolicy *string `field:"optional" json:"tlsInspectionPolicy" yaml:"tlsInspectionPolicy"`
}

