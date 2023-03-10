package googlecomputerouterinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRouterInterfaceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// A unique name for the interface, required by GCE. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#name GoogleComputeRouterInterface#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the router this interface will be attached to.
	//
	// Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#router GoogleComputeRouterInterface#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#id GoogleComputeRouterInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name or resource link to the VLAN interconnect for this interface.
	//
	// Changing this forces a new interface to be created. Only one of vpn_tunnel and interconnect_attachment can be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#interconnect_attachment GoogleComputeRouterInterface#interconnect_attachment}
	InterconnectAttachment *string `field:"optional" json:"interconnectAttachment" yaml:"interconnectAttachment"`
	// IP address and range of the interface.
	//
	// The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#ip_range GoogleComputeRouterInterface#ip_range}
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// The ID of the project in which this interface's router belongs.
	//
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#project GoogleComputeRouterInterface#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region this interface's router sits in.
	//
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#region GoogleComputeRouterInterface#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#timeouts GoogleComputeRouterInterface#timeouts}
	Timeouts *GoogleComputeRouterInterfaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The name or resource link to the VPN tunnel this interface will be linked to.
	//
	// Changing this forces a new interface to be created. Only one of vpn_tunnel and interconnect_attachment can be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_router_interface#vpn_tunnel GoogleComputeRouterInterface#vpn_tunnel}
	VpnTunnel *string `field:"optional" json:"vpnTunnel" yaml:"vpnTunnel"`
}

