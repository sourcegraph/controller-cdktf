package googlecomputevpntunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeVpnTunnelConfig struct {
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
	// Name of the resource.
	//
	// The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#name GoogleComputeVpnTunnel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#shared_secret GoogleComputeVpnTunnel#shared_secret}
	SharedSecret *string `field:"required" json:"sharedSecret" yaml:"sharedSecret"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#description GoogleComputeVpnTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#id GoogleComputeVpnTunnel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway.
	//
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#ike_version GoogleComputeVpnTunnel#ike_version}
	IkeVersion *float64 `field:"optional" json:"ikeVersion" yaml:"ikeVersion"`
	// Labels to apply to this VpnTunnel.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#labels GoogleComputeVpnTunnel#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway.
	//
	// The value should be a CIDR formatted string,
	// for example '192.168.0.0/16'. The ranges should be disjoint.
	// Only IPv4 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#local_traffic_selector GoogleComputeVpnTunnel#local_traffic_selector}
	LocalTrafficSelector *[]*string `field:"optional" json:"localTrafficSelector" yaml:"localTrafficSelector"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#peer_external_gateway GoogleComputeVpnTunnel#peer_external_gateway}
	PeerExternalGateway *string `field:"optional" json:"peerExternalGateway" yaml:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#peer_external_gateway_interface GoogleComputeVpnTunnel#peer_external_gateway_interface}
	PeerExternalGatewayInterface *float64 `field:"optional" json:"peerExternalGatewayInterface" yaml:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	//
	// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
	// ID in the peer GCP VPN gateway.
	// This field must reference a 'google_compute_ha_vpn_gateway' resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#peer_gcp_gateway GoogleComputeVpnTunnel#peer_gcp_gateway}
	PeerGcpGateway *string `field:"optional" json:"peerGcpGateway" yaml:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#peer_ip GoogleComputeVpnTunnel#peer_ip}
	PeerIp *string `field:"optional" json:"peerIp" yaml:"peerIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#project GoogleComputeVpnTunnel#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#region GoogleComputeVpnTunnel#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway.
	//
	// The value should be a CIDR formatted string,
	// for example '192.168.0.0/16'. The ranges should be disjoint.
	// Only IPv4 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#remote_traffic_selector GoogleComputeVpnTunnel#remote_traffic_selector}
	RemoteTrafficSelector *[]*string `field:"optional" json:"remoteTrafficSelector" yaml:"remoteTrafficSelector"`
	// URL of router resource to be used for dynamic routing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#router GoogleComputeVpnTunnel#router}
	Router *string `field:"optional" json:"router" yaml:"router"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#target_vpn_gateway GoogleComputeVpnTunnel#target_vpn_gateway}
	TargetVpnGateway *string `field:"optional" json:"targetVpnGateway" yaml:"targetVpnGateway"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#timeouts GoogleComputeVpnTunnel#timeouts}
	Timeouts *GoogleComputeVpnTunnelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// URL of the VPN gateway with which this VPN tunnel is associated.
	//
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a 'google_compute_ha_vpn_gateway' resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#vpn_gateway GoogleComputeVpnTunnel#vpn_gateway}
	VpnGateway *string `field:"optional" json:"vpnGateway" yaml:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_vpn_tunnel#vpn_gateway_interface GoogleComputeVpnTunnel#vpn_gateway_interface}
	VpnGatewayInterface *float64 `field:"optional" json:"vpnGatewayInterface" yaml:"vpnGatewayInterface"`
}

