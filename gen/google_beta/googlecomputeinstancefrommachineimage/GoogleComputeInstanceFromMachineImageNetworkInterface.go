package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageNetworkInterface struct {
	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#access_config GoogleComputeInstanceFromMachineImage#access_config}
	AccessConfig interface{} `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// An array of alias IP ranges for this network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#alias_ip_range GoogleComputeInstanceFromMachineImage#alias_ip_range}
	AliasIpRange interface{} `field:"optional" json:"aliasIpRange" yaml:"aliasIpRange"`
	// The prefix length of the primary internal IPv6 range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#internal_ipv6_prefix_length GoogleComputeInstanceFromMachineImage#internal_ipv6_prefix_length}
	InternalIpv6PrefixLength *float64 `field:"optional" json:"internalIpv6PrefixLength" yaml:"internalIpv6PrefixLength"`
	// ipv6_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#ipv6_access_config GoogleComputeInstanceFromMachineImage#ipv6_access_config}
	Ipv6AccessConfig interface{} `field:"optional" json:"ipv6AccessConfig" yaml:"ipv6AccessConfig"`
	// An IPv6 internal network address for this network interface.
	//
	// If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#ipv6_address GoogleComputeInstanceFromMachineImage#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The name or self_link of the network attached to this interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#network GoogleComputeInstanceFromMachineImage#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#network_attachment GoogleComputeInstanceFromMachineImage#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
	// The private IP address assigned to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#network_ip GoogleComputeInstanceFromMachineImage#network_ip}
	NetworkIp *string `field:"optional" json:"networkIp" yaml:"networkIp"`
	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#nic_type GoogleComputeInstanceFromMachineImage#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// The networking queue count that's specified by users for the network interface.
	//
	// Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#queue_count GoogleComputeInstanceFromMachineImage#queue_count}
	QueueCount *float64 `field:"optional" json:"queueCount" yaml:"queueCount"`
	// A full or partial URL to a security policy to add to this instance.
	//
	// If this field is set to an empty string it will remove the associated security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#security_policy GoogleComputeInstanceFromMachineImage#security_policy}
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not.
	//
	// If not specified, IPV4_ONLY will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#stack_type GoogleComputeInstanceFromMachineImage#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// The name or self_link of the subnetwork attached to this interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#subnetwork GoogleComputeInstanceFromMachineImage#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The project in which the subnetwork belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#subnetwork_project GoogleComputeInstanceFromMachineImage#subnetwork_project}
	SubnetworkProject *string `field:"optional" json:"subnetworkProject" yaml:"subnetworkProject"`
}

