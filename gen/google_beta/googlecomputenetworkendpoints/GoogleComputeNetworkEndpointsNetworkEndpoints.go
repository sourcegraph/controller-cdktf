package googlecomputenetworkendpoints


type GoogleComputeNetworkEndpointsNetworkEndpoints struct {
	// IPv4 address of network endpoint.
	//
	// The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_network_endpoints#ip_address GoogleComputeNetworkEndpoints#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The name for a specific VM instance that the IP address belongs to.
	//
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone as the network endpoint group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_network_endpoints#instance GoogleComputeNetworkEndpoints#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type of 'GCE_VM_IP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_network_endpoints#port GoogleComputeNetworkEndpoints#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

