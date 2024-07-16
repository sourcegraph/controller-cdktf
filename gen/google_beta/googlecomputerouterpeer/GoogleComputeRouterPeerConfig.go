package googlecomputerouterpeer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRouterPeerConfig struct {
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
	// Name of the interface the BGP peer is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#interface GoogleComputeRouterPeer#interface}
	Interface *string `field:"required" json:"interface" yaml:"interface"`
	// Name of this BGP peer.
	//
	// The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#name GoogleComputeRouterPeer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Peer BGP Autonomous System Number (ASN). Each BGP interface may use a different value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#peer_asn GoogleComputeRouterPeer#peer_asn}
	PeerAsn *float64 `field:"required" json:"peerAsn" yaml:"peerAsn"`
	// The name of the Cloud Router in which this BgpPeer will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#router GoogleComputeRouterPeer#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// User-specified list of prefix groups to advertise in custom mode, which currently supports the following option:.
	//
	// * 'ALL_SUBNETS': Advertises all of the router's own VPC subnets.
	// This excludes any routes learned for subnets that use VPC Network
	// Peering.
	//
	//
	// Note that this field can only be populated if advertiseMode is 'CUSTOM'
	// and overrides the list defined for the router (in the "bgp" message).
	// These groups are advertised in addition to any specified prefixes.
	// Leave this field blank to advertise no custom groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#advertised_groups GoogleComputeRouterPeer#advertised_groups}
	AdvertisedGroups *[]*string `field:"optional" json:"advertisedGroups" yaml:"advertisedGroups"`
	// advertised_ip_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#advertised_ip_ranges GoogleComputeRouterPeer#advertised_ip_ranges}
	AdvertisedIpRanges interface{} `field:"optional" json:"advertisedIpRanges" yaml:"advertisedIpRanges"`
	// The priority of routes advertised to this BGP peer.
	//
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#advertised_route_priority GoogleComputeRouterPeer#advertised_route_priority}
	AdvertisedRoutePriority *float64 `field:"optional" json:"advertisedRoutePriority" yaml:"advertisedRoutePriority"`
	// User-specified flag to indicate which mode to use for advertisement.
	//
	// Valid values of this enum field are: 'DEFAULT', 'CUSTOM' Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#advertise_mode GoogleComputeRouterPeer#advertise_mode}
	AdvertiseMode *string `field:"optional" json:"advertiseMode" yaml:"advertiseMode"`
	// bfd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#bfd GoogleComputeRouterPeer#bfd}
	Bfd *GoogleComputeRouterPeerBfd `field:"optional" json:"bfd" yaml:"bfd"`
	// The status of the BGP peer connection.
	//
	// If set to false, any active session
	// with the peer is terminated and all associated routing information is removed.
	// If set to true, the peer connection can be established with routing information.
	// The default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#enable GoogleComputeRouterPeer#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Enable IPv4 traffic over BGP Peer. It is enabled by default if the peerIpAddress is version 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#enable_ipv4 GoogleComputeRouterPeer#enable_ipv4}
	EnableIpv4 interface{} `field:"optional" json:"enableIpv4" yaml:"enableIpv4"`
	// Enable IPv6 traffic over BGP Peer. If not specified, it is disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#enable_ipv6 GoogleComputeRouterPeer#enable_ipv6}
	EnableIpv6 interface{} `field:"optional" json:"enableIpv6" yaml:"enableIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#id GoogleComputeRouterPeer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#ip_address GoogleComputeRouterPeer#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// IPv4 address of the interface inside Google Cloud Platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#ipv4_nexthop_address GoogleComputeRouterPeer#ipv4_nexthop_address}
	Ipv4NexthopAddress *string `field:"optional" json:"ipv4NexthopAddress" yaml:"ipv4NexthopAddress"`
	// IPv6 address of the interface inside Google Cloud Platform.
	//
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#ipv6_nexthop_address GoogleComputeRouterPeer#ipv6_nexthop_address}
	Ipv6NexthopAddress *string `field:"optional" json:"ipv6NexthopAddress" yaml:"ipv6NexthopAddress"`
	// md5_authentication_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#md5_authentication_key GoogleComputeRouterPeer#md5_authentication_key}
	Md5AuthenticationKey *GoogleComputeRouterPeerMd5AuthenticationKey `field:"optional" json:"md5AuthenticationKey" yaml:"md5AuthenticationKey"`
	// IP address of the BGP interface outside Google Cloud Platform. Only IPv4 is supported. Required if 'ip_address' is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#peer_ip_address GoogleComputeRouterPeer#peer_ip_address}
	PeerIpAddress *string `field:"optional" json:"peerIpAddress" yaml:"peerIpAddress"`
	// IPv4 address of the BGP interface outside Google Cloud Platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#peer_ipv4_nexthop_address GoogleComputeRouterPeer#peer_ipv4_nexthop_address}
	PeerIpv4NexthopAddress *string `field:"optional" json:"peerIpv4NexthopAddress" yaml:"peerIpv4NexthopAddress"`
	// IPv6 address of the BGP interface outside Google Cloud Platform.
	//
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#peer_ipv6_nexthop_address GoogleComputeRouterPeer#peer_ipv6_nexthop_address}
	PeerIpv6NexthopAddress *string `field:"optional" json:"peerIpv6NexthopAddress" yaml:"peerIpv6NexthopAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#project GoogleComputeRouterPeer#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the router and BgpPeer reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#region GoogleComputeRouterPeer#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The URI of the VM instance that is used as third-party router appliances such as Next Gen Firewalls, Virtual Routers, or Router Appliances.
	//
	// The VM instance must be located in zones contained in the same region as
	// this Cloud Router. The VM instance is the peer side of the BGP session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#router_appliance_instance GoogleComputeRouterPeer#router_appliance_instance}
	RouterApplianceInstance *string `field:"optional" json:"routerApplianceInstance" yaml:"routerApplianceInstance"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#timeouts GoogleComputeRouterPeer#timeouts}
	Timeouts *GoogleComputeRouterPeerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

