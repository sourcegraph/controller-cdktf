package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceHaPolicy struct {
	// Specifies whether fast IP move is enabled, and if so, the mechanism to achieve it. Supported values are:.
	//
	// * 'DISABLED': Fast IP Move is disabled. You can only use the haPolicy.leader API to
	//               update the leader.
	//
	// * 'GARP_RA': Provides a method to very quickly define a new network endpoint as the
	//              leader. This method is faster than updating the leader using the
	//              haPolicy.leader API. Fast IP move works as follows: The VM hosting the
	//              network endpoint that should become the new leader sends either a
	//              Gratuitous ARP (GARP) packet (IPv4) or an ICMPv6 Router Advertisement(RA)
	//              packet (IPv6). Google Cloud immediately but temporarily associates the
	//              forwarding rule IP address with that VM, and both new and in-flight packets
	//              are quickly delivered to that VM. Possible values: ["DISABLED", "GARP_RA"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#fast_ip_move GoogleComputeRegionBackendService#fast_ip_move}
	FastIpMove *string `field:"optional" json:"fastIpMove" yaml:"fastIpMove"`
	// leader block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#leader GoogleComputeRegionBackendService#leader}
	Leader *GoogleComputeRegionBackendServiceHaPolicyLeader `field:"optional" json:"leader" yaml:"leader"`
}

