package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateNetworkInterface struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#access_config GoogleComputeInstanceTemplate#access_config}
	AccessConfig interface{} `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// alias_ip_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#alias_ip_range GoogleComputeInstanceTemplate#alias_ip_range}
	AliasIpRange interface{} `field:"optional" json:"aliasIpRange" yaml:"aliasIpRange"`
	// ipv6_access_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#ipv6_access_config GoogleComputeInstanceTemplate#ipv6_access_config}
	Ipv6AccessConfig interface{} `field:"optional" json:"ipv6AccessConfig" yaml:"ipv6AccessConfig"`
	// The name or self_link of the network to attach this interface to.
	//
	// Use network attribute for Legacy or Auto subnetted networks and subnetwork for custom subnetted networks.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#network GoogleComputeInstanceTemplate#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The private IP address to assign to the instance. If empty, the address will be automatically assigned.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#network_ip GoogleComputeInstanceTemplate#network_ip}
	NetworkIp *string `field:"optional" json:"networkIp" yaml:"networkIp"`
	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#nic_type GoogleComputeInstanceTemplate#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// The networking queue count that's specified by users for the network interface.
	//
	// Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#queue_count GoogleComputeInstanceTemplate#queue_count}
	QueueCount *float64 `field:"optional" json:"queueCount" yaml:"queueCount"`
	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not.
	//
	// If not specified, IPV4_ONLY will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#stack_type GoogleComputeInstanceTemplate#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// The name of the subnetwork to attach this interface to.
	//
	// The subnetwork must exist in the same region this instance will be created in. Either network or subnetwork must be provided.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#subnetwork GoogleComputeInstanceTemplate#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The ID of the project in which the subnetwork belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#subnetwork_project GoogleComputeInstanceTemplate#subnetwork_project}
	SubnetworkProject *string `field:"optional" json:"subnetworkProject" yaml:"subnetworkProject"`
}

