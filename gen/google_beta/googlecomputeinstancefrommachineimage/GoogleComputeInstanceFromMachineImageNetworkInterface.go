package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageNetworkInterface struct {
	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#access_config GoogleComputeInstanceFromMachineImage#access_config}
	AccessConfig interface{} `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// An array of alias IP ranges for this network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#alias_ip_range GoogleComputeInstanceFromMachineImage#alias_ip_range}
	AliasIpRange interface{} `field:"optional" json:"aliasIpRange" yaml:"aliasIpRange"`
	// ipv6_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#ipv6_access_config GoogleComputeInstanceFromMachineImage#ipv6_access_config}
	Ipv6AccessConfig interface{} `field:"optional" json:"ipv6AccessConfig" yaml:"ipv6AccessConfig"`
	// The name or self_link of the network attached to this interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#network GoogleComputeInstanceFromMachineImage#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The private IP address assigned to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#network_ip GoogleComputeInstanceFromMachineImage#network_ip}
	NetworkIp *string `field:"optional" json:"networkIp" yaml:"networkIp"`
	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#nic_type GoogleComputeInstanceFromMachineImage#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// The networking queue count that's specified by users for the network interface.
	//
	// Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#queue_count GoogleComputeInstanceFromMachineImage#queue_count}
	QueueCount *float64 `field:"optional" json:"queueCount" yaml:"queueCount"`
	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not.
	//
	// If not specified, IPV4_ONLY will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#stack_type GoogleComputeInstanceFromMachineImage#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// The name or self_link of the subnetwork attached to this interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#subnetwork GoogleComputeInstanceFromMachineImage#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The project in which the subnetwork belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_from_machine_image#subnetwork_project GoogleComputeInstanceFromMachineImage#subnetwork_project}
	SubnetworkProject *string `field:"optional" json:"subnetworkProject" yaml:"subnetworkProject"`
}

