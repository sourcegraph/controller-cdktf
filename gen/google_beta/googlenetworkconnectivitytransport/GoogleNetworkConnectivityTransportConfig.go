package googlenetworkconnectivitytransport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkConnectivityTransportConfig struct {
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
	// Name of the resource, see google.aip.dev/122 for resource naming.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#name GoogleNetworkConnectivityTransport#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Resource URL of the Network that will be peered with this Transport.
	//
	// This field must be provided during resource creation and cannot be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#network GoogleNetworkConnectivityTransport#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The region of this resource.
	//
	// This is required to construct the resource name, but is not sent to the API since the region is already contained in the parent field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#region GoogleNetworkConnectivityTransport#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Resource URL of the remoteTransportProfile that this Transport is connecting to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#remote_profile GoogleNetworkConnectivityTransport#remote_profile}
	RemoteProfile *string `field:"required" json:"remoteProfile" yaml:"remoteProfile"`
	// Administrative state of the underlying connectivity.
	//
	// If set to true (default), connectivity should be available between your environments. If set to false, the connectivity over these links is disabled. Disabling your Transport does not affect billing, and retains the underlying network bandwidth associated with the connectivity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#admin_enabled GoogleNetworkConnectivityTransport#admin_enabled}
	AdminEnabled interface{} `field:"optional" json:"adminEnabled" yaml:"adminEnabled"`
	// List of IP Prefixes that will be advertised to the remote provider. Both IPv4 and IPv6 addresses are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#advertised_routes GoogleNetworkConnectivityTransport#advertised_routes}
	AdvertisedRoutes *[]*string `field:"optional" json:"advertisedRoutes" yaml:"advertisedRoutes"`
	// Controls whether resources proposed by the Transport are automatically accepted on behalf of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#auto_accept GoogleNetworkConnectivityTransport#auto_accept}
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// Bandwidth of the Transport. This must be one of the supported bandwidths for the remote profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#bandwidth GoogleNetworkConnectivityTransport#bandwidth}
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#description GoogleNetworkConnectivityTransport#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The NCC Hub that the Transport should attach to.
	//
	// The hub must be in the same project as the Transport.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#hub GoogleNetworkConnectivityTransport#hub}
	Hub *string `field:"optional" json:"hub" yaml:"hub"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#id GoogleNetworkConnectivityTransport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#labels GoogleNetworkConnectivityTransport#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// [Output only] The maximum transmission unit (MTU) of a packet that can be sent over this transport.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#mtu_limit GoogleNetworkConnectivityTransport#mtu_limit}
	MtuLimit *float64 `field:"optional" json:"mtuLimit" yaml:"mtuLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#project GoogleNetworkConnectivityTransport#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Key used for establishing a connection with the remote transport.
	//
	// This key can only be provided if the profile supports an INPUT key flow and the resource is in the PENDING_KEY state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#provided_activation_key GoogleNetworkConnectivityTransport#provided_activation_key}
	ProvidedActivationKey *string `field:"optional" json:"providedActivationKey" yaml:"providedActivationKey"`
	// Controls whether a Routing VPC Spoke should be created and attached to the NCC Hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#psc_routing_enabled GoogleNetworkConnectivityTransport#psc_routing_enabled}
	PscRoutingEnabled interface{} `field:"optional" json:"pscRoutingEnabled" yaml:"pscRoutingEnabled"`
	// The user supplied account id for the CSP associated with the remote profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#remote_account_id GoogleNetworkConnectivityTransport#remote_account_id}
	RemoteAccountId *string `field:"optional" json:"remoteAccountId" yaml:"remoteAccountId"`
	// IP version stack for the established connectivity. Possible values: ["IPV4_IPV6", "IPV4_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#stack_type GoogleNetworkConnectivityTransport#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport#timeouts GoogleNetworkConnectivityTransport#timeouts}
	Timeouts *GoogleNetworkConnectivityTransportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

