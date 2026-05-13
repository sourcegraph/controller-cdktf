package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceDynamicForwarding struct {
	// forward_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#forward_proxy GoogleComputeRegionBackendService#forward_proxy}
	ForwardProxy *GoogleComputeRegionBackendServiceDynamicForwardingForwardProxy `field:"optional" json:"forwardProxy" yaml:"forwardProxy"`
	// ip_port_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#ip_port_selection GoogleComputeRegionBackendService#ip_port_selection}
	IpPortSelection *GoogleComputeRegionBackendServiceDynamicForwardingIpPortSelection `field:"optional" json:"ipPortSelection" yaml:"ipPortSelection"`
}

